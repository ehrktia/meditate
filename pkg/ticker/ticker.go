package ticker

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	interimInterval       = "INTERIM_INTERVAL"
	defaultTickerInterval = 1 * time.Second
	sessionDuration       = "SESSION_DURATION"
)

// Timer holds session and interim timer.
type Timer struct {
	sessionDuration        time.Duration
	defaultInterimInterval time.Duration
}

// DefaultTimerEnv uses all default values from environment
func DefaultTimerFromEnv() (Timer, error) {
	defaultSession := os.Getenv(sessionDuration)
	duration, err := time.ParseDuration(defaultSession)
	if err != nil {
		return Timer{}, err
	}
	defaultInterim := os.Getenv(interimInterval)
	defaultDuration, err := time.ParseDuration(defaultInterim)
	if err != nil {
		return Timer{}, err
	}
	return Timer{
		sessionDuration:        duration,
		defaultInterimInterval: defaultDuration,
	}, nil
}

// IntialiseTimer starts a timer for provided duration using custom input values
func IntialiseTimer(duration, interim string) (Timer, error) {
	sDuration, err := time.ParseDuration(duration)
	if err != nil {
		return Timer{}, err
	}
	interimDuration, err := time.ParseDuration(interim)
	if err != nil {
		return Timer{}, err
	}
	return Timer{
		sessionDuration:        sDuration,
		defaultInterimInterval: interimDuration,
	}, nil
}
func (t Timer) validateDuration() bool {
	return t.sessionDuration > t.defaultInterimInterval
}


// CountInterimTimers provides total number of interim counters required for the session based on defaultInterimInterval
func CountInterimTimers(t Timer) int {
	if t.validateDuration() {
	timerCount := t.sessionDuration.Seconds() / t.defaultInterimInterval.Seconds()
	return int(timerCount)
	}
	return 0
}

// InitiateInterimTimer starts a ticker
func InitiateInterimTimer(t Timer, status chan bool) {
	time.Sleep(t.defaultInterimInterval)
	status <- true
}

// StartInterimTimer starts an interim Timer.
func StartInterimTimer(w io.Writer, t Timer, status chan bool) {
	ticker := time.NewTicker(defaultTickerInterval)
	defer ticker.Stop()
	go InitiateInterimTimer(t, status)
	for {
		select {
		case tc := <-ticker.C:
			fmt.Fprintf(w, "time Now: %v\n", tc)
		case <-status:
			fmt.Fprintf(w, "completed\n")
			return
		}
	}
}
