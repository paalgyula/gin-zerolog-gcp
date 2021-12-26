package gcp

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type loggerConfig struct {
	pretty   bool
	logLevel zerolog.Level
}

// LoggerOption function definition for configuring GCP logger
type LoggerOption = func(*loggerConfig)

// You can turn on pretty human readable format instead of json output
func WithPrettyOutput() LoggerOption {
	return func(opt *loggerConfig) {
		opt.pretty = true
	}
}

// WithLogLevel you can specify the logging level
func WithLogLevel(lvl zerolog.Level) LoggerOption {
	return func(opt *loggerConfig) {
		opt.logLevel = lvl
	}
}

// SetupLogger sets up the logger for GKE log friendly format.
func SetupLogger(opts ...LoggerOption) {
	opt := &loggerConfig{
		pretty:   false,
		logLevel: zerolog.InfoLevel,
	}

	for _, o := range opts {
		o(opt)
	}

	zerolog.SetGlobalLevel(opt.logLevel)

	if opt.pretty {

		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    false,
			TimeFormat: time.Kitchen,
		})

		log.Logger = log.With().
			Caller().
			Logger()

		return
	}

	log.Logger = log.Hook(SeverityHook{})
	log.Output(os.Stdout)
}
