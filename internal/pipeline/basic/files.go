// Package basic ...
package basic

import (
	"net/http"
	"strings"

	"github.com/koha90/platform/internal/config"
	"github.com/koha90/platform/internal/pipeline"
	"github.com/koha90/platform/internal/services"
)

// StaticFileComponent ...
type StaticFileComponent struct {
	urlPrefix     string
	stdLibHandler http.Handler
}

// Init ...
func (sfc *StaticFileComponent) Init() {
	var cfg config.Configuration
	services.GetService(&cfg)
	sfc.urlPrefix = cfg.GetStringDefault("file:urlprefix", "/file/")
	path, ok := cfg.GetString("files:path")
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
