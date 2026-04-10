package main

import (
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// HTTP Metrics
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

	// DB Latency
	dbQueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_query_duration_seconds",
			Help:    "Database query latency",
			Buckets: []float64{0.01, 0.05, 0.1, 0.2, 0.5, 1, 2},
		},
		[]string{"operation"},
	)

	// 🔥 NEW: Total DB queries
	dbQueriesTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "db_queries_total",
			Help: "Total number of DB queries",
		},
		[]string{"operation"},
	)

	// 🔥 NEW: DB errors
	dbErrorsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "db_errors_total",
			Help: "Total number of DB errors",
		},
		[]string{"operation"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(dbQueryDuration)
	prometheus.MustRegister(dbQueriesTotal)
	prometheus.MustRegister(dbErrorsTotal)
}

func isSystemError(err error) bool {
	msg := err.Error()

	//  business / expected errors → NOT system failures
	if strings.Contains(msg, "duplicate") ||
		strings.Contains(msg, "unique") ||
		strings.Contains(msg, "constraint") {
		return false
	}

	// ✅ everything else → treat as system failure
	return true
}

// Improved DB observer
func observeDB(operation string, fn func() error) error {
	start := time.Now()

	err := fn()

	duration := time.Since(start).Seconds()

	dbQueryDuration.WithLabelValues(operation).Observe(duration)
	dbQueriesTotal.WithLabelValues(operation).Inc()

	// ✅ Only count REAL system failures
	if err != nil && isSystemError(err) {
		dbErrorsTotal.WithLabelValues(operation).Inc()
	}

	return err
}
