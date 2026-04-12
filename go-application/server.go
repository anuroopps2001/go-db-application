package main

import (
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server interface {
	Start() error
	routes()
}

type MuxServer struct {
	gorilla *mux.Router
	Client
}

func NewServer(db Client) Server {
	server := &MuxServer{
		mux.NewRouter(),
		db,
	}

	server.routes()
	server.gorilla.Use(observabilityMiddleware) // 🔥 changed here
	return server
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// 🔥 NEW: Combined logging + metrics middleware
func observabilityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestID := uuid.New().String()

		rec := &statusRecorder{ResponseWriter: w, status: 200}

		start := time.Now()

		next.ServeHTTP(rec, r)

		duration := time.Since(start).Seconds()

		// Metrics
		httpRequestsTotal.WithLabelValues(
			r.Method,
			r.URL.Path,
			strconv.Itoa(rec.status),
		).Inc()

		httpRequestDuration.WithLabelValues(
			r.Method,
			r.URL.Path,
		).Observe(duration)

		// 🔥 Structured logging
		slog.Info("http request",
			"request_id", requestID,
			"method", r.Method,
			"path", r.URL.Path,
			"status", rec.status,
			"duration_seconds", duration,
			"user_agent", r.UserAgent(),
			"remote_addr", r.RemoteAddr,
		)
	})
}

func (s *MuxServer) Start() error {
	slog.Info("server starting", "port", 8080)
	log.Fatal(http.ListenAndServe(":8080", s.gorilla))
	return nil
}
