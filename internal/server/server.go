package server

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/lai0xn/orka/internal/api/rpc"
	"github.com/lai0xn/orka/pkg/pb"
)

type server struct {
	port      string
	twirpPort string
	router    chi.Router
}

type Config struct {
	PORT      string
	TwirpPORT string
	Router    chi.Router
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
	log.Info("Starting Twirp server")

	tw := pb.NewSchedulerServer(&rpc.Scheduler{})
	mux := http.NewServeMux()
	mux.Handle(tw.PathPrefix(), tw)

	go func() {
		if err := http.ListenAndServe(s.twirpPort, mux); err != nil {
			log.Fatal("Couldn't start twirp server", err)
		}
	}()

	if err := http.ListenAndServe(s.port, s.router); err != nil {
		log.Fatal("Server failed to start", err)
	}
}
