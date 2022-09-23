package services

import (
	"platform/config"
	"platform/logging"
)

func RegisterDefaultServices() {
	err := AddSingleton(func() (c config.Configuration) {
		c, loadErr := config.Load("config.json")
		if loadErr != nil {
			panic(loadErr)
		}
		return
	})

	err = AddSingleton(func(appConfig config.Configuration) logging.Logger {
		return logging.NewDefaultLogger(appConfig)
	})
	if err != nil {
		panic(err)
	}
}
