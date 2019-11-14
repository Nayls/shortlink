package main

import (
	"context"
	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/traicing"
	"github.com/batazor/shortlink/pkg/api"
	"github.com/batazor/shortlink/pkg/api/graphql"
	grpcweb "github.com/batazor/shortlink/pkg/api/grpc-web"
	httpchi "github.com/batazor/shortlink/pkg/api/http-chi"
	"github.com/heptiolabs/healthcheck"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"net/http"
)

type Service struct {
	log         logger.Logger
	tracer      opentracing.Tracer
	tracerClose io.Closer
}

func (s *Service) initLogger() {
	var err error
	conf := logger.Configuration{
		Level: logger.INFO_LEVEL,
	}

	if s.log, err = logger.NewLogger(logger.Zap, conf); err != nil {
		panic(err)
	}
}

func (s *Service) initTracer() {
	var err error
	if s.tracer, s.tracerClose, err = traicing.Init(); err != nil {
		s.log.Error(err.Error())
	}
}

// runAPIServer - start HTTP-server
func (s *Service) runAPIServer(ctx context.Context) {
	var API api.API
	serverType := "http-chi"

	switch serverType {
	case "http-chi":
		API = &httpchi.API{}
	case "gRPC-web":
		API = &grpcweb.API{}
	case "graphql":
		API = &graphql.API{}
	default:
		API = &httpchi.API{}
	}

	if err := API.Run(ctx); err != nil {
		s.log.Fatal(err.Error())
	}
}

func (s *Service) initMonitoring() *http.ServeMux {
	// Create a new Prometheus registry
	registry := prometheus.NewRegistry()

	// Create a metrics-exposing Handler for the Prometheus registry
	// The healthcheck related metrics will be prefixed with the provided namespace
	health := healthcheck.NewMetricsHandler(registry, "common")

	// Add a liveness check that always succeeds
	health.AddLivenessCheck("successful-check", func() error {
		return nil
	})

	// Create an "common" listener
	commonMux := http.NewServeMux()

	// Expose prometheus metrics on /metrics
	commonMux.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	// Expose a liveness check on /live
	commonMux.HandleFunc("/live", health.LiveEndpoint)

	// Expose a readiness check on /ready
	commonMux.HandleFunc("/ready", health.ReadyEndpoint)

	return commonMux
}

// Start - run this a service
func (s *Service) Start() {
	// Create a new context
	ctx := context.Background()

	// Logger
	s.initLogger()
	ctx = logger.WithLogger(ctx, s.log) // Add logger to context

	// Add Tracer
	s.initTracer()
	ctx = traicing.WithTraicer(ctx, s.tracer) // Add tracer to context

	// Monitoring endpoints
	monitoringServer := s.initMonitoring()
	go http.ListenAndServe("0.0.0.0:9090", monitoringServer) // nolint errcheck

	// Run API server
	s.runAPIServer(ctx)
}

// Stop - stop this a service
func (s *Service) Stop() {
	if err := s.tracerClose.Close(); err != nil {
		s.log.Error(err.Error())
	}

	// flushes buffer, if any
	s.log.Close()
}
