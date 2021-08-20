package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/web-alytics/meditate/pkg/model"
)

func homeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
func loginHandler() gin.HandlerFunc {
	userToken := new(model.User)
	return func(c *gin.Context) {
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

func (s *HTTPServer) register(routes *routeList) error {
	for _, route := range routes.routeList {
		switch route.method {
		case http.MethodGet:
			s.Engine.GET(route.path, route.handler)
		case http.MethodPost:
			s.Engine.POST(route.path, route.handler)
		default:
			return fmt.Errorf("not a valid method")
		}
	}
	return nil
}
