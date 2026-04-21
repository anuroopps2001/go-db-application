package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s *MuxServer) routes() {

	s.router.HandleFunc("/user", s.addUser).Methods("POST")
	s.router.HandleFunc("/users", s.listUsers).Methods("GET")
	s.router.HandleFunc("/user/{id}", s.getUser).Methods("GET")
	s.router.HandleFunc("/user/{id}", s.updateUser).Methods("PUT")
	s.router.HandleFunc("/user/{id}", s.deleteUser).Methods("DELETE")

	s.router.HandleFunc("/upload/{id}", s.uploadProfileImage).Methods("POST")

	s.router.Handle("/metrics", promhttp.Handler())

	s.router.HandleFunc("/healthz", s.health).Methods("GET")
	s.router.HandleFunc("/ready", s.ready).Methods("GET")

	s.router.PathPrefix("/ui/").
		Handler(http.StripPrefix("/ui/", http.FileServer(http.Dir("/url"))))
}
