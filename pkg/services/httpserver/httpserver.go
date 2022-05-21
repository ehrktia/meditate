package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/web-alytics/meditate/pkg/logging"
)

const (
	httpPort    = "PORT"
	defaultPort = "8083"
)

//go:generate mockgen -package=mocks -destination=mocks/${GOFILE} github.com/opentracing/opentracing-go Tracer
//HTTPServer holds required dependencies for http server.
type HTTPServer struct {
	Engine *gin.Engine
	Logger logging.Logger
	Server *http.Server
	Tracer opentracing.Tracer
}

// NewHTTPServer creates new instance of http server.
func NewHTTPServer(log logging.Logger,
	engine *gin.Engine,
	tracer opentracing.Tracer) (*HTTPServer, error) {
	var port string
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	engine.Use(cors.New(config))

	if port = os.Getenv(httpPort); port == "" {
		port = defaultPort
	}
	h := &HTTPServer{
		Engine: engine,
		Logger: log,
		Tracer: tracer,
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

// Run starts http server.
func (s *HTTPServer) Run(ctx context.Context) error {
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

// RegisterRoutes registers handlers and roues.
func (s *HTTPServer) RegisterRoutes() error {
	rList := &routeList{routeList: []*routes{}}
	rList.addRoutes()
	return s.register(rList)
}
