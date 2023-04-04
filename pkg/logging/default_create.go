// Package logging ...
package logging

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/koha90/platform/internal/config"
)

var (
	trace = fmt.Sprintf("%c[1;40;36m%s%c[0m", 0x1B, "[TRACE]: ", 0x1B)
	debug = fmt.Sprintf("%c[1;40;37m%s%c[0m", 0x1B, "[DEBUG]: ", 0x1B)
	info  = fmt.Sprintf("%c[1;40;34m%s%c[0m", 0x1B, "[INFO]: ", 0x1B)
	warn  = fmt.Sprintf("%c[1;40;33m%s%c[0m", 0x1B, "[WARN]: ", 0x1B)
	fatal = fmt.Sprintf("%c[1;40;31m%s%c[0m", 0x1B, "[FATAL]: ", 0x1B)
)

// NewDefaultLogger create default logger
func NewDefaultLogger(cfg config.Configuration) Logger {
	var level LogLevel = Debug
	if configLevelString, found := cfg.GetString("logging:level"); found {
		level = LogLevelFromString(configLevelString)
	}

	flags := log.Lmsgprefix | log.Ltime | log.Ldate
	return &DefaultLogger{
		minLevel: level,
		loggers: map[LogLevel]*log.Logger{
			Trace:       log.New(os.Stdout, trace, flags),
			Debug:       log.New(os.Stdout, debug, flags),
			Information: log.New(os.Stdout, info, flags),
			Warning:     log.New(os.Stdout, warn, flags),
			Fatal:       log.New(os.Stdout, fatal, flags),
		},
		triggerPanic: true,
	}
}

// LogLevelFromString set level log
func LogLevelFromString(val string) (level LogLevel) {
	switch strings.ToLower(val) {
	case "debug":
		level = Debug
	case "information":
		level = Information
	case "warning":
		level = Warning
	case "fatal":
		level = Fatal
	case "none":
		level = None
	default:
		level = Debug
	}

	return level
}
