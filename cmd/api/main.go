package main

import (
	"context"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/config"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/repo"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/ws"
	"github.com/go-redis/redis/v8"
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

	wsURL := env.PragmaticFeedWsURL
	casinoID := env.GetCasinoID()
	tableIDs := env.GetTableIDs()
	currencyIDs := env.GetCurrencyIDs()
	wsClient := ws.NewClient(ctx, &redisRepo, wsURL, casinoID, tableIDs, currencyIDs)
	wsClient.StartClients()
}
