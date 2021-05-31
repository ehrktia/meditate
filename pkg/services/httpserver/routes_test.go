package httpserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_add_new_route(t *testing.T) {
	r := &routeList{routeList: []*routes{}}
	t.Run("should create a new route list", func(t *testing.T) {
		assert.NotEmpty(t, r)
	})

}
