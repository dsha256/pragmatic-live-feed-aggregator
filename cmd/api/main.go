package main

import (
	"context"
	"fmt"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/config"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/pragmaticlivefeed"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/repository"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/server"
	"github.com/pusher/pusher-http-go/v5"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// @title Pragmatic Live Feed API Documentation
// @version 1.0.0
// @host localhost:8080
// @BasePath /api/v1/pragmatic_live_feed
func main() {
	errC, err := bootstrap()
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func bootstrap() (<-chan error, error) {
	errC := make(chan error, 1)
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ctx.Done()
		defer func() {
			stop()
			close(errC)
		}()
	}()

	env := config.ENV{}
	env.Load()
	pragmaticFeedWsURL := env.GetPragmaticFeedWsURL()
	casinoID := env.GetCasinoID()
	tableIDs := env.GetTableIDs()
	currencyIDs := env.GetCurrencyIDs()
	serverPort := env.GetServerPort()
	pusherChannelID := env.GetPusherChannelID()
	pusherPeriodMinutes := env.GetPusherPeriodMinutes()
	pusherAppID := env.GetPusherAppID()
	pusherKey := env.GetPusherKey()
	pusherSecret := env.GetPusherSecret()
	pusherCluster := env.GetPusherCluster()

	aggregateRepo := repository.NewAggregator()
	service := pragmaticlivefeed.NewService(aggregateRepo)

	httpServer := server.NewHTTP(service)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), httpServer)
		if err != nil {
			errC <- err
		}
	}()

	pusherClient := pusher.Client{
		AppID:   pusherAppID,
		Key:     pusherKey,
		Secret:  pusherSecret,
		Cluster: pusherCluster,
		Secure:  true,
	}
	pusherSvc := pragmaticlivefeed.NewPusherService(&pusherClient, pusherChannelID, pusherPeriodMinutes, aggregateRepo)
	pusherSvc.StartPushing(ctx)

	wsClient := pragmaticlivefeed.NewWSService(ctx, aggregateRepo, pragmaticFeedWsURL, casinoID, tableIDs, currencyIDs)
	wsClient.StartClients()

	return errC, nil
}
