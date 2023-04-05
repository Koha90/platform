// Package basic ...
package basic

import (
	"net/http"

	"github.com/koha90/platform/internal/pipeline"
	"github.com/koha90/platform/pkg/logging"
)

// LoggingResponseWritter ...
type LoggingResponseWritter struct {
	statusCode int
	http.ResponseWriter
}

// WriteHeader ...
func (w *LoggingResponseWritter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Write ...
func (w *LoggingResponseWritter) Write(b []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = http.StatusOK
	}
	return w.ResponseWriter.Write(b)
}

// LoggingComponent ...
type LoggingComponent struct{}

// ImplementsProcessRequestWithServices ...
func (lc *LoggingComponent) ImplementsProcessRequestWithServices() {}

// Init ...
func (lc *LoggingComponent) Init() {}

// ProcessRequestWithServices ...
func (lc *LoggingComponent) ProcessRequestWithServices(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	logger logging.Logger,
) {
	loggingWriter := LoggingResponseWritter{0, ctx.ResponseWriter}
	ctx.ResponseWriter = &loggingWriter

	logger.Infof("REQ --- %v - %v", ctx.Request.Method, ctx.Request.URL)
	next(ctx)

	logger.Infof("RSP %v %v", loggingWriter.statusCode, ctx.Request.URL)
}
