// Package logging ...
package logging

import (
	"fmt"
	"log"
)

// DefaultLogger ...
type DefaultLogger struct {
	minLevel     LogLevel
	loggers      map[LogLevel]*log.Logger
	triggerPanic bool
}

// MinLogLevel ...
func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

func (l *DefaultLogger) write(level LogLevel, message string) {
	if l.minLevel <= level {
		l.loggers[level].Output(2, message)
	}
}

// Trace ...
func (l *DefaultLogger) Trace(msg string) {
	l.write(Trace, msg)
}

// Tracef ...
func (l *DefaultLogger) Tracef(template string, vals ...interface{}) {
	l.write(Trace, fmt.Sprintf(template, vals...))
}

// Debug ...
func (l *DefaultLogger) Debug(msg string) {
	l.write(Debug, msg)
}

// Debugf ...
func (l *DefaultLogger) Debugf(template string, vals ...interface{}) {
	l.write(Debug, fmt.Sprintf(template, vals...))
}

// Info ...
func (l *DefaultLogger) Info(msg string) {
	l.write(Information, msg)
}

// Infof ...
func (l *DefaultLogger) Infof(template string, vals ...interface{}) {
	l.write(Information, fmt.Sprintf(template, vals...))
}

// Warn log level
func (l *DefaultLogger) Warn(msg string) {
	l.write(Warning, msg)
}

// Warnf log level
func (l *DefaultLogger) Warnf(template string, vals ...interface{}) {
	l.write(Warning, fmt.Sprintf(template, vals...))
}

// Panic log level
func (l *DefaultLogger) Panic(msg string) {
	l.write(Fatal, msg)
	if l.triggerPanic {
		panic(msg)
	}
}

// Panicf ...
func (l *DefaultLogger) Panicf(template string, vals ...interface{}) {
	formattedMsg := fmt.Sprintf(template, vals...)
	l.write(Fatal, formattedMsg)
	if l.triggerPanic {
		panic(formattedMsg)
	}
}
