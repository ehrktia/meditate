package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func Test_new_logger(t *testing.T) {
    got,err:=NewLogger()
    assert.Nil(t,err)
    assert.NotNil(t, got)
    t.Run("should be able to use the logger", func(t *testing.T) {
        got.Info("logger initializes")
    })
}
