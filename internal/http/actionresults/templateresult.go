// Package actionresults ...
package actionresults

import "github.com/koha90/platform/internal/templates"

// NewTemplateAction ...
func NewTemplateAction(name string, data interface{}) ActionResult {
	return &TemplateActionResult{templateName: name, data: data}
}

// TemplateActionResult ...
type TemplateActionResult struct {
	templateName string
	data         interface{}
	templates.TemplateExecutor
	templates.InvokeHandleFunc
}

// Execute ...
func (action *TemplateActionResult) Execute(ctx *ActionContext) error {
	return action.TemplateExecutor.ExecTemplateWithFunc(
		ctx.ResponseWriter,
		action.templateName,
		action.data,
		action.InvokeHandleFunc,
	)
}
