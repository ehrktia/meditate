package main

import (
	"context"
	"io"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
	"github.com/web-alytics/meditate/pkg/logging"
	"github.com/web-alytics/meditate/pkg/services/httpserver"
	"go.uber.org/zap"
)

	var svcname = "meditate"
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	prodLogger, err := zap.NewProduction()
	if err != nil {
		cancel()
		panic(err)
	}
	log, err := logging.NewLogger(prodLogger.Sugar())
	if err != nil {
		cancel()
		panic(err)
	}
	tracer, closer, err := initJaegerTracer(svcname)
	defer closer.Close()
	if err != nil {
		cancel()
		log.Errorf("error initializing tracing: %v", err)
	}
	engine := gin.New()
	server, err := httpserver.NewHTTPServer(log,
		engine, tracer)
	if err != nil {
		cancel()
		log.Errorf("error creating new server: %v", err)
	}
	if err := server.RegisterRoutes(); err != nil {
		cancel()
		log.Errorf("error registering server routes: %v", err)
	}
	log.Info("starting app server..")
	if err := server.Run(ctx); err != nil {
		log.Errorf("error starting server", err)
	}
	span := tracer.StartSpan("main")
	defer span.Finish()
	span.SetTag("event", "mainfn")
}

// initJaegerTracer creates and intialise tracer.
func initJaegerTracer(svcname string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: svcname,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	return tracer, closer, err
}
