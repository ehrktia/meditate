package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_new_logger(t *testing.T) {
	got, err := NewLogger()
	assert.Nil(t, err)
	assert.NotNil(t, got)
	expectedType := &zap.SugaredLogger{}
	t.Run("should be able to use the logger", func(t *testing.T) {
		got.Info("logger initializes")
		assert.IsType(t, expectedType, got)
	})
	t.Run("should return existing instance of logger", func(t *testing.T) {
		logger = expectedType
		got, err := NewLogger()
		assert.Nil(t, err)
		assert.NotNil(t, got)
	})
}
