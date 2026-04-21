package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	blobpkg "go-application/internal/blob"
	dbpkg "go-application/internal/db"
)

type Server interface {
	Start() error
}

type MuxServer struct {
	router   *mux.Router
	Client   dbpkg.Client        // ✅ FIX
	blob     *blobpkg.BlobClient // ✅ FIX
	producer *KafkaProducer
}

func NewServer(dbClient dbpkg.Client, blobClient *blobpkg.BlobClient, producer *KafkaProducer) Server {
	s := &MuxServer{
		router:   mux.NewRouter(),
		Client:   dbClient,
		blob:     blobClient,
		producer: producer,
	}

	s.routes()

	s.router.Use(corsMiddleware)
	s.router.Use(observabilityMiddleware)

	return s
}

func (s *MuxServer) Start() error {
	slog.Info("server starting", "port", 8080)
	return http.ListenAndServe(":8080", s.router)
}

func observabilityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		r = r.WithContext(ctx)

		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start).Seconds()

		slog.Info("http request",
			"request_id", requestID,
			"method", r.Method,
			"path", r.URL.Path,
			"duration", duration,
			"user_agent", r.UserAgent(),
		)
	})
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
