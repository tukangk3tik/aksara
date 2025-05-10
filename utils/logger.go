package utils

import (
	"context"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var GlobalLogger *zap.Logger

type ctxKeyLogger struct{}

func WithContext(ctx context.Context, baseLogger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxKeyLogger{}, baseLogger)
}

func FromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(ctxKeyLogger{}).(*zap.Logger)
    if ok && logger != nil {
        return logger
    }
    if GlobalLogger != nil {
        return GlobalLogger
    }
    return zap.NewNop()
}

func SetupLogger(config Config) {
    if config.AppEnv == "test" {
		GlobalLogger = zap.NewNop()
		return
	}

	logFile, _ := os.OpenFile(config.AppLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// Configure encoder (JSON format)
	var core zapcore.Core
	encoderConfig := zap.NewProductionEncoderConfig()

	// Create core for file logging
	if config.AppEnv == "production" {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // Use JSON format
			zapcore.AddSync(multiWriter),          // Output to file
			zap.ErrorLevel,                        // Log level
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // Use JSON format
			zapcore.AddSync(multiWriter),          // Output to file
			zap.DebugLevel,                        // Log level
		)
	}
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	GlobalLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
	defer GlobalLogger.Sync() // Flush logs before exiting
}