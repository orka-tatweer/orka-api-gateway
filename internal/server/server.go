package server

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/lai0xn/orka/internal/api/handler"
	"github.com/lai0xn/orka/internal/api/router"
	"github.com/lai0xn/orka/internal/api/rpc"
	"github.com/lai0xn/orka/internal/app/auth"
	"github.com/lai0xn/orka/internal/app/keys"
	"github.com/lai0xn/orka/internal/app/production"
	"github.com/lai0xn/orka/internal/app/scheduler"
	"github.com/lai0xn/orka/internal/repo"
	"github.com/lai0xn/orka/pkg/pb"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Server struct {
	PORT      string
	TwirpPORT string
	DB        *gorm.DB
	Router    *chi.Mux
	Cache     *redis.Client
}

type Config struct {
	PORT      string
	TwirpPORT string
	DB        *gorm.DB
	Router    *chi.Mux
	Cache     *redis.Client
}

func NewServer(cfg *Config) *Server {
	return &Server{
		PORT:      cfg.PORT,
		Router:    cfg.Router,
		TwirpPORT: cfg.TwirpPORT,
		DB:        cfg.DB,
		Cache:     cfg.Cache,
	}
}

func (s *Server) Run() {
	log.Info("🚀 Server Starting")
	log.Info("Starting Twirp Server")

	tw := pb.NewSchedulerServer(&rpc.Scheduler{})
	mux := http.NewServeMux()
	mux.Handle(tw.PathPrefix(), tw)

	go func() {
		if err := http.ListenAndServe(s.TwirpPORT, mux); err != nil {
			log.Fatal("Couldn't start twirp Server", err)
		}
	}()

	apiHandler := handler.NewAPIHandler(
		auth.NewAuthService(repo.NewAuthRepo(s.DB)),
		scheduler.NewTaskService(repo.NewTaskRepo(s.DB)),
		production.NewProductionScheduler(),
		keys.NewApiKeyService(repo.NewKeyRepository(s.DB)),
	)
	router.Route(s.Router, apiHandler, s.DB)
	if err := http.ListenAndServe(s.PORT, s.Router); err != nil {
		log.Fatal("Server failed to start", err)
	}
}
