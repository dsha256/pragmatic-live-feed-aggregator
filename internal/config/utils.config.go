package config

import (
	"log"
	"strings"
)

const (
	UTF8CodeOfComma     = 44
	UTF8CodeOfDigitZero = 48
	UTF8CodeOfDigitNine = 57
	UTF8CodeOfUpperA    = 65
	UTF8CodeOfUpperZ    = 90
	UTF8CodeOfLowerA    = 97
	UTF8CodeOfLowerZ    = 122
)

// strCleanUp removes all the extra characters added by different OSs environments.
func strCleanUp(strToCleanUp string) string {
	var builder strings.Builder
	for _, char := range strToCleanUp {
		if (char == UTF8CodeOfComma) ||
			(char >= UTF8CodeOfDigitZero && char <= UTF8CodeOfDigitNine) ||
			(char >= UTF8CodeOfUpperA && char <= UTF8CodeOfUpperZ) ||
			(char >= UTF8CodeOfLowerA && char <= UTF8CodeOfLowerZ) {
			builder.WriteRune(char)
		}
	}

	return builder.String()
}

func checkEnvVarOnEmptiness(name string, value any) {
	switch value.(type) {
	case int, int8, int16, int32, uint, uint8, uint16, uint32:
		if value == 0 {
			log.Fatalf("Zero value env var: %s", name)
		}
	case string:
		if value == "" {
			log.Fatalf("Zero value env var: %s", name)
		}
	case any:
		if value == nil {
			log.Fatalf("Zero value env var: %s", name)
		}
	}
}
