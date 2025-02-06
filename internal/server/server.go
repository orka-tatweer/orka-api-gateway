package server

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
)

type server struct {
	port   string
	router chi.Router
}

type Config struct {
	PORT   string
	Router chi.Router
}

func NewServer(cfg *Config) *server {
	return &server{
		port:   cfg.PORT,
		router: cfg.Router,
	}
}

func (s *server) Setup() {

}

func (s *server) Run() {
	log.Info("🚀 Server Starting")
	if err := http.ListenAndServe(s.port, s.router); err != nil {
		log.Fatal("Server failed to start")
	}
}
