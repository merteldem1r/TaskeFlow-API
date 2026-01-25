package repositories

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/merteldem1r/TaskeFlow-API/internal/models"
)

type TaskRepository struct {
	DB driver.Conn
}

func NewTaskRepository(db driver.Conn) *TaskRepository {
	return &TaskRepository{DB: db}
}

// Get all tasks
func (r *TaskRepository) GetAll(ctx context.Context) ([]*models.Task, error) {
	rows, err := r.DB.Query(ctx, "SELECT id, title, description, status, user_id, created_at, updated_at FROM tasks")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task

	for rows.Next() {
		var t models.Task
		if err := rows.Scan(
			&t.ID, &t.Title, &t.Description, &t.Status, &t.UserID, &t.CreatedAt, &t.UpdatedAt,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, &t)
	}

	if tasks == nil {
		tasks = []*models.Task{}
	}

	return tasks, nil
}

func (r *TaskRepository) Create(ctx context.Context, t *models.Task) error {
	return r.DB.Exec(
		ctx,
		`INSERT INTO tasks (id, title, description, status, user_id, created_at, updated_at)
         VALUES (?, ?, ?, ?, ?, ?, ?)`,
		t.ID, t.Title, t.Description, t.Status, t.UserID, t.CreatedAt, t.UpdatedAt,
	)
}
