// Package placeholder generate response
package placeholder

import (
	"errors"
	"io"

	"github.com/koha90/platform/internal/config"
	"github.com/koha90/platform/internal/pipeline"
	"github.com/koha90/platform/internal/services"
)

type SimpleMessageComponent struct{}

// Init ...
func (c *SimpleMessageComponent) Init() {}

func (c *SimpleMessageComponent) ProcessRequest(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	var cfg config.Configuration
	services.GetService(&cfg)

	msg, ok := cfg.GetString("main:message")
	if ok {
		io.WriteString(ctx.ResponseWriter, msg)
	} else {
		ctx.Error(errors.New("Cannot find config setting"))
	}

	next(ctx)
}
