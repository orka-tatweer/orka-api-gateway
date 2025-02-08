package handler

import (
	"github.com/lai0xn/orka/internal/app/auth"
	"github.com/lai0xn/orka/internal/app/scheduler"
)

type APIHandler struct {
	authService auth.AuthService
	taskService scheduler.TaskService
}

func NewAPIHandler(asrv auth.AuthService) *APIHandler {
	return &APIHandler{
		authService: asrv,
	}
}
