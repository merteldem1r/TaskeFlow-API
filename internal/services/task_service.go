package services

import (
	"context"

	"github.com/merteldem1r/TaskeFlow-API/internal/models"
	"github.com/merteldem1r/TaskeFlow-API/internal/repositories"
)

type TaskService struct {
	Repo *repositories.TaskRepository
}

func NewTaskService(repo *repositories.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]*models.Task, error) {
	return s.Repo.GetAll(ctx)
}

func (s *TaskService) CreateTask(ctx context.Context, t *models.Task) error {
	return s.Repo.Create(ctx, t)
}
