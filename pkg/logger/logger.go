package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Init() {
	config := DefaultConfig()

	var err error
	Logger, err = config.Build()
	if err != nil {
		panic("Failed create logger: " + err.Error())
	}
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Sync() {
	Logger.Sync()
}
