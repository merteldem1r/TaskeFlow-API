package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/merteldem1r/TaskeFlow-API/internal/services"
	"github.com/merteldem1r/TaskeFlow-API/internal/utils"
)

type TaskHandler struct {
	Service *services.TaskService
	// Later: TaskService will be injected here
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Service.GetAllTasks(r.Context())

	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.APIResponse{
			Status: "error",
			Error:  "Failed to fetch tasks",
		})
		return
	}

	utils.JSON(w, http.StatusOK, utils.APIResponse{
		Status: "success",
		Data:   tasks,
		Count:  len(tasks),
	})
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")

	response := map[string]string{
		"message": "Get task by ID",
		"id":      taskID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Create task"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")

	response := map[string]string{
		"message": "Update task",
		"id":      taskID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")

	response := map[string]string{
		"message": "Delete task",
		"id":      taskID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
