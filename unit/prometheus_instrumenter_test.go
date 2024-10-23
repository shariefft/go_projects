// metrics/prometheus_instrumenter_test.go
package metrics

import (
    "strconv"
    "testing"
    "github.com/prometheus/client_golang/prometheus"
)

// TestPrometheusInstrumenter_RecordRequestWithoutDTO tests that the request counter increments properly without dto
func TestPrometheusInstrumenter_RecordRequestWithoutDTO(t *testing.T) {
    // Create a new registry to avoid polluting the default one
    reg := prometheus.NewRegistry()

    // Create a new instrumenter
    p := NewPrometheusInstrumenter()

    // Register the counters with the test registry
    reg.MustRegister(p.requestCounter)
    reg.MustRegister(p.totalRequestsCounter)

    // Define test cases
    testCases := []struct {
        statusCode int
        method     string
        expected   float64
    }{
        {statusCode: 200, method: "GET", expected: 1},
        {statusCode: 404, method: "POST", expected: 1},
        {statusCode: 500, method: "PUT", expected: 1},
    }

    var totalExpected float64

    // Execute test cases
    for _, tc := range testCases {
        totalExpected += tc.expected
        t.Run(strconv.Itoa(tc.statusCode)+"_"+tc.method, func(t *testing.T) {
            // Record a request
            p.RecordRequest(tc.statusCode, tc.method)

            // Gather all metrics from the custom registry
            metricFamilies, err := reg.Gather()
            if err != nil {
                t.Fatalf("error gathering metrics: %v", err)
            }

            // Find the metric family for "s3_requests_total"
            var found bool
            for _, mf := range metricFamilies {
                if *mf.Name == "s3_requests_total" {
                    found = true

                    // Iterate through the metrics to find the one with matching labels
                    for _, m := range mf.Metric {
                        labels := m.GetLabel()
                        var statusCodeLabel, methodLabel string
                        for _, label := range labels {
                            if *label.Name == "status_code" {
                                statusCodeLabel = *label.Value
                            }
                            if *label.Name == "method" {
                                methodLabel = *label.Value
                            }
                        }

                        if statusCodeLabel == strconv.Itoa(tc.statusCode) && methodLabel == tc.method {
                            // Check if the counter value matches the expected value
                            if *m.Counter.Value != tc.expected {
                                t.Errorf("expected counter to be %v, got %v", tc.expected, *m.Counter.Value)
                            }
                        }
                    }
                }
            }

            if !found {
                t.Fatalf("metric s3_requests_total not found")
            }
        })
    }

    // Check total requests counter
    metricFamilies, err := reg.Gather()
    if err != nil {
        t.Fatalf("error gathering metrics: %v", err)
    }

    var totalRequestsCounterFound bool
    for _, mf := range metricFamilies {
        if *mf.Name == "total_s3_requests" {
            totalRequestsCounterFound = true
            if *mf.Metric[0].Counter.Value != totalExpected {
                t.Errorf("expected total requests counter to be %v, got %v", totalExpected, *mf.Metric[0].Counter.Value)
            }
        }
    }

    if !totalRequestsCounterFound {
        t.Fatalf("metric total_s3_requests not found")
    }
}
