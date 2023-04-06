// Package main is general package for platform
package main

import (
	"github.com/koha90/platform/internal/placeholder"
	"github.com/koha90/platform/internal/services"
)

func main() {
	services.RegistrationDefaultServices()
	placeholder.Start()
}
