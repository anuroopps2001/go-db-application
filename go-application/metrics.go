package main

import "github.com/prometheus/client_golang/prometheus"

var (
	// Counter: total HTTP requests
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total", // FIXED name
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	// Histogram: request latency
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency",
			Buckets: []float64{0.01, 0.05, 0.1, 0.2, 0.5, 1, 2}, // FIXED buckets
		},
		[]string{"method", "path"}, // FIXED labels
	)
)

var dbQueryDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "db_query_duration_seconds",
		Help:    "Database query latency",
		Buckets: []float64{0.01, 0.05, 0.1, 0.2, 0.5, 1, 2},
	},
	[]string{"operation"},
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(dbQueryDuration)
}
