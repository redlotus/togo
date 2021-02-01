package main

import (
	"os"

	"go.uber.org/zap"
)

const DebugFlag = "DEBUG_MODE"

func initLogger() *zap.SugaredLogger {
	debugFlag := os.Getenv(DebugFlag)
	var logger *zap.Logger
	if debugFlag == "true" {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}

	sugar := logger.Sugar()

	return sugar
}
