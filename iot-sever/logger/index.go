package logger

import (
	"go.uber.org/zap"
	"log"
)

var Logger *zap.Logger

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf(err.Error())
	}
	Logger = logger
}
