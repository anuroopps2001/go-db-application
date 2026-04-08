package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency",
			Buckets: []float64{0.01, 0.05, 0.1, 0.2, 0.5, 1, 2},
		},
		[]string{"method", "path"},
	)

	dbQueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_query_duration_seconds",
			Help:    "Database query latency",
			Buckets: []float64{0.01, 0.05, 0.1, 0.2, 0.5, 1, 2},
		},
		[]string{"operation"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(dbQueryDuration)
}

// 🔥 Helper function (important)
func observeDB(operation string, fn func() error) error {
	start := time.Now()

	err := fn()

	duration := time.Since(start).Seconds()
	dbQueryDuration.WithLabelValues(operation).Observe(duration)

	return err
}
