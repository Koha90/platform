// Package services ...
package services

import (
	"github.com/koha90/platform/internal/config"
	"github.com/koha90/platform/internal/templates"
	"github.com/koha90/platform/pkg/logging"
)

// RegistrationDefaultServices ...
func RegistrationDefaultServices() {
	err := AddSingletone(func() (c config.Configuration) {
		c, loadErr := config.Load("config.json")
		if loadErr != nil {
			panic(loadErr)
		}

		return
	})

	err = AddSingletone(func(appconfig config.Configuration) logging.Logger {
		return logging.NewDefaultLogger(appconfig)
	})
	if err != nil {
		panic(err)
	}

	err = AddSingletone(
		func(c config.Configuration) templates.TemplateExecutor {
			templates.LoadTemplates(c)
			return &templates.LayoutTemplateProcessor{}
		})
	if err != nil {
		panic(err)
	}
}
