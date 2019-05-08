package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func Info(tag, label, message string) {
	Logger.Info(tag, zap.String(label, message))
}

func Error(tag, label, message string) {
	Logger.Error(tag, zap.String(label, message))
}
