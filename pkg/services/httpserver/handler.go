package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meditate/pkg/model"
)

func homeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func logoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "logout successful",
		})
	}
}

func loginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken := new(model.User)
		if err := c.ShouldBind(userToken); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "not ok",
				"token":  userToken.IDToken,
				"error":  err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"token": userToken.IDToken,
		})
	}
}

func (h *httpServer) register(routes *routeList) error {
	for _, route := range routes.routeList {
		switch route.method {
		case http.MethodGet:
			h.engine.GET(route.path, route.handler)
		case http.MethodPost:
			h.engine.POST(route.path, route.handler)
		default:
			return fmt.Errorf("not a valid method")
		}
	}
	return nil
}
