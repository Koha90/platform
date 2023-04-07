// Package actionresults ...
package actionresults

import "encoding/json"

// NewJsonActionResult ...
func NewJsonActionResult(data interface{}) ActionResult {
	return &JsonActionResult{data: data}
}

// JsonActionResult ...
type JsonActionResult struct {
	data interface{}
}

func (action *JsonActionResult) Execute(ctx *ActionContext) error {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.ResponseWriter)
	return encoder.Encode(action.data)
}
