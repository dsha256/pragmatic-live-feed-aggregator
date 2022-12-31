package pusher

import (
	"context"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/repo"
	"github.com/pusher/pusher-http-go/v5"
	"log"
	"time"
)

type Client struct {
	pusherClient          *pusher.Client
	channelID             string
	pusherPeriodInMinutes int
	repo                  *repo.RedisRepository
}

func NewClient(
	appID string,
	key string,
	secret string,
	cluster string,
	channelID string,
	pusherPeriodInMinutes int,
	repo *repo.RedisRepository,
) *Client {
	return &Client{
		pusherClient: &pusher.Client{
			AppID:   appID,
			Key:     key,
			Secret:  secret,
			Cluster: cluster,
			Secure:  true,
		},
		channelID:             channelID,
		pusherPeriodInMinutes: pusherPeriodInMinutes,
		repo:                  repo,
	}
}

func (c *Client) StartPushing(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(c.pusherPeriodInMinutes) * time.Minute)
	done := make(chan bool)
	fatal := make(chan string)
	go func() {
		for {
			select {
			case <-done:
				log.Println("Pusher client has done")
			case <-fatal:
				log.Fatalf("Pusher client accidently stopped: %v", <-fatal)
			case <-ticker.C:
				log.Println("Pushed updated pragmatic table's data to the pusher channel")
				err := c.pushData(ctx)
				if err != nil {
					fatal <- err.Error()
				}
			}
		}
	}()
}

func (c *Client) pushData(ctx context.Context) error {
	pragmaticTables, err := c.repo.ListTables(ctx)
	if err != nil {
		return err
	}
	err = c.pusherClient.Trigger(c.channelID, "update-pragmatic-tables", pragmaticTables)
	if err != nil {
		return err
	}
	return nil
}
