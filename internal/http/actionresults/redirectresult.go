// Package actionresults ...
package actionresults

import "net/http"

// NewRedirectAction ...
func NewRedirectAction(url string) ActionResult {
	return &RedirectActionResult{url: url}
}

// RedirectActionResult ...
type RedirectActionResult struct {
	url string
}

func (action *RedirectActionResult) Execute(ctx *ActionContext) error {
	ctx.ResponseWriter.Header().Set("Location", action.url)
	ctx.ResponseWriter.WriteHeader(http.StatusSeeOther)
	return nil
}
