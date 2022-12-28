package config

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
	//values := reflect.ValueOf(.config)
	//for i := 0; i < values.NumField(); i++ {
	//	if values.Field(i).String() == "" {
	//		log.Fatalln("Env value is empty")
	//	}
	//}

	return env
}

func (env *ENV) GetTableIDs() []string {
	return strings.Split(strCleanUp(env.TableIDs), ",")
}

func (env *ENV) GetCurrencyIDs() []string {
	//commas := strings.Count(config.CurrencyIDs, ",")
	return strings.Split(strCleanUp(env.CurrencyIDs), ",")
}

func (env *ENV) GetCasinoID() string {
	return strCleanUp(env.CasinoID)
}

// strCleanUp removes all the extra characters added by different OSs environments.
func strCleanUp(strToCleanUp string) string {
	var builder strings.Builder
	for _, char := range strToCleanUp {
		if (char == ',') || (char >= 65 && char <= 90) || (char >= 97 && char <= 122) || (char >= 48 && char <= 57) {
			builder.WriteRune(char)
		}
	}

	return builder.String()
}
