// Package logging to write log
package logging

// LogLevel ...
type LogLevel int

// Trace, Debug, Information, Warning, Fatal, None is level of log
const (
	Trace LogLevel = iota
	Debug
	Information
	Warning
	Fatal
	None
)

// Logger ...
type Logger interface {
	Trace(string)
	Tracef(string, ...interface{})

	Debug(string)
	Debugf(string, ...interface{})

	Info(string)
	Infof(string, ...interface{})

	Warn(string)
	Warnf(string, ...interface{})

	Panic(string)
	Panicf(string, ...interface{})
}
