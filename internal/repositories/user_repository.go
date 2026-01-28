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

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	row := r.DB.QueryRow(ctx, "SELECT id, email, password, role, created_at, updated_at FROM users WHERE email = ?", email)
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}
