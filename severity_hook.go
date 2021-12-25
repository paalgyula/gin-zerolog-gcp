package gcp

import (
	"strings"

	"github.com/rs/zerolog"
)

type SeverityHook struct{}

func (h SeverityHook) Run(e *zerolog.Event, level zerolog.Level, _ string) {
	if level != zerolog.NoLevel {
		e.Str("severity", strings.ToUpper(level.String()))
	}
}
