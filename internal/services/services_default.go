// Package services ...
package services

import (
	"github.com/koha90/platform/internal/config"
	"github.com/koha90/platform/pkg/logging"
)

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
}
