package main

import (
	"context"
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
	blob     *BlobClient
	producer *KafkaProducer // 👈 NEW
}

func NewServer(db Client, blob *BlobClient, producer *KafkaProducer) Server {
	server := &MuxServer{
		gorilla:  mux.NewRouter(),
		Client:   db,
		blob:     blob,
		producer: producer, // 👈 NEW
	}

	server.routes()

	server.gorilla.Use(corsMiddleware)
	server.gorilla.Use(observabilityMiddleware)

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

func observabilityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 🔥 filter noisy endpoints
		if r.URL.Path == "/healthz" || r.URL.Path == "/ready" || r.URL.Path == "/metrics" {
			next.ServeHTTP(w, r)
			return
		}

		requestID := uuid.New().String()

		// 🔥 propagate request_id via context
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		r = r.WithContext(ctx)

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

		// Logs
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

func getRequestID(ctx context.Context) string {
	if v, ok := ctx.Value(requestIDKey).(string); ok {
		return v
	}
	return ""
}

func (s *MuxServer) Start() error {
	slog.Info("server starting", "port", 8080)
	log.Fatal(http.ListenAndServe(":8080", s.gorilla))
	return nil
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
