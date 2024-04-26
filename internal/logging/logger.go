package logging

import "go.uber.org/zap"

var zapLogger *zap.Logger

func Setup() error {
	var err error
	zapLogger, err = zap.NewProduction()
	return err
}

func Debug(message string, fields ...zap.Field) {
	zapLogger.Debug(message, fields...)
}

func Info(message string, fields ...zap.Field) {
	zapLogger.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	zapLogger.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	zapLogger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	zapLogger.Fatal(message, fields...)
}
