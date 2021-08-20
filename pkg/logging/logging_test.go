package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_new_logger(t *testing.T) {
	prodLogger,err:=zap.NewProduction()
if 	assert.NoError(t,err){
	got,err:=NewLogger(prodLogger.Sugar())
	if assert.NoError(t, err) {
		assert.NotNil(t, got)
		assert.IsType(t, &PkgLogger{}, got)
	}
}

}
