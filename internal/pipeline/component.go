// Package pipeline middleware application
package pipeline

import "net/http"

// ComponentContext struct of middleware application
type ComponentContext struct {
	*http.Request
	http.ResponseWriter
	error
}

// Error ...
func (mwc *ComponentContext) Error(err error) {
	mwc.error = err
}

// GetErr ...
func (mwc *ComponentContext) GetErr() error {
	return mwc.error
}

// MiddlewareComponent interface of function
type MiddlewareComponent interface {
	Init()
	ProcessRequest(context *ComponentContext, next func(*ComponentContext))
}

// ServiceMiddleWareComponent ...
type ServiceMiddleWareComponent interface {
	Init()
	ImplementsProcessRequestWithServices()
}
