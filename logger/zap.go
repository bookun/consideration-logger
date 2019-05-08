package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Info(tag string, fields map[string]string) {
	var zapStrings []zapcore.Field
	for k, v := range fields {
		zapStrings = append(zapStrings, zap.String(k, v))
	}
	Logger.Info(tag, zapStrings...)
}

func Error(tag string, fields map[string]string) {
	var zapStrings []zapcore.Field
	for k, v := range fields {
		zapStrings = append(zapStrings, zap.String(k, v))
	}
	Logger.Error(tag, zapStrings...)
}
