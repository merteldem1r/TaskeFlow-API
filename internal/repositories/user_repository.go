package repositories

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/merteldem1r/TaskeFlow-API/internal/models"
)

type UserRepository struct {
	DB     driver.Conn
	DBName string
}

func NewUserRepository(db driver.Conn, dbName string) *UserRepository {
	return &UserRepository{
		DB:     db,
		DBName: dbName,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, u *models.User) error {
	query := fmt.Sprintf(
		"INSERT INTO %s.users (id, email, password, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		r.DBName,
	)
	return r.DB.Exec(ctx, query, u.ID, u.Email, u.Password, u.Role, u.CreatedAt, u.UpdatedAt)
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := fmt.Sprintf(
		"SELECT id, email, password, role, created_at, updated_at FROM %s.users WHERE email = ?",
		r.DBName,
	)
	row := r.DB.QueryRow(ctx, query, email)
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	query := fmt.Sprintf(
		"SELECT id, email, password, role, created_at, updated_at FROM %s.users WHERE id = ?",
		r.DBName,
	)

	row := r.DB.QueryRow(ctx, query, id)
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}
