package server

import (
	"fmt"
	"log"
	"net/http"
	"squawk/config"
	"squawk/server/essentialhandlers"
)

type Server struct {
	server *http.Server
	Mux    *http.ServeMux
}

func CreateNewServer(serverConfig *config.ServerConfig) *Server {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", serverConfig.Port),
		Handler: mux,
	}
	return &Server{
		Mux:    mux,
		server: server,
	}
}

func (s *Server) setEssentialRoutes() {
	log.Println("Setting server essential API routes")
	s.Mux.HandleFunc("GET /ping", essentialhandlers.Ping)
}

func (s *Server) Start() {
	log.Printf("chat server started and listening on: %s", s.server.Addr)

	s.setEssentialRoutes()

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
