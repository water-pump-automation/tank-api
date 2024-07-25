package stdout

import (
	"context"
	"fmt"
	"time"
	"water-tank-api/app/core/entity/logs"
)

type STDOutLogger struct {
	infoPrefix  string
	fatalPrefix string
	errorPrefix string
}

func NewSTDOutLogger() *STDOutLogger {
	return &STDOutLogger{
		infoPrefix:  "INFO",
		errorPrefix: "ERROR",
		fatalPrefix: "FATAL",
	}
}

func (logger *STDOutLogger) Context(ctx context.Context) logs.ILogger {
	return &STDOutLogger{}
}

func (logger *STDOutLogger) Error(message string) time.Time {
	now := time.Now()
	fmt.Printf("[%s] %s\n", logger.errorPrefix, message)
	return now
}

func (logger *STDOutLogger) Fatal(message string) time.Time {
	now := time.Now()
	fmt.Printf("[%s] %s\n", logger.fatalPrefix, message)
	return now
}

func (logger *STDOutLogger) Info(message string) time.Time {
	now := time.Now()
	fmt.Printf("[%s] %s\n", logger.infoPrefix, message)
	return now
}
