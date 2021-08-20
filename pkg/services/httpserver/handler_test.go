package httpserver

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
)

func Test_Home_Handler(t *testing.T) {
	newTestServer := &HTTPServer{
		Engine: gin.Default(),
		Server: &http.Server{},
	}
	assert.Nil(t, newTestServer.RegisterRoutes(), "error registering routes")
	apitest.New("homehandler").
		Debug().
		Handler(newTestServer.Engine).Get("/").
		Expect(t).
		Status(http.StatusOK).
		End()
}
func Test_Login_Handler(t *testing.T) {
	newTestServer := &HTTPServer{
		Engine: gin.Default(),
		Server: &http.Server{},
	}
	assert.Nil(t, newTestServer.RegisterRoutes(), "error registering routes")
	apitest.New("loginhandler").
		Debug().
		Handler(newTestServer.Engine).Post("/login").
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}
