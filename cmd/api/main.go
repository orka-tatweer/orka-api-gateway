package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/lai0xn/orka/internal/server"
)

func main() {
	r := chi.NewRouter()
	s := server.NewServer(&server.Config{
		PORT:   ":8080",
		Router: r,
	})

	s.Run()
}
