// Package actionresults ...
package actionresults

// NewErrorAction ...
func NewErrorAction(err error) ActionResult {
	return &ErrorActionResult{err}
}

// ErrorActionResult ...
type ErrorActionResult struct {
	error
}

// Execute error
func (action *ErrorActionResult) Execute(*ActionContext) error {
	return action.error
}
