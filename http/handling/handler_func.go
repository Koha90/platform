package handling

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"strings"

	"platform/http/actionresults"
	"platform/services"
	"platform/templates"
)

func createInvokeHandlerFunc(ctx context.Context, routes []Route) templates.InvokeHandlerFunc {
	return func(handlerName string, methodName string, args ...interface{}) interface{} {
		var err error
		for _, route := range routes {
			if strings.EqualFold(handlerName, route.handlerName) &&
				strings.EqualFold(methodName, route.handlerMethod.Name) {
				paramsVals := make([]reflect.Value, len(args))
				for i := 0; i < len(args); i++ {
					paramsVals[i] = reflect.ValueOf(args[i])
				}
				structVal := reflect.New(route.handlerMethod.Type.In(0))
				services.PopulateForContext(ctx, structVal.Interface())
				paramsVals = append([]reflect.Value{structVal.Elem()},
					paramsVals...)
				result := route.handlerMethod.Func.Call(paramsVals)
				if action, ok := result[0].Interface().(*actionresults.TemplateActionResult); ok {
					invoker := createInvokeHandlerFunc(ctx, routes)
					err = services.PopulateForContextWithExtras(ctx, action, map[reflect.Type]reflect.Value{
						reflect.TypeOf(invoker): reflect.ValueOf(invoker),
					})
					writer := &stringResponseWriter{Builder: &strings.Builder{}}
					if err == nil {
						err = action.Execute(&actionresults.ActionContext{
							Context:        ctx,
							ResponseWriter: writer,
						})
						if err == nil {
							return (template.HTML)(writer.Builder.String())
						}
					}
				} else {
					return fmt.Sprint(result[0])
				}
			}
		}
		if err == nil {
			err = fmt.Errorf("no route found for %v %v", handlerName, methodName)
		}
		panic(err)
	}
}

type stringResponseWriter struct {
	*strings.Builder
}

func (sw *stringResponseWriter) Write(data []byte) (int, error) {
	return sw.Builder.Write(data)
}

func (sw *stringResponseWriter) WriteHeader(statusCode int) {}

func (sw *stringResponseWriter) Header() http.Header {
	return http.Header{}
}
