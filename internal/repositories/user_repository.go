package repositories

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/merteldem1r/TaskeFlow-API/internal/models"
)

type UserRepository struct {
	DB driver.Conn
}

func NewUserRepository(db driver.Conn) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, u *models.User) error {
	return r.DB.Exec(
		ctx,
		`INSERT INTO users (id, email, password, role, created_at, updated_at)
         VALUES (?, ?, ?, ?, ?, ?)`,
		u.ID, u.Email, u.Password, u.Role, u.CreatedAt, u.UpdatedAt,
	)
}
