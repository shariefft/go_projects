package main

import (
	"fmt"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

// Instrumenter interface for recording metrics
type Instrumenter interface {
	RecordRequest(statusCode int, method string)
}

// PrometheusInstrumenter is the implementation of the Instrumenter interface
type PrometheusInstrumenter struct {
	requestCounter       *prometheus.CounterVec
	totalRequestsCounter prometheus.Counter
}

func NewPrometheusInstrumenter() *PrometheusInstrumenter {
	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "s3_requests_total",
			Help: "Total number of S3 requests grouped by status code and method",
		},
		[]string{"status_code", "method"},
	)

	totalRequestsCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_s3_requests",
			Help: "Total number of S3 requests",
		},
	)

	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(totalRequestsCounter)

	return &PrometheusInstrumenter{
		requestCounter:       requestCounter,
		totalRequestsCounter: totalRequestsCounter,
	}
}

func (p *PrometheusInstrumenter) RecordRequest(statusCode int, method string) {
	p.requestCounter.WithLabelValues(strconv.Itoa(statusCode), method).Inc()
	p.totalRequestsCounter.Inc()
}

// NoOpInstrumenter is a no-op implementation of the Instrumenter interface
type NoOpInstrumenter struct{}

func (n *NoOpInstrumenter) RecordRequest(statusCode int, method string) {
	// No operation
}

// ProxyConfig struct with an Instrumenter interface field
type ProxyConfig struct {
	Instrumenter Instrumenter
	// Additional proxy configuration fields can go here
}

// CreateProxyServer creates a proxy server with the given config and instrumenter
func CreateProxyServer(conf ProxyConfig) {
	// Example usage of the instrumenter within the proxy server
	conf.Instrumenter.RecordRequest(200, "GET")
	fmt.Println("Proxy server created with metrics instrumentation")
}

func main() {
	// Example configuration: toggle Prometheus metrics here
	usePrometheus := true

	// Select the appropriate instrumenter based on the config
	var instrumenter Instrumenter
	if usePrometheus {
		instrumenter = NewPrometheusInstrumenter()
	} else {
		instrumenter = &NoOpInstrumenter{}
	}

	// Create a ProxyConfig with the chosen instrumenter
	proxyConfig := ProxyConfig{
		Instrumenter: instrumenter,
	}

	// Pass the ProxyConfig to CreateProxyServer
	CreateProxyServer(proxyConfig)
}
