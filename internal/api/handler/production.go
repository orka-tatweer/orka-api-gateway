package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lai0xn/orka/internal/api"
	"github.com/lai0xn/orka/pkg/client"
)

func (h *APIHandler) ProductionPlan(w http.ResponseWriter, r *http.Request) error {
	var payload client.ProductionPlanRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	resp, err := h.prodService.GeneratePlan(payload)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	api.WriteJSON(w, http.StatusCreated, resp)
	return nil
}
