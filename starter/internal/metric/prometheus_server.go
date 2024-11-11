package metric

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	prometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusServer struct {
	registry *prometheus.Registry
	host     PrometheusHost
	port     PrometheusPort
	path     PrometheusPath
}

func NewPrometheusServer(registry *prometheus.Registry, host PrometheusHost, port PrometheusPort, path PrometheusPath) (*PrometheusServer, error) {
	// Validate port range
	if port < 0 || port > 65535 {
		return nil, errors.New("invalid port: must be between 0 and 65535")
	}

	// Initialize and return the exporter
	return &PrometheusServer{registry: registry, host: host, port: port, path: path}, nil
}

func (ps *PrometheusServer) Addr() string {
	// Construct the server address
	h := strings.Trim(string(ps.host), "/")
	addr := fmt.Sprintf("%s:%d", h, ps.port)

	return addr
}

func (ps *PrometheusServer) Path() string {
	return "/" + strings.Trim(string(ps.path), "/")
}

// Start begins serving metrics asynchronously and returns a shutdown function.
func (ps *PrometheusServer) Start(ctx context.Context) (func(), error) {
	mux := http.NewServeMux()
	mux.Handle(ps.Path(), ps)
	srv := &http.Server{
		Addr:    ps.Addr(),
		Handler: mux,
	}

	// Start the server in a new goroutine to run asynchronously
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// TODO: handle server startup failure
		}
	}()

	// Define and return shutdown function
	shutdown := func() {
		if err := srv.Shutdown(ctx); err != nil {
			// TODO: handle shutdown error
		}
	}

	return shutdown, nil
}

func (ps *PrometheusServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	promhttp.HandlerFor(ps.registry, promhttp.HandlerOpts{}).ServeHTTP(w, req)
}
