// Package handling
package handling

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/koha90/platform/internal/http/handling/params"
	"github.com/koha90/platform/internal/pipeline"
	"github.com/koha90/platform/internal/services"
)

// RouterComponent ...
type RouterComponent struct {
	routes []Route
}

// NewRouter ...
func NewRouter(handlers ...HandlerEntry) *RouterComponent {
	return &RouterComponent{generateRoutes(handlers...)}
}

// Init ...
func (router *RouterComponent) Init() {}

// ProcessRequest ...
func (router *RouterComponent) ProcessRequest(
	context *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	for _, route := range router.routes {
		if strings.EqualFold(context.Request.Method, route.httpMethod) {
			matches := route.expression.FindAllStringSubmatch(context.URL.Path, -1)
			if len(matches) > 0 {
				rawParamVals := []string{}
				if len(matches[0]) > 1 {
					rawParamVals = matches[0][1:]
				}

				err := router.invokeHandler(route, rawParamVals, context)
				if err == nil {
					next(context)
				} else {
					context.Error(err)
				}
				return
			}
		}
	}

	context.ResponseWriter.WriteHeader(http.StatusNotFound)
}

func (router *RouterComponent) invokeHandler(
	route Route,
	rawParams []string,
	context *pipeline.ComponentContext,
) error {
	paramVals, err := params.GetParametrsFromRequest(
		context.Request,
		route.handlerMethod,
		rawParams,
	)
	if err == nil {
		structVal := reflect.New(route.handlerMethod.Type.In(0))
		services.PopulateForContext(context.Context(), structVal.Interface())
		paramVals = append([]reflect.Value{structVal.Elem()}, paramVals...)

		result := route.handlerMethod.Func.Call(paramVals)
		io.WriteString(context.ResponseWriter, fmt.Sprint(result[0].Interface()))
	}

	return err
}
