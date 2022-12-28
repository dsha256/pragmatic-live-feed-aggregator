package main

import (
	"context"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/repo"
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
	"github.com/go-redis/redis/v8"
	"log"
)

func main() {
	ctx := context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	redisRepo := repo.NewRedisRepository(redisClient)
	_ = redisRepo
	redisRepo.AddTable(ctx, dto.PragmaticTable{TableId: "255", Currency: "USD"})
	redisRepo.AddTable(ctx, dto.PragmaticTable{TableId: "256", Currency: "EUR"})
	redisRepo.AddTable(ctx, dto.PragmaticTable{TableId: "256", Currency: "USD"})
	a, _ := redisRepo.GetTableByTableAndCurrencyIDs(ctx, "256", "USD")
	log.Println(a)
}
