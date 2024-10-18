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

	// Initialize the S3 proxy configuration (inject the instrumenter)
	proxyConfig := &proxy.ProxyConfig{
		S3Client:     initializeS3Client(), // Assume this is your S3 client initialization function
		Instrumenter: instrumenter,
	}

	// Set up the HTTP handlers
	http.HandleFunc("/s3/get", proxyConfig.HandleGetRequest)

	// Expose the /metrics endpoint for Prometheus
	http.Handle("/metrics", promhttp.Handler())

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}

func initializeS3Client() *S3Client {
	// Your logic for initializing and returning the S3 client
	return &S3Client{}
}
