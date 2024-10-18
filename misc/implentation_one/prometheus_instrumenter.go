// metrics/prometheus_instrumenter.go
package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
    "strconv"
)

// PrometheusInstrumenter is the implementation of the Instrumenter interface
type PrometheusInstrumenter struct {
    requestCounter *prometheus.CounterVec
}

// NewPrometheusInstrumenter creates a new instrumenter for custom metrics
func NewPrometheusInstrumenter() *PrometheusInstrumenter {
    requestCounter := prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "s3_requests_total",
            Help: "Total number of S3 requests grouped by status code and method",
        },
        []string{"status_code", "method"},
    )
    
    // Register the counter with Prometheus' default registry
    prometheus.MustRegister(requestCounter)
    
    return &PrometheusInstrumenter{
        requestCounter: requestCounter,
    }
}

// RecordRequest records an S3 request with the given status code and method
func (p *PrometheusInstrumenter) RecordRequest(statusCode int, method string) {
    p.requestCounter.WithLabelValues(strconv.Itoa(statusCode), method).Inc()
}
