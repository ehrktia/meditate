package ticker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)
func TestCreateticker(t *testing.T) {
        got:=IntialiseTimer(time.Minute,30*time.Second)
    t.Run("should create timer", func(t *testing.T) {
        assert.NotNil(t, got)
        assert.Equal(t, time.Minute, got.tickerDuration)
    })
    t.Run("should get the ticker initialised", func(t *testing.T) {
        assert.NotNil(t, got.GetTicker())
    })
    t.Run("should set interval timer for provided interval", func(t *testing.T) {
        assert.NotNil(t, got.GetIntervalTimer())
    })
}
