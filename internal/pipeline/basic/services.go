// Package basic ...
package basic

import (
	"github.com/koha90/platform/internal/pipeline"
	"github.com/koha90/platform/internal/services"
)

// ServicesComponent ...
type ServicesComponent struct{}

// Init ...
func (c *ServicesComponent) Init() {}

// ProcessRequest this component of middleware changed Context
func (c *ServicesComponent) ProcessRequest(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	reqContext := ctx.Request.Context()
	ctx.Request.WithContext(services.NewServiceContext(reqContext))
	next(ctx)
}
