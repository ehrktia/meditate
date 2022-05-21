package ticker

import (
	"os"
	"strings"
	"testing"
	"time"

	"go.uber.org/zap/buffer"
)

func TestCreateticker(t *testing.T) {
	interimDuration := "30s"
	sDuration := "1m"
	t.Run("create default timer from env vars", func(t *testing.T) {
		if err := os.Setenv(interimInterval, interimDuration); err != nil {
			t.Error(err)
		}
		if err := os.Setenv(sessionDuration, sDuration); err != nil {
			t.Error(err)
		}
		_, err := DefaultTimerFromEnv()
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("create custom timer from input val", func(t *testing.T) {
	_, err := IntialiseTimer(sDuration, interimDuration)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Cleanup(func() {
		if err := os.Unsetenv(sessionDuration); err != nil {
			t.Error(err)
		}
		if err := os.Unsetenv(interimInterval); err != nil {
			t.Error(err)
		}
	})
}

func Test_count_timers(t *testing.T) {
	tests := []struct {
		name    string
		timer   Timer
		want    int
		wantErr bool
		err     error
	}{
		{
			name: "valid counter",
			timer: Timer{
				sessionDuration:        time.Duration(60),
				defaultInterimInterval: time.Duration(10),
			},
			want: int(time.Duration(60).Seconds() / time.Duration(10).Seconds()),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CountInterimTimers(test.timer)
			if got != test.want {
				t.Logf("[%v]-got,[%v]-want", got, test.want)
			}

		})
	}
}
func TestInterimTicker(t *testing.T) {
	timer, err := IntialiseTimer("10s", "5s")
	if err != nil {
		t.Fatal(err)
	}
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
}
