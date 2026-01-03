package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s *MuxServer) routes() {
	s.gorilla.HandleFunc("/user", s.addUser).Methods("POST")
	s.gorilla.HandleFunc("/users", s.listUsers).Methods("GET")
	s.gorilla.HandleFunc("/user/{id}", s.updateUser).Methods("PUT")
	s.gorilla.HandleFunc("/user/{id}", s.deleteUser).Methods("DELETE")

	//prometheus endpoint
	s.gorilla.Handle("/metrics", promhttp.Handler())
}
