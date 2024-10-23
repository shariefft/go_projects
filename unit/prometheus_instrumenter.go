// metrics/prometheus_instrumenter.go
package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

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

	// Counter that tracks total number of requests
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
