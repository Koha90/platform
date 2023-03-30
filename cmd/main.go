// package main is general package for platform
package main

import (
	"github.com/koha90/platform/pkg/logging"
)

// writeMessage write "Hi, Platform" in terminal
func writeMessage(logger logging.Logger) {
	logger.Info("Hello!")
}

func main() {
	var logger logging.Logger = logging.NewDefaultLogger(logging.Information)

	writeMessage(logger)
}
