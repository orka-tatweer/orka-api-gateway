package middleware

import (
	"net/http"
	"strings"

	"github.com/lai0xn/orka/internal/api"
	"github.com/lai0xn/orka/internal/app/keys"
)

func RequireAPIKey(keyService keys.ApiKeyService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("Authorization")

			if apiKey == "" || !strings.HasPrefix(apiKey, "Bearer ") {
				api.WriteJSON(w, http.StatusUnauthorized, api.Map{"error": "missing API key"})
				return
			}

			apiKey = strings.TrimPrefix(apiKey, "Bearer ")

			// Validate API key
			key, err := keyService.ValidateApiKey(apiKey)
			if err != nil || key == nil {
				api.WriteJSON(w, http.StatusUnauthorized, api.Map{"error": "invalid API key"})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
