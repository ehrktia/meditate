package logging

import (
	"context"
	"sync"

	"go.uber.org/zap"
)

// Key is generic key used for ctx value load
type Key string

var (
	once = &sync.Once{}
	// GlobalLogger is the zap logger used across application
	GlobalLogger *zap.Logger
)

// New generates a singleton instance of logger
func New() *zap.Logger {
	once.Do(func() {
		GlobalLogger = zap.New(nil)
	})
	return GlobalLogger
}

// LogToCtx loads logger to context
func LogToCtx(ctx context.Context, log *zap.Logger) context.Context {
	ctxVal := context.WithValue(ctx, Key("LOG"), log)
	return ctxVal
}

// LogFromCtx gives an instance of zap logger to caller
func LogFromCtx(ctx context.Context) *zap.Logger {
	v := ctx.Value(Key("LOG"))
	// when missing logger in ctx load a new logger
	if v == nil {
		ctx = LogToCtx(ctx, New())
		v := ctx.Value(Key("LOG"))
		return v.(*zap.Logger)
	}
	return v.(*zap.Logger)
}
