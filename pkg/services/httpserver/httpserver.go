package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/web-alytics/meditate/pkg/logging"
)

const (
	httpPort    = "PORT"
	defaultPort = "8083"
)

type httpServer struct {
	Engine *gin.Engine
	Logger logging.Logger
	Server *http.Server
}

func NewHTTPServer(log logging.Logger,
	engine *gin.Engine) (*httpServer, error) {
	var port string
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Content-Type"}
	engine.Use(cors.New(config))

	if port = os.Getenv(httpPort); port == "" {
		port = defaultPort
	}
	h := &httpServer{
		Engine: engine,
		Logger: log,
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: engine,
		},
	}
	h.Logger.Info("server initalised in address ", port)
	if err := h.RegisterRoutes(); err != nil {
		return nil, err
	}
	h.Logger.Info("routes registration completed")
	return h, nil
}
func (s *httpServer) Run(ctx context.Context) error {
	errCh := make(chan error)
	go func(e chan error) {
		<-ctx.Done()
		if err := s.Server.Shutdown(ctx); err != nil {
			errCh <- err
		}
	}(errCh)
	select {
	case err := <-errCh:
		s.Logger.Errorf("error shutting down server: %v", err)
		return err
	default:
		if err := s.Server.ListenAndServe(); err != nil {
			return err
		}
		return nil
	}
}

func (h *httpServer) RegisterRoutes() error {
	rList := &routeList{routeList: []*routes{}}
	rList.addRoutes()
	return h.register(rList)
}
