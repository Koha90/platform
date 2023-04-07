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
}

func (action *TemplateActionResult) Execute(ctx *ActionContext) error {
	return action.TemplateExecutor.ExecTemplate(
		ctx.ResponseWriter,
		action.templateName,
		action.data,
	)
}
