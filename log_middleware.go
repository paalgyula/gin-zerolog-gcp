package gcp

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// The following field name list have been composed to be StackTrace compatible.
// Not all fields are currently provided.
// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#HttpRequest
const (
	FieldNameRequest          = "httpRequest"
	FieldNameRequestMethod    = "requestMethod"
	FieldNameRequestURL       = "requestUrl"
	FieldNameRequestProtocol  = "protocol"
	FieldNameRequestUserAgent = "userAgent"
	FieldNameResponseSize     = "responseSize"
	FieldNameResponseStatus   = "status"
	FieldNameLatency          = "latency"
)

func Logger(c *gin.Context) zerolog.Logger {
	return c.MustGet("logger").(zerolog.Logger)
}

func WithAccessLog() gin.HandlerFunc {
	logContextRoot := log.With()

	return func(ctx *gin.Context) {
		start := time.Now()

		// Set logger in context.
		logger := logContextRoot.Logger()
		ctx.Set("logger", logger)

		ctx.Next()

		// Compose logger for access log entry with request and response data.
		var e *zerolog.Event

		if ctx.Writer.Status() < 400 {
			e = logger.Debug()
		} else {
			e = logger.Info()
		}

		r := ctx.Request

		rd := zerolog.Dict()
		rd = rd.Str(FieldNameRequestMethod, r.Method)
		rd = rd.Str(FieldNameRequestURL, r.URL.String())
		rd = rd.Str(FieldNameRequestUserAgent, r.UserAgent())
		rd = rd.Int(FieldNameResponseStatus, ctx.Writer.Status())
		rd = rd.Str(FieldNameResponseSize, strconv.Itoa(ctx.Writer.Size()))
		rd = rd.Str(FieldNameLatency, fmt.Sprintf("%dms", time.Since(start).Milliseconds()))
		rd = rd.Str(FieldNameRequestProtocol, r.Proto)

		// Log access.
		e.Dict(FieldNameRequest, rd).Msgf("Served HTTP request: %s", r.URL.String())
	}
}
