package pragmaticlivefeed

import (
	"context"
	"github.com/pusher/pusher-http-go/v5"
	"log"
	"time"
)

type pusherService struct {
	client                *pusher.Client
	channelID             string
	pusherPeriodInMinutes int
	repo                  AggregateRepository
}

func NewPusherService(
	client *pusher.Client,
	channelID string,
	pusherPeriodInMinutes int,
	repo AggregateRepository,
) *pusherService {
	return &pusherService{
		client:                client,
		channelID:             channelID,
		pusherPeriodInMinutes: pusherPeriodInMinutes,
		repo:                  repo,
	}
}

func (s *pusherService) StartPushing(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(s.pusherPeriodInMinutes) * time.Minute)
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
				log.Println("Pushed updated data of pragmatic live feed tables to the pusher channel")
				err := pushData(ctx, s.client, s.channelID, s.repo)
				if err != nil {
					fatal <- err.Error()
				}
			}
		}
	}()
}

func pushData(ctx context.Context, ps *pusher.Client, channelID string, repo AggregateRepository) error {
	pragmaticTables, err := repo.ListTables(ctx)
	if err != nil {
		return err
	}
	err = ps.Trigger(channelID, "update-pragmatic-live-feed-tables", pragmaticTables)
	if err != nil {
		return err
	}
	return nil
}
