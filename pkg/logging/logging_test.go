package logging

import (
	"testing"
)

func Test_new_logger(t *testing.T) {
	t.Run("initialize new logger", func(t *testing.T) {
		got := New()
		if got == nil {
			t.Fatal("expected to have valid zap logger")
		}

	})
}
