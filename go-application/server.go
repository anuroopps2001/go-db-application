package main

import (
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"

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
	server.gorilla.Use(metricsMiddleware)
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

func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w, status: 200}

		start := time.Now()

		next.ServeHTTP(rec, r)

		duration := time.Since(start).Seconds()

		httpRequestsTotal.WithLabelValues(
			r.Method,
			r.URL.Path,
			strconv.Itoa(rec.status),
		).Inc()

		httpRequestDuration.WithLabelValues(
			r.Method,
			r.URL.Path,
		).Observe(duration)
	})
}

func (s *MuxServer) Start() error {
	slog.Info("Serving at port 8080")
	log.Fatal(http.ListenAndServe(":8080", s.gorilla))
	return nil
}
