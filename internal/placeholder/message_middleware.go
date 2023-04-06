// Package placeholder generate response
package placeholder

import (
	"github.com/koha90/platform/internal/config"
	"github.com/koha90/platform/internal/pipeline"
	"github.com/koha90/platform/internal/templates"
)

// SimpleMessageComponent ...
type SimpleMessageComponent struct {
	Message string
	config.Configuration
}

// ImplementsProcessRequestWithServices ...
func (c *SimpleMessageComponent) ImplementsProcessRequestWithServices() {}

// Init ...
func (c *SimpleMessageComponent) Init() {
	c.Message = c.Configuration.GetStringDefault("main:message",
		"Default Message")
}

// ProcessRequestWithServices ...
func (c *SimpleMessageComponent) ProcessRequestWithServices(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	executor templates.TemplateExecutor,
) {
	err := executor.ExecTemplate(ctx.ResponseWriter, "simple_message.html", c.Message)
	if err != nil {
		ctx.Error(err)
	} else {
		next(ctx)
	}
}
