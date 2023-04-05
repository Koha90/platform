// Package basic ...
package basic

import (
	"net/http"
	"strings"

	"github.com/koha90/platform/internal/config"
	"github.com/koha90/platform/internal/pipeline"
)

// StaticFileComponent ...
type StaticFileComponent struct {
	urlPrefix     string
	stdLibHandler http.Handler
	Config        config.Configuration
}

// Init ...
func (sfc *StaticFileComponent) Init() {
	sfc.urlPrefix = sfc.Config.GetStringDefault("file:urlprefix", "/files/")
	path, ok := sfc.Config.GetString("files:path")
	if ok {
		sfc.stdLibHandler = http.StripPrefix(sfc.urlPrefix, http.FileServer(http.Dir(path)))
	} else {
		panic("Cannot load file configuration settings")
	}
}

// ProcessRequest ...
func (sfc *StaticFileComponent) ProcessRequest(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	if !strings.EqualFold(ctx.Request.URL.Path, sfc.urlPrefix) &&
		strings.HasPrefix(ctx.Request.URL.Path, sfc.urlPrefix) {
		sfc.stdLibHandler.ServeHTTP(ctx.ResponseWriter, ctx.Request)
	} else {
		next(ctx)
	}
}
