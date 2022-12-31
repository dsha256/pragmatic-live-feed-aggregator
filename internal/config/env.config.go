package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	PragmaticFeedWsURL  = "PRAGMATIC_FEED_WS_URL"
	CasinoID            = "CASINO_ID"
	TableIDs            = "TABLE_IDS"
	CurrencyIDs         = "CURRENCY_IDS"
	RedisPort           = "REDIS_PORT"
	ServerPort          = "SERVER_PORT"
	PusherChannelID     = "PUSHER_CHANNEL_ID"
	PusherPeriodMinutes = "PUSHER_PERIOD_MINUTES"
	PusherAppID         = "PUSHER_APP_ID"
	PusherKey           = "PUSHER_KEY"
	PusherSecret        = "PUSHER_SECRET"
	PusherCluster       = "PUSHER_CLUSTER"
)

type ENV struct {
	PragmaticFeedWsURL  string
	CasinoID            string
	TableIDs            string
	CurrencyIDs         string
	RedisPort           string
	ServerPort          string
	PusherChannelID     string
	PusherPeriodMinutes string
	PusherAppID         string
	PusherKey           string
	PusherSecret        string
	PusherCluster       string
}

func (env *ENV) Load() *ENV {
	env.PragmaticFeedWsURL = os.Getenv(PragmaticFeedWsURL)
	env.CasinoID = os.Getenv(CasinoID)
	env.TableIDs = os.Getenv(TableIDs)
	env.CurrencyIDs = os.Getenv(CurrencyIDs)
	env.RedisPort = os.Getenv(RedisPort)
	env.ServerPort = os.Getenv(ServerPort)
	env.PusherChannelID = os.Getenv(PusherChannelID)
	env.PusherPeriodMinutes = os.Getenv(PusherPeriodMinutes)
	env.PusherAppID = os.Getenv(PusherAppID)
	env.PusherKey = os.Getenv(PusherKey)
	env.PusherSecret = os.Getenv(PusherSecret)
	env.PusherCluster = os.Getenv(PusherCluster)

	return env
}

// TODO: Reflection can make your life easier. Give this present to yourself this new year :)

func (env *ENV) GetTableIDs() []string {
	res := strings.Split(strCleanUp(env.TableIDs), ",")
	checkEnvVarOnEmptiness(TableIDs, res)
	return res
}

func (env *ENV) GetCurrencyIDs() []string {
	res := strings.Split(strCleanUp(env.CurrencyIDs), ",")
	checkEnvVarOnEmptiness(CurrencyIDs, res)
	return res
}

func (env *ENV) GetCasinoID() string {
	checkEnvVarOnEmptiness(CasinoID, env.CasinoID)
	return strCleanUp(env.CasinoID)
}

func (env *ENV) GetRedisPort() string {
	checkEnvVarOnEmptiness(RedisPort, env.RedisPort)
	return strCleanUp(env.RedisPort)
}

func (env *ENV) GetServerPort() string {
	checkEnvVarOnEmptiness(ServerPort, env.ServerPort)
	return strCleanUp(env.ServerPort)
}

func (env *ENV) GetPusherChannelID() string {
	checkEnvVarOnEmptiness(PusherChannelID, env.PusherChannelID)
	return strCleanUp(env.PusherChannelID)
}

func (env *ENV) GetPusherPeriodMinutes() int {
	cleanedUpMinutes := strCleanUp(env.PusherPeriodMinutes)
	minutes, err := strconv.Atoi(cleanedUpMinutes)
	if err != nil {
		log.Fatalf("Error convertin %s to int", PusherPeriodMinutes)
	}
	checkEnvVarOnEmptiness(PusherPeriodMinutes, minutes)
	return minutes
}

func (env *ENV) GetPusherAppID() string {
	checkEnvVarOnEmptiness(PusherAppID, env.PusherAppID)
	return strCleanUp(env.PusherAppID)
}

func (env *ENV) GetPusherKey() string {
	checkEnvVarOnEmptiness(PusherKey, env.PusherKey)
	return strCleanUp(env.PusherKey)
}

func (env *ENV) GetPusherSecret() string {
	checkEnvVarOnEmptiness(PusherSecret, env.PusherSecret)
	return strCleanUp(env.PusherSecret)
}

func (env *ENV) GetPusherCluster() string {
	checkEnvVarOnEmptiness(PusherCluster, env.PusherCluster)
	return strCleanUp(env.PusherCluster)
}
