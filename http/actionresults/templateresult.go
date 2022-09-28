package actionresults

import "platform/templates"

func NewTemplateAction(name string, data interface{}) ActionResult {
	return &TemplateActionResult{templateName: name, data: data}
}

type TemplateActionResult struct {
	templateName string
	data         interface{}
	templates.TemplateExecutor
}

func (action *TemplateActionResult) Execute(ctx *ActionContext) error {
	return action.TemplateExecutor.ExecTemplate(ctx.ResponseWriter, action.templateName, action.data)
}
