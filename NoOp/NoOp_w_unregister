package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

// Instrumenter interface for recording metrics
type Instrumenter interface {
	RecordRequest(statusCode int, method string)
	Unregister()
}

// PrometheusInstrumenter is the implementation of the Instrumenter interface
type PrometheusInstrumenter struct {
	requestCounter       *prometheus.CounterVec
	totalRequestsCounter prometheus.Counter
}

// NewPrometheusInstrumenter creates a new instrumenter for custom metrics
func NewPrometheusInstrumenter() *PrometheusInstrumenter {
	// Define the counters
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

	// Register the counters
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(totalRequestsCounter)

	return &PrometheusInstrumenter{
		requestCounter:       requestCounter,
		totalRequestsCounter: totalRequestsCounter,
	}
}

// RecordRequest records an S3 request with the given status code and method
func (p *PrometheusInstrumenter) RecordRequest(statusCode int, method string) {
	p.requestCounter.WithLabelValues(strconv.Itoa(statusCode), method).Inc()
	p.totalRequestsCounter.Inc()
}

// Unregister unregisters the metrics, useful for cleanup in tests
func (p *PrometheusInstrumenter) Unregister() {
	prometheus.Unregister(p.requestCounter)
	prometheus.Unregister(p.totalRequestsCounter)
}

// NoOpInstrumenter is a no-op implementation of the Instrumenter interface
type NoOpInstrumenter struct{}

func (n *NoOpInstrumenter) RecordRequest(statusCode int, method string) {
	// No operation
}

func (n *NoOpInstrumenter) Unregister() {
	// No operation for no-op instrumenter
}
