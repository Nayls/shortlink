/*
Shortlink application

API-service
*/
package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/batazor/shortlink/internal/config"
	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/error/status"
	"github.com/batazor/shortlink/pkg/api"
)

func init() {
	// Read ENV variables
	if err := config.Init(); err != nil {
		fmt.Println(err.Error())
		os.Exit(status.ERROR_CONFIG)
	}
}

func main() {
	// Create a new context
	ctx := context.Background()

	// Init a new service
	s, cleanup, err := di.InitializeFullService(ctx)
	if err != nil { // TODO: use as helpers
		if r, ok := err.(*net.OpError); ok {
			panic(fmt.Errorf("address %s already in use. Set GRPC_SERVER_PORT enviroment", r.Addr.String()))
		}

		panic(err)
	}

	// Monitoring endpoints
	go http.ListenAndServe("0.0.0.0:9090", s.Monitoring) // nolint errcheck

	var profiling *http.ServeMux = s.PprofEndpoint
	go http.ListenAndServe("0.0.0.0:7071", profiling) // nolint errcheck

	// Run API server
	var API api.Server
	API.RunAPIServer(ctx, s.Log, s.Tracer, s.ServerRPC, s.ClientRPC)

	defer func() {
		if r := recover(); r != nil {
			s.Log.Error(r.(string))
		}
	}()

	// Handle SIGINT and SIGTERM.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Context close
	ctx.Done()

	// Stop the service gracefully.
	// close DB
	if err := s.DB.Store.Close(); err != nil {
		s.Log.Error(err.Error())
	}

	cleanup()
}
