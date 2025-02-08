package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/lai0xn/orka/internal/api"
	"github.com/lai0xn/orka/internal/api/handler"
	"github.com/lai0xn/orka/internal/api/middleware"
	"github.com/lai0xn/orka/internal/app/keys"
	"github.com/lai0xn/orka/internal/repo"
	"gorm.io/gorm"
)

func Route(r *chi.Mux, apiHandler *handler.APIHandler, db *gorm.DB) {
	r.Use(middleware.RateLimiter(5, 10))

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.RequireAPIKey(*keys.NewApiKeyService(repo.NewKeyRepository(db))))
		r.Post("/schedule", api.MakeHandler(apiHandler.ScheduleTask))
		r.Post("/prod/plan", api.MakeHandler(apiHandler.ProductionPlan))
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", api.MakeHandler(apiHandler.Login))
		r.Post("/signup", api.MakeHandler(apiHandler.Register))
	})

	r.Route("/apikeys", func(r chi.Router) {
		r.Post("/", api.MakeHandler(apiHandler.GenerateApiKey))
		r.Post("/", api.MakeHandler(apiHandler.ValidateApiKey))
		r.Get("/", api.MakeHandler(apiHandler.GetUserApiKeys))
		r.Delete("/{id}", api.MakeHandler(apiHandler.RevokeApiKey))
	})
}
