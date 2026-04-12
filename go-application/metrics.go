package main

import (
	"context"
	"strings"
	"time"

	"log/slog"

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

	dbQueriesTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "db_queries_total",
			Help: "Total number of DB queries",
		},
		[]string{"operation"},
	)

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

	if strings.Contains(msg, "duplicate") ||
		strings.Contains(msg, "unique") ||
		strings.Contains(msg, "constraint") ||
		strings.Contains(msg, "record not found") {
		return false
	}

	return true
}

// 🔥 UPDATED: context-aware logging + metrics
func observeDBWithContext(ctx context.Context, operation string, fn func() error) error {
	start := time.Now()

	err := fn()

	duration := time.Since(start).Seconds()

	dbQueryDuration.WithLabelValues(operation).Observe(duration)
	dbQueriesTotal.WithLabelValues(operation).Inc()

	reqID := getRequestID(ctx)

	if err != nil {
		if isSystemError(err) {
			dbErrorsTotal.WithLabelValues(operation).Inc()
		}

		slog.Error("db operation failed",
			"request_id", reqID,
			"operation", operation,
			"error", err.Error(),
			"duration_seconds", duration,
		)
	} else {
		slog.Info("db operation success",
			"request_id", reqID,
			"operation", operation,
			"duration_seconds", duration,
		)
	}

	return err
}
