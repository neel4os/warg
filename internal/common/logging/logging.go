package logging

import (
	"github.com/rs/zerolog"
)

func SetLogConfig(isDebugLog bool) {
	if isDebugLog {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	// any other customization go here
}
