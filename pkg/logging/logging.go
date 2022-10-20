package logging

import (
	"context"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Key is generic key used for ctx value load
type Key string

var (
	once = &sync.Once{}
	// GlobalLogger is the zap logger used across application
	GlobalLogger *zap.Logger
)

type config struct {
	loglevel string
}

// New generates a singleton instance of logger
func New() *zap.Logger {
	c := newCoreConfig("Debug")
	core := zapcore.NewCore(c, os.Stderr, zap.DebugLevel)
	once.Do(func() {
		GlobalLogger = zap.New(zapcore.NewCore())
	})
	return GlobalLogger
}
func newCoreConfig(level string) zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       level,
		CallerKey:      "log",
		MessageKey:     "msg",
		StacktraceKey:  "stack-trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
	}
	zap.NewProductionEncoderConfig()

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
