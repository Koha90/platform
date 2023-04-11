// Package handling
package handling

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"

	"github.com/koha90/platform/internal/http/actionresults"
	"github.com/koha90/platform/internal/services"
)

// AddMethodAlias ...
func (rc *RouterComponent) AddMethodAlias(
	srcURL string,
	method interface{},
	data ...interface{},
) *RouterComponent {
	var urlgen URLGenerator
	services.GetService(&urlgen)
	url, err := urlgen.GenerateURL(method, data...)
	if err == nil {
		return rc.AddURLAlias(srcURL, url)
	}
	panic(err)
}

// AddURLAlias ...
func (rc *RouterComponent) AddURLAlias(srcURL string, targetURL string) *RouterComponent {
	aliasFunc := func(interface{}) actionresults.ActionResult {
		return actionresults.NewRedirectAction(targetURL)
	}

	alias := Route{
		httpMethod:  http.MethodGet,
		handlerName: "Alias",
		actionName:  "Redirect",
		expression:  *regexp.MustCompile(fmt.Sprintf("^%v[/]?$", srcURL)),
		handlerMethod: reflect.Method{
			Type: reflect.TypeOf(aliasFunc),
			Func: reflect.ValueOf(aliasFunc),
		},
	}

	rc.routes = append([]Route{alias}, rc.routes...)
	return rc
}
