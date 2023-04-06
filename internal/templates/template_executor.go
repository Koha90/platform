// Package templates ...
package templates

import "io"

// TemplateExecutor ...
type TemplateExecutor interface {
	ExecTemplate(writer io.Writer, name string, data interface{}) (err error)
}
