package ticker

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	defaultInterimInterval = "INTERIM_INTERVAL"
	defaultTickerInterval  = 1 * time.Second
	defaultSessionDuration = "SESSION_DURATION"
)

type timer struct {
	sessionDuration              string
	sessionDurationFormat        time.Duration
	defaultInterimIntervalFormat time.Duration
}

// IntialiseTimer starts a timer for provided duration
func IntialiseTimer() (timer, error) {
	defaultSession := os.Getenv(defaultSessionDuration)
	duration, err := time.ParseDuration(defaultSession)
	if err != nil {
		return timer{}, err
	}
	defaultInterim := os.Getenv(defaultInterimInterval)
	defaultDuration, err := time.ParseDuration(defaultInterim)
	if err != nil {
		return timer{}, err
	}
	if duration.Seconds() < defaultDuration.Seconds() {
		return timer{}, fmt.Errorf("error session time can not be lower than defaultInterimInterval")
	}
	return timer{
		sessionDuration:              defaultSession,
		sessionDurationFormat:        duration,
		defaultInterimIntervalFormat: defaultDuration,
	}, nil
}

// CountInterimTimers provides total number of interim counters required for the session based on defaultInterimInterval
func CountInterimTimers(t timer) int {
	timerCount := t.sessionDurationFormat.Seconds() / t.defaultInterimIntervalFormat.Seconds()
	return int(timerCount)
}

// InitiateTicker starts a ticker
func InitiateInterimTimer(t timer, status chan bool) {
	time.Sleep(t.defaultInterimIntervalFormat)
	status <- true
}
func StartInterimTimer(w io.Writer, t timer, status chan bool) {
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
