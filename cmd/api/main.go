package main

import (
	"context"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/config"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/repo"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/server"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/ws"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
)

func main() {
	env := config.ENV{}
	env.Load()

	ctx := context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	redisRepo := repo.NewRedisRepository(redisClient)

	httpServer := server.NewHTTP(&redisRepo)
	// TODO: Get port from env vars
	err := http.ListenAndServe(":8080", httpServer)
	if err != nil {
		log.Printf("can't run the server on port 8080: %v", err)
	}

	wsURL := env.PragmaticFeedWsURL
	casinoID := env.GetCasinoID()
	tableIDs := env.GetTableIDs()
	currencyIDs := env.GetCurrencyIDs()
	wsClient := ws.NewClient(ctx, &redisRepo, wsURL, casinoID, tableIDs, currencyIDs)
	//_ = wsClient
	wsClient.StartClients()
}
