package main

import "github.com/prometheus/client_golang/prometheus"

var (
	// // Counter: how many HTTP requests
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	// // Histogram: how long HTTP requests take
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "HTTP request latency",
		},
		[]string{"path"},
	)

	// Histogram: how long the Database calls take (The "Cross-Cloud" bridge)
	sqlRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "sql_request_duration_seconds",
			Help:    "Duration of SQL queries reaching from AKS to GCP",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1}, // Precise buckets for latency
		},
		[]string{"query_type"}, // e.g., "insert", "select"
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(sqlRequestDuration)
}
