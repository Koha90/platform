// Package actionresult ...
package actionresults

import (
	"context"
	"net/http"
)

// ActionContext ...
type ActionContext struct {
	context.Context
	http.ResponseWriter
}

// ActionResult ...
type ActionResult interface {
	Execute(*ActionContext) error
}
