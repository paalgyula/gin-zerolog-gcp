package gcp

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// SetupLogger sets up the logger for GKE log friendly format.
func SetupLogger(debug bool) {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	if debug {
		// nolint:exhaustivestruct
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
		log.Logger = log.With().Caller().Logger()

		return
	}

	log.Logger = log.Hook(SeverityHook{})
	log.Output(os.Stdout)
}
