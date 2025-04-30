package utils

import (
	"context"

	"go.uber.org/zap"
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
