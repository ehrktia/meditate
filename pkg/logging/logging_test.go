package logging

import (
	"testing"
)
func Test_new_logger(t *testing.T) {
    got,err:=NewLogger()
    if err!=nil {
        t.Error(err)
    }
    if got==nil {
        t.Error("can not create logger")
    }
    t.Run("should be able to use the logger", func(t *testing.T) {
        got.Info("logger initializes")
    })
}
