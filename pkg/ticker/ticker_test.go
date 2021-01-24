package ticker

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/buffer"
)

func TestCreateticker(t *testing.T) {
	sessionDuration := "1m"
	interimDuration := "30s"
	errDuration := "3m"
	if err := os.Setenv(defaultInterimInterval, interimDuration); err != nil {
		t.Error(err)
	}
	if err := os.Setenv(defaultSessionDuration, sessionDuration); err != nil {
		t.Error(err)
	}
	got, err := IntialiseTimer()
	t.Run("should create timer", func(t *testing.T) {
		assert.Equal(t, sessionDuration, got.sessionDuration)
		assert.Nil(t, err)
	})
	if err := os.Unsetenv(defaultSessionDuration); err != nil {
		t.Error(err)
	}
	if err := os.Unsetenv(defaultInterimInterval); err != nil {
		t.Error(err)
	}
	t.Run("should return error when sessionDuration is lower than defaultInterimInterval", func(t *testing.T) {
		if err := os.Setenv(defaultInterimInterval, errDuration); err != nil {
			t.Error(err)
		}
		_, err := IntialiseTimer()
		assert.NotNil(t, err)
	})
	t.Cleanup(func() {
		if err := os.Unsetenv(defaultSessionDuration); err != nil {
			t.Error(err)
		}
		if err := os.Unsetenv(defaultInterimInterval); err != nil {
			t.Error(err)
		}
	})
}
func TestInterimTicker(t *testing.T) {
	if err := os.Setenv(defaultInterimInterval, "5s"); err != nil {
		t.Error(err)
	}
	if err := os.Setenv(defaultSessionDuration, "10s"); err != nil {
		t.Error(err)
	}
	timer, err := IntialiseTimer()
	assert.Nil(t, err)
	status := make(chan bool)
	t.Run("should start a ticker for the interim interval period set", func(t *testing.T) {
		go InitiateInterimTimer(timer, status)
		select {
		case <-status:
			t.Log("completed")
		case <-time.After(6 * time.Second):
			t.Error("timer ran out")
		}
	})
	t.Run("should provide no of interim timers for session", func(t *testing.T) {
		got := CountInterimTimers(timer)
		if got != 2 {
			t.Errorf("expected: %v got: %v", 2, got)
		}
	})
	buf := new(buffer.Buffer)
	t.Run("should run countdown for each sec for one interim timer", func(t *testing.T) {
		StartInterimTimer(buf, timer, status)
		if strings.Count(buf.String(), "\n") < 6 {
			t.Error("expected 6 lines in output")
		}
	})
	t.Cleanup(func() {
		if err := os.Unsetenv(defaultInterimInterval); err != nil {
			t.Error(err)
		}
	})
}
