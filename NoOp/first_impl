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

// NewPrometheusInstrumenter creates a new instrumenter for custom metrics
func NewPrometheusInstrumenter() *PrometheusInstrumenter {
	// Counter that tracks requests by status code and method
	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "s3_requests_total",
			Help: "Total number of S3 requests grouped by status code and method",
		},
		[]string{"status_code", "method"},
	)

	// Counter that tracks the total number of requests
	totalRequestsCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_s3_requests",
			Help: "Total number of S3 requests",
		},
	)

	// Register both counters with Prometheus' default registry
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(totalRequestsCounter)

	return &PrometheusInstrumenter{
		requestCounter:       requestCounter,
		totalRequestsCounter: totalRequestsCounter,
	}
}

// RecordRequest records an S3 request with the given status code and method
func (p *PrometheusInstrumenter) RecordRequest(statusCode int, method string) {
	// Increment the request counter with labels
	p.requestCounter.WithLabelValues(strconv.Itoa(statusCode), method).Inc()
	// Increment the total requests counter
	p.totalRequestsCounter.Inc()
}

// NoOpInstrumenter is a no-op implementation of the Instrumenter interface
type NoOpInstrumenter struct{}

func (n *NoOpInstrumenter) RecordRequest(statusCode int, method string) {
	// No operation
}

// processWithMetrics takes an Instrumenter to record metrics if enabled
func processWithMetrics(instr Instrumenter, statusCode int, method string) {
	instr.RecordRequest(statusCode, method)
	fmt.Println("Processing request with metrics instrumentation")
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

	// Pass the instrumenter to the function with example values
	processWithMetrics(instrumenter, 200, "GET")
}
