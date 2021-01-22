package httpserver

import (
	"context"
	"os"
	"testing"
	"time"
)
func Test_create_new_server(t *testing.T) {
    customPort:="0.0.0.0:8080"
    t.Run("should create new server", func(t *testing.T) {
        s:=NewHTTPServer()
        if s==nil {
            t.Error("can not create new instance of http server")
        }
    })
    t.Run("should have a port for server", func(t *testing.T) {
        defaultServer:=NewHTTPServer()
        if defaultServer.Addr=="" {
            t.Error("should have valid port assigned")
        }
    })
    t.Run("should be able to assign custom port", func(t *testing.T) {
        if err:=os.Setenv(httpPort, customPort);err!=nil {
            t.Error(err)
        }
        defer func(){
            if err:=os.Unsetenv(httpPort);err!=nil {
                t.Error(err)
            }
        }()
        s:=NewHTTPServer()
        if s.Addr!=customPort {
            t.Errorf("expected port: %v got: %v",customPort,s.Addr)
        }
    })
    t.Run("should have a valid mux for routing", func(t *testing.T) {
        mux:=NewHTTPServer()
        if mux.Handler==nil {
            t.Errorf("no mux router present for routing")

        }
    })
    t.Run("should be able to run ", func(t *testing.T) {
        ctx,cancel:=context.WithCancel(context.Background())
        srv:=NewHTTPServer()
        go func() {
        if err:=Run(ctx,srv);err!=nil {
            t.Error(err)
        }
        }()
        time.Sleep(2*time.Second)
        cancel()
    })
}
