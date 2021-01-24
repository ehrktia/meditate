package logging

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

// NewLogger initializes a new instance of logger
func NewLogger() (*zap.SugaredLogger, error) {
	if logger == nil {
		prodLogger, err := zap.NewProduction()
		if err != nil {
			return nil, err
		}
		return prodLogger.Sugar(), nil
	}
	return logger, nil
}
