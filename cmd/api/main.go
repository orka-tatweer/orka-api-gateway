package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/lai0xn/orka/internal/infra/db"
	"github.com/lai0xn/orka/internal/infra/redis"
	"github.com/lai0xn/orka/internal/server"
)

func main() {
	r := chi.NewRouter()
	db, err := db.Connect()

	rdb := redis.Connect()

	s := server.NewServer(&server.Config{
		PORT:      ":8080",
		Router:    r,
		TwirpPORT: ":8000",
		DB:        db,
		Cache:     rdb,
	})

	if err != nil {
		panic(err)
	}

	s.Run()
}
