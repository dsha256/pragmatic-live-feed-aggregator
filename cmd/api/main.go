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
)

// @title Pragmatic Live Feed API Documentation
// @version 1.0.0
// @host localhost:8080
// @BasePath /api/v1/pragmatic_live_feed
func main() {
	start()
}

func start() {

	env := config.ENV{}
	env.Load()
	pragmaticFeedWsURL := env.GetPragmaticFeedWsURL()
	_ = pragmaticFeedWsURL
	casinoID := env.GetCasinoID()
	_ = casinoID
	tableIDs := env.GetTableIDs()
	_ = tableIDs
	currencyIDs := env.GetCurrencyIDs()
	_ = currencyIDs
	serverPort := env.GetServerPort()
	pusherChannelID := env.GetPusherChannelID()
	_ = pusherChannelID
	pusherPeriodMinutes := env.GetPusherPeriodMinutes()
	_ = pusherPeriodMinutes
	pusherAppID := env.GetPusherAppID()
	_ = pusherAppID
	pusherKey := env.GetPusherKey()
	_ = pusherKey
	pusherSecret := env.GetPusherSecret()
	_ = pusherSecret
	pusherCluster := env.GetPusherCluster()
	_ = pusherCluster
	ctx := context.Background()
	_ = ctx

	aggregateRepo := repository.NewAggregator()
	service := pragmaticlivefeed.NewService(aggregateRepo)

	httpServer := server.NewHTTP(service)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), httpServer)
		if err != nil {
			panicOnErr(err, fmt.Sprintf("failed to start server on port %s", serverPort))
		}
	}()

	//err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), httpServer)
	//if err != nil {
	//	panicOnErr(err, fmt.Sprintf("failed to start server on port %s", serverPort))
	//}

	pusherClient := pusher.Client{
		AppID:   pusherAppID,
		Key:     pusherKey,
		Secret:  pusherSecret,
		Cluster: pusherCluster,
		Secure:  true,
	}
	pusherSvc := pragmaticlivefeed.NewPusherService(&pusherClient, pusherChannelID, pusherPeriodMinutes, aggregateRepo)
	_ = pusherSvc
	//pusherSvc.StartPushing(ctx)

	// TODO: improve the runner func related processes synchronization using the channels to avoid strict ordering.
	// 	At this time, the code below must be at the end of the runner func, cause it takes an additional responsibility
	// 	to force the main thread to wait until all the other services are done their works.
	wsClient := pragmaticlivefeed.NewWSService(ctx, aggregateRepo, pragmaticFeedWsURL, casinoID, tableIDs, currencyIDs)
	wsClient.StartClients()
}

func panicOnErr(err error, msg string) {
	if err != nil {
		log.Panic(msg, err)
	}
}
