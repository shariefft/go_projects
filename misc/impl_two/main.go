// main.go
package main

import (
	"net/http"
	"s3proxy/metrics"
	"s3proxy/proxy"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Initialize the Prometheus instrumenter
	instrumenter := metrics.NewPrometheusInstrumenter()

	// Set up the HTTP handlers
	http.HandleFunc("/s3/get", func(w http.ResponseWriter, r *http.Request) {
		proxy.HandleGetRequest(w, r, instrumenter) // Pass instrumenter to the handler
	})

	// Expose the /metrics endpoint for Prometheus
	http.Handle("/metrics", promhttp.Handler())

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
