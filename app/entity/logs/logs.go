package logs

import (
	"context"
	"errors"
	"time"
)

var (
	TankErrorLoggerAlreadyInitialized = errors.New("Log gateway already initialized")
)

type ILogger interface {
	Context(ctx context.Context) ILogger
	Error(message string) time.Time
	Fatal(message string) time.Time
	Info(message string) time.Time
}

var loggerGateway ILogger = nil
var emptyLogger ILogger = &_empty{}

func SetLogger(logger ILogger) (err error) {
	if loggerGateway == nil {
		loggerGateway = logger
	}

	return TankErrorLoggerAlreadyInitialized
}

func Gateway() ILogger {
	if loggerGateway != nil {
		return loggerGateway
	}
	return emptyLogger
}

type _empty struct {
}

func (logger *_empty) Context(ctx context.Context) ILogger {
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
