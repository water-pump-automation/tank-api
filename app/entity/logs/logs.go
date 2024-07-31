package logs

import (
	"context"
	"errors"
)

var (
	ErrTankLoggerAlreadyInitialized = errors.New("log gateway already initialized")
)

type ILogger interface {
	Context(ctx context.Context) ILogger
	Error(message string)
	Fatal(message string)
	Info(message string)
}

var loggerGateway ILogger = nil
var emptyLogger ILogger = &_empty{}

func SetLogger(logger ILogger) (err error) {
	if loggerGateway == nil {
		loggerGateway = logger
	}

	return ErrTankLoggerAlreadyInitialized
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

func (logger *_empty) Error(message string) {}

func (logger *_empty) Fatal(message string) {}

func (logger *_empty) Info(message string) {}
