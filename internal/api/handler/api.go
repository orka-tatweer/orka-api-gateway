package handler

import (
	"github.com/lai0xn/orka/internal/app/auth"
	"github.com/lai0xn/orka/internal/app/keys"
	"github.com/lai0xn/orka/internal/app/production"
	"github.com/lai0xn/orka/internal/app/scheduler"
)

type APIHandler struct {
	authService *auth.AuthService
	taskService *scheduler.TaskService
	prodService *production.ProductionScheduler
	keyService  *keys.ApiKeyService
}

func NewAPIHandler(asrv *auth.AuthService, tsrv *scheduler.TaskService, psrv *production.ProductionScheduler, ksrv *keys.ApiKeyService) *APIHandler {
	return &APIHandler{
		authService: asrv,
		taskService: tsrv,
		prodService: psrv,
		keyService:  ksrv,
	}
}
