package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/config"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/pragmaticlivefeed"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/repository"
	"github.com/dsha256/pragmatic-live-feed-aggregator/internal/server"
	"github.com/go-redis/redis/v8"
	"github.com/pusher/pusher-http-go/v5"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"os"
)

const serviceName = "pragmatic_live_feed_aggregator"

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
	redisPort := env.GetRedisPort()
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

	db := redis.NewClient(&redis.Options{
		Addr:     "redis:" + redisPort,
		Password: "",
		DB:       0,
	})
	aggregateRepo := repository.NewAggregator(db)
	service := pragmaticlivefeed.NewService(aggregateRepo)

	httpServer := server.NewHTTP(service)

	//err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), httpServer)
	//if err != nil {
	//	//log.Fatalf("Error starting a server on port %s: %s", serverPort, err)
	//	panicOnErr(err, fmt.Sprintf("failed to start server on port %s", serverPort))
	//}
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), httpServer)
		if err != nil {
			//log.Fatalf("Error starting a server on port %s: %s", serverPort, err)
			panicOnErr(err, fmt.Sprintf("failed to start server on port %s", serverPort))
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
	_ = pusherSvc
	//pusherSvc.StartPushing(ctx)

	// TODO: improve the runner func related processes synchronization using the channels to avoid strict ordering.
	// At this time, the code below must be at the end of the runner func, cause it takes an additional responsibility
	// to force the main thread to wait until all the other services are done their works.
	wsClient := pragmaticlivefeed.NewWSService(ctx, aggregateRepo, pragmaticFeedWsURL, casinoID, tableIDs, currencyIDs)
	wsClient.StartClients()
}

func panicOnErr(err error, msg string) {
	if err != nil {
		log.Panic(msg, err)
	}
}

func getLogger() *zap.Logger {
	encoderConfig := ecszap.EncoderConfig{
		//EncodeName:     customNameEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   ecszap.FullCallerEncoder,
	}
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger
}

func getZapLogger() *zap.Logger {
	checkErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	// TODO: Create JSON file
	rawJSON := []byte(`{
						   "level": "debug",
						   "encoding": "json",
						   "outputPaths": ["stdout"],
						   "errorOutputPaths": ["stderr"],
						   "encoderConfig": {
							 "messageKey": "message",
							 "levelKey": "level",
							 "levelEncoder": "lowercase"
						   }
 						}`)
	var cfg zap.Config
	err := json.Unmarshal(rawJSON, &cfg)
	checkErr(err)

	logger, err := cfg.Build()
	checkErr(err)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		checkErr(err)
	}(logger)
	return logger
}
