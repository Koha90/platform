// Package templates
package templates

import (
	"errors"
	"html/template"
	"sync"

	"github.com/koha90/platform/internal/config"
)

var (
	once      = sync.Once{}
	templates *template.Template
)

// LoadTemplates ...
func LoadTemplates(c config.Configuration) (err error) {
	path, ok := c.GetString("temlates:path")
	if !ok {
		return errors.New("Cannot load template config")
	}

	reload := c.GetBoolDefault("templates:reload", false)

	once.Do(func() {
		doLoad := func() (t *template.Template) {
			t = template.New("htmlTemplates")
			t.Funcs(map[string]interface{}{
				"body":   func() string { return "" },
				"layout": func() string { return "" },
			})
			t, err = t.ParseGlob(path)
			return t
		}
		if reload {
			getTemplates = doLoad
		} else {
			// var templates *template.Template
			templates = doLoad()
			getTemplates = func() *template.Template {
				t, _ := templates.Clone()
				return t
			}
		}
	})
	return
}
