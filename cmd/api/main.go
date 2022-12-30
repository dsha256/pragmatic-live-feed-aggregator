package main

import (
	"context"
	"fmt"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/config"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/repo"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/server"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/ws"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
)

func main() {
	start()
}

func start() {
	env := config.ENV{}
	env.Load()

	ctx := context.Background()

	redisPort := env.GetRedisPort()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + redisPort,
		Password: "",
		DB:       0,
	})
	redisRepo := repo.NewRedisRepository(redisClient)

	httpServer := server.NewHTTP(&redisRepo)
	serverPort := env.GetServerPort()
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), httpServer)
		if err != nil {
			log.Fatalf("Error starting a server on port %s: %s", serverPort, err)
		}
	}()

	wsURL := env.PragmaticFeedWsURL
	casinoID := env.GetCasinoID()
	tableIDs := env.GetTableIDs()
	currencyIDs := env.GetCurrencyIDs()
	wsClient := ws.NewClient(ctx, &redisRepo, wsURL, casinoID, tableIDs, currencyIDs)
	//_ = wsClient
	wsClient.StartClients()
}
