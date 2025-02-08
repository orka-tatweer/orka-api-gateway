package handler

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lai0xn/orka/internal/api"
	"github.com/lai0xn/orka/internal/app/keys"
)

func (h *APIHandler) GenerateApiKey(w http.ResponseWriter, r *http.Request) error {
	var payload keys.GenerateApiKeyDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	apiKey, err := h.keyService.GenerateApiKey(payload.UserID, payload.Key)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "failed to create API key",
		})
	}

	api.WriteJSON(w, http.StatusCreated, api.Map{
		"api_key": apiKey.Key,
	})
	return nil
}

func (h *APIHandler) ValidateApiKey(w http.ResponseWriter, r *http.Request) error {
	key := r.URL.Query().Get("key")
	apiKey, err := h.keyService.ValidateApiKey(key)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "API key not found",
		})
	}

	api.WriteJSON(w, http.StatusOK, api.Map{
		"api_key": apiKey.Key,
		"user_id": apiKey.UserID,
	})
	return nil
}

func (h *APIHandler) GetUserApiKeys(w http.ResponseWriter, r *http.Request) error {
	user, ok := r.Context().Value("user").(jwt.MapClaims)
	if !ok {
		return api.InvalidJSON(api.Map{
			"error": "invalid user ID",
		})
	}
	userID, ok := user["Id"].(uint)
	apiKeys, err := h.keyService.GetUserApiKeys(userID)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "failed to retrieve API keys",
		})
	}

	api.WriteJSON(w, http.StatusOK, api.Map{
		"api_keys": apiKeys,
	})
	return nil
}

func (h *APIHandler) RevokeApiKey(w http.ResponseWriter, r *http.Request) error {
	key := r.URL.Query().Get("key")
	if err := h.keyService.RevokeApiKey(key); err != nil {
		return api.InvalidJSON(api.Map{
			"error": "failed to revoke API key",
		})
	}

	api.WriteJSON(w, http.StatusOK, api.Map{
		"message": "API key revoked successfully",
	})
	return nil
}
