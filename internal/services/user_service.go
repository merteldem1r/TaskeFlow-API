package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/merteldem1r/TaskeFlow-API/internal/models"
	"github.com/merteldem1r/TaskeFlow-API/internal/repositories"
	"github.com/merteldem1r/TaskeFlow-API/internal/utils"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, email, password, role string) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:        uuid.New().String(),
		Email:     email,
		Password:  hashedPassword,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.Repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
