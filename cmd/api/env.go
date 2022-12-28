package main

import (
	"os"
	"strings"
)

const (
	PushedPeriodMinutes = "PUSHER_PERIOD_MINUTES"
	PragmaticFeedWsURL  = "PRAGMATIC_FEED_WS_URL"
	PusherChannelID     = "PUSHER_CHANNEL_ID"
	CasinoID            = "CASINO_ID"
	TableIDs            = "TABLE_IDS"
	CurrencyIDs         = "CURRENCY_IDS"
)

type ENV struct {
	PushedPeriodMinutes string
	PragmaticFeedWsURL  string
	PusherChannelID     string
	CasinoID            string
	TableIDs            string
	CurrencyIDs         string
}

func (env *ENV) Load() *ENV {
	env.PushedPeriodMinutes = os.Getenv(PushedPeriodMinutes)
	env.PragmaticFeedWsURL = os.Getenv(PragmaticFeedWsURL)
	env.PusherChannelID = os.Getenv(PusherChannelID)
	env.CasinoID = os.Getenv(CasinoID)
	env.TableIDs = os.Getenv(TableIDs)
	env.CurrencyIDs = os.Getenv(CurrencyIDs)

	// TODO: check on empty
	//values := reflect.ValueOf(env)
	//for i := 0; i < values.NumField(); i++ {
	//	if values.Field(i).String() == "" {
	//		log.Fatalln("Env value is empty")
	//	}
	//}

	return env
}

func (env *ENV) GetTableIDs() []string {
	return strings.SplitN(env.TableIDs, ",", -1)
}

func (env *ENV) GetCurrencyIDs() []string {
	return strings.SplitN(env.CurrencyIDs, ",", -1)
}
