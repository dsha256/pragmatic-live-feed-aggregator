package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/repo"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	// golang.org/x/net/websocket currently lacks some features: https://pkg.go.dev/golang.org/x/net/websocket
	"github.com/gorilla/websocket"
)

type WS struct {
	sync.RWMutex
	ctx         context.Context
	DB          repo.InMemoryDataBase
	WSURL       string
	CasinoID    string
	TableIDs    []string
	CurrencyIDs []string
}

func NewClient(
	ctx context.Context,
	db repo.InMemoryDataBase,
	wsURL string,
	casinoID string,
	tableIDs []string,
	currencyIDs []string,
) WS {
	return WS{
		ctx:         ctx,
		DB:          db,
		WSURL:       wsURL,
		CasinoID:    casinoID,
		TableIDs:    tableIDs,
		CurrencyIDs: currencyIDs,
	}
}

var done chan any
var interrupt chan os.Signal
var socketUrl = "wss://dga.pragmaticplaylive.net/ws"

func (ws *WS) StartClients() {
	done = make(chan interface{})    // Channel to indicate that the receiverHandler is done
	interrupt = make(chan os.Signal) // Channel to listen for interrupt signal to terminate gracefully

	signal.Notify(interrupt, os.Interrupt) // Notify the interrupt channel for SIGINT

	wg := sync.WaitGroup{}
	for _, message := range ws.generateWSMessages() {
		log.Println(string(message))

		wg.Add(1)
		msg := message
		go func([]byte) {
			go ws.startClient(msg, &wg)
		}(msg)
	}

	wg.Wait()
}

func (ws *WS) pushReceivedDataToDB(connection *websocket.Conn) {
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		var pragmaticTable dto.PragmaticTable
		err = json.Unmarshal(msg, &pragmaticTable)
		if err != nil {
			log.Println("Error in unmarshal:", err)
		}

		ws.Lock()
		err = ws.DB.AddTable(ws.ctx, pragmaticTable)
		ws.Unlock()
		if err != nil {
			log.Println("Error in adding table to the DB:", err)
		}
	}
}

func (ws *WS) startClient(msg []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	socketUrl := socketUrl
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()

	go ws.pushReceivedDataToDB(conn)

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

func (ws *WS) generateWSMessages() [][]byte {
	var messages [][]byte
	messageTemplate := "{\"type\":\"subscribe\",\"key\":\"%s\",\"casinoId\":\"%s\",\"currency\":\"%s\"}"
	for _, tableID := range ws.TableIDs {
		for _, currencyID := range ws.CurrencyIDs {
			messages = append(messages, []byte(fmt.Sprintf(messageTemplate, tableID, ws.CasinoID, currencyID)))
		}
	}
	return messages
}
