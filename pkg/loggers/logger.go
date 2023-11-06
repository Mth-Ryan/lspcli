package loggers

import "fmt"

type Logger interface {
	Log(string, ...any)
}

type StdOutLogger struct{}

func NewStdOutLogger() *StdOutLogger {
	return &StdOutLogger{}
}

func (l *StdOutLogger) Log(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}

type QuietLogger struct{}

func NewQuietLogger() *QuietLogger {
	return &QuietLogger{}
}

func (l *QuietLogger) Log(format string, args ...any) {
}
