package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	// Later: TaskService will be injected here
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// TODO: Get from the service/database
	response := map[string]string{"message": "Get all tasks"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
