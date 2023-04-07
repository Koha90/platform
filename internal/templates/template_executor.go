// Package templates ...
package templates

import "io"

// TemplateExecutor ...
type TemplateExecutor interface {
	ExecTemplate(writer io.Writer, name string, data interface{}) (err error)
	ExecTemplateWithFunc(
		writer io.Writer,
		name string,
		data interface{},
		handlerFunc InvokeHandleFunc,
	) (err error)
}

type InvokeHandleFunc func(handlerName string, methodName string, args ...interface{}) interface{}
