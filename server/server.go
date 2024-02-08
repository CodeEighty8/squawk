package server

import (
	"fmt"
	"log"
	"net/http"
	"squawk/config"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	server *http.Server
	Mux    *chi.Mux
}

func CreateNewServer(serverConfig *config.ServerConfig) *Server {
	mux := chi.NewRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", serverConfig.Port),
		Handler: mux,
	}
	return &Server{
		Mux:    mux,
		server: server,
	}
}

func (s *Server) Start() {
	log.Printf("chat server started and listening on: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
