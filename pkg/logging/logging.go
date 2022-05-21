package logging

import (
	"sync"

	"go.uber.org/zap"
)

var (
	once = &sync.Once{}
	// GlobalLogger is the zap logger used across application
	GlobalLogger *zap.Logger
)

// New generates a singleton instance of logger
func New() *zap.Logger {
	once.Do(func() {
		if GlobalLogger == nil {
			GlobalLogger = zap.New(nil)
		}
	})
	return GlobalLogger
}
