package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/merteldem1r/TaskeFlow-API/internal/models"
	"github.com/merteldem1r/TaskeFlow-API/internal/services"
	"github.com/merteldem1r/TaskeFlow-API/internal/utils"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.APIResponse{
			Status: "error",
			Error:  "Invalid request payload",
		})
		return
	}

	user, err := h.Service.Register(r.Context(), input.Email, input.Password, string(models.RoleUser))

	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.APIResponse{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	utils.JSON(w, http.StatusCreated, utils.APIResponse{
		Status: "success",
		Data:   user,
	})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.APIResponse{
			Status: "error",
			Error:  "Invalid request payload",
		})
		return
	}

	user, err := h.Service.Authenticate(r.Context(), input.Email, input.Password)
	if err != nil {
		utils.JSON(w, http.StatusUnauthorized, utils.APIResponse{
			Status: "error",
			Error:  "Invalid email or password",
		})
		return
	}

	token, err := utils.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, utils.APIResponse{
			Status: "error",
			Error:  "Failed to generate token",
		})
		return
	}

	utils.JSON(w, http.StatusOK, utils.APIResponse{
		Status: "success",
		Data: map[string]interface{}{
			"token": token,
			"user":  user,
		},
	})
}
