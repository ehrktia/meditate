package logging

import (
	"go.uber.org/zap"
)

//go:generate mockgen -package=mocks -source=${GOFILE} -destination=mocks/${GOFILE}
// Logger interface to create logger
type Logger interface {
	With(args ...interface{}) *zap.SugaredLogger
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	DPanic(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	DPanicf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
}
type PkgLogger struct {
	Logger
}

// NewLogger initializes a new instance of logger
func NewLogger(logger Logger) (*PkgLogger, error) {
	return &PkgLogger{
		logger,
	}, nil
}
