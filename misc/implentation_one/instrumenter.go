// metrics/instrumenter.go
package metrics

type Instrumenter interface {
    RecordRequest(statusCode int, method string)
}
