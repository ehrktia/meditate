package ticker

import "time"
type timer struct {
    tickerDuration time.Duration
    startTime time.Time
    ticker *time.Ticker
    intervalDuration time.Duration
    intervalTicker *time.Ticker
}
// IntialiseTimer starts a timer for provided duration
func IntialiseTimer(d ,intervalTime time.Duration) timer {
    return timer{startTime:time.Now(),
    ticker: time.NewTicker(d),
    intervalDuration: intervalTime,
    intervalTicker: time.NewTicker(intervalTime),
    tickerDuration: d,
}
}
// GetTicker gives the ticker for set duration
func (t *timer) GetTicker() *time.Ticker {
    return t.ticker
}
// GetIntervalTimer provides the ticker to go off in set interval
func (t *timer) GetIntervalTimer() *time.Ticker {
    return t.intervalTicker
}
// TODO:  <23-01-21, karthick> recursive call to interval timer based
// on check to the total duration and interval


