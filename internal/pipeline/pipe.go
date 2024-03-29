// Package pipeline ...
package pipeline

import (
	"net/http"
	"reflect"

	"github.com/koha90/platform/internal/services"
)

// RequestPipeline ...
type RequestPipeline func(*ComponentContext)

var emptyPipeline RequestPipeline = func(cc *ComponentContext) {
	/* do nothing */
}

// CreatePipeline ...
func CreatePipeline(components ...interface{}) RequestPipeline {
	f := emptyPipeline
	for i := len(components) - 1; i >= 0; i-- {
		currentComponent := components[i]
		services.Populate(currentComponent)
		nextFunc := f
		if servComp, ok := currentComponent.(ServiceMiddleWareComponent); ok {
			f = createServiceDependentFunction(currentComponent, nextFunc)
			servComp.Init()
		} else if stdComp, ok := currentComponent.(MiddlewareComponent); ok {
			f = func(context *ComponentContext) {
				if context.error == nil {
					stdComp.ProcessRequest(context, nextFunc)
				}
			}
			stdComp.Init()
		} else {
			panic("Value is not a middleware component")
		}
	}

	return f
}

func createServiceDependentFunction(
	component interface{},
	nextFunc RequestPipeline,
) RequestPipeline {
	method := reflect.ValueOf(component).MethodByName("ProcessRequestWithServices")
	if method.IsValid() {
		return func(context *ComponentContext) {
			if context.error == nil {
				_, err := services.CallForContext(
					context.Request.Context(),
					method.Interface(),
					context,
					nextFunc,
				)
				if err != nil {
					context.Error(err)
				}
			}
		}
	}
	panic("No ProcessRequestWithServices method defined")
}

// ProcessRequest ...
func (pl RequestPipeline) ProcessRequest(req *http.Request, resp http.ResponseWriter) error {
	ctx := ComponentContext{
		Request:        req,
		ResponseWriter: resp,
	}

	pl(&ctx)

	return ctx.error
}
