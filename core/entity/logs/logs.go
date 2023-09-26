package logs

import (
	"context"
	"time"
)

type Logger interface {
	Context(ctx context.Context) Logger
	Error(message string) time.Time
	Fatal(message string) time.Time
	Info(message string) time.Time
}

var loggerGateway Logger = nil
var emptyLogger Logger = &_empty{}

func SetLogger(logger Logger) (err error) {
	if loggerGateway == nil {
		loggerGateway = logger
	}

	return WaterTankErrorLoggerAlreadyInitialized
}

func Gateway() Logger {
	if loggerGateway != nil {
		return loggerGateway
	}
	return emptyLogger
}

type _empty struct {
}

func (logger *_empty) Context(ctx context.Context) Logger {
	return &_empty{}
}

func (logger *_empty) Error(message string) time.Time {
	return time.Now()
}

func (logger *_empty) Fatal(message string) time.Time {
	return time.Now()
}

func (logger *_empty) Info(message string) time.Time {
	return time.Now()
}
