// Package main is general package for platform
package main

import (
	"github.com/koha90/platform/internal/config"
	"github.com/koha90/platform/internal/services"
	"github.com/koha90/platform/pkg/logging"
)

// writeMessage write "Hi, Platform" in terminal
func writeMessage(logger logging.Logger, cfg config.Configuration) {
	section, ok := cfg.GetSection("main")
	if ok {
		message, ok := section.GetString("message")
		if ok {
			logger.Info(message)
		} else {
			logger.Panic("Cannot find configuration setting")
		}
	} else {
		logger.Panic("Config section not found")
	}
}

func main() {
	services.RegistrationDefaultServices()

	var cfg config.Configuration
	services.GetService(&cfg)

	var logger logging.Logger
	services.GetService(&logger)

	writeMessage(logger, cfg)
}
