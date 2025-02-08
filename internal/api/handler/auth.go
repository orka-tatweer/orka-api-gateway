package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lai0xn/orka/internal/api"
	"github.com/lai0xn/orka/internal/app/auth"
	"github.com/lai0xn/orka/pkg/utils"
)

func (h *APIHandler) Login(w http.ResponseWriter, r *http.Request) error {
	var payload auth.LoginDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}
	user, err := h.authService.Authenticate(payload)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "wrong credentials",
		})
	}
	token, err := utils.GenerateToken(*user)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "wrong credentials",
		})
	}

	api.WriteJSON(w, http.StatusCreated, api.Map{
		"token": token,
	})
	return nil

}

func (h *APIHandler) Register(w http.ResponseWriter, r *http.Request) error {
	var payload auth.SignupDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}
	if err := h.authService.Signup(payload); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}
	api.WriteJSON(w, http.StatusCreated, api.Map{
		"message": "user created",
	})
	return nil

}
