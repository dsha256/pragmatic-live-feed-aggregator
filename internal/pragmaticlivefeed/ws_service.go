package pragmaticlivefeed

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

var done chan any
var interrupt chan os.Signal

type wsService struct {
	sync.RWMutex
	ctx         context.Context
	repo        AggregateRepository
	wsURL       string
	casinoID    string
	tableIDs    []string
	currencyIDs []string
}

func NewWSService(
	ctx context.Context,
	repo AggregateRepository,
	wsURL string,
	casinoID string,
	tableIDs []string,
	currencyIDs []string,
) *wsService {
	return &wsService{
		ctx:         ctx,
		repo:        repo,
		wsURL:       wsURL,
		casinoID:    casinoID,
		tableIDs:    tableIDs,
		currencyIDs: currencyIDs,
	}
}

func (s *wsService) StartClients() {
	done = make(chan interface{})    // Channel to indicate that the receiverHandler is done
	interrupt = make(chan os.Signal) // Channel to listen for interrupt signal to terminate gracefully

	signal.Notify(interrupt, os.Interrupt) // Notify the interrupt channel for SIGINT

	wg := sync.WaitGroup{}
	generatedMsgs := generateWSMessages(s.casinoID, s.tableIDs, s.currencyIDs)
	for _, message := range generatedMsgs {
		log.Println(string(message))

		wg.Add(1)
		msg := message
		go func([]byte) {
			go s.StartClient(msg, &wg)
		}(msg)
	}

	wg.Wait()
}

func (s *wsService) PushReceivedDataToDB(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		var pragmaticTable dto.PragmaticTable
		err = json.Unmarshal(msg, &pragmaticTable)
		if err != nil {
			log.Println("Error in unmarshal:", err)
		}

		s.Lock()
		err = s.repo.AddTable(s.ctx, pragmaticTable)
		s.Unlock()
		if err != nil {
			log.Println("Error in adding table to the DB:", err)
		}
	}
}

func (s *wsService) StartClient(msg []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	socketUrl := s.wsURL
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()

	go s.PushReceivedDataToDB(conn)

	for {
		select {
		case <-time.After(time.Duration(1) * time.Second * 3):
			// Send an echo packet every second
			err := conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error during writing to websocket:", err)
				return
			}

		case <-interrupt:
			// Received a SIGINT (Ctrl + C). Terminate gracefully...
			log.Println("Received SIGINT interrupt signal. Closing all pending connections")

			// Close the websocket connection
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error during closing websocket:", err)
				return
			}

			select {
			case <-done:
				log.Println("Receiver Channel Closed! Exiting....")
				//wg.Done()
			case <-time.After(time.Duration(1) * time.Second):
				log.Println("Timeout in closing receiving channel. Exiting...")
				//wg.Done()
			}
			return
		}
	}
}

func generateWSMessages(casinoID string, tableIDs, currencyIDs []string) [][]byte {
	var messages [][]byte
	messageTemplate := "{\"type\":\"subscribe\",\"key\":\"%s\",\"casinoId\":\"%s\",\"currency\":\"%s\"}"
	for _, tableID := range tableIDs {
		for _, currencyID := range currencyIDs {
			messages = append(messages, []byte(fmt.Sprintf(messageTemplate, tableID, casinoID, currencyID)))
		}
	}
	return messages
}
