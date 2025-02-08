package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lai0xn/orka/internal/api"
	"github.com/lai0xn/orka/internal/app/scheduler"
	"github.com/lai0xn/orka/pkg/client"
)

func (h *APIHandler) ScheduleTask(w http.ResponseWriter, r *http.Request) error {
	var payload client.LogisticsRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	resp, err := h.taskService.ScheduleTask(payload)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	api.WriteJSON(w, http.StatusCreated, resp)
	return nil
}

func (h *APIHandler) CreateTask(w http.ResponseWriter, r *http.Request) error {
	var payload scheduler.TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	if err := h.taskService.CreateTask(payload); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	api.WriteJSON(w, http.StatusCreated, api.Map{
		"message": "task created",
	})
	return nil
}

func (h *APIHandler) DeleteTask(w http.ResponseWriter, r *http.Request) error {
	taskIDStr := r.URL.Query().Get("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		return api.InvalidJSON(api.Map{
			"error": "invalid task ID",
		})
	}

	if err := h.taskService.RemoveTask(uint(taskID)); err != nil {
		return api.InvalidJSON(api.Map{
			"error": err.Error(),
		})
	}

	api.WriteJSON(w, http.StatusOK, api.Map{
		"message": "task deleted",
	})
	return nil
}
