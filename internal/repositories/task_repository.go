package repositories

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/merteldem1r/TaskeFlow-API/internal/models"
)

type TaskRepository struct {
	DB     driver.Conn
	DBName string
}

func NewTaskRepository(db driver.Conn, dbName string) *TaskRepository {
	return &TaskRepository{DB: db, DBName: dbName}
}

// Get all tasks
func (r *TaskRepository) GetAll(ctx context.Context) ([]*models.Task, error) {
	query := fmt.Sprintf("SELECT id, title, description, status, user_id, created_at, updated_at FROM %s.tasks", r.DBName)

	rows, err := r.DB.Query(ctx, query)

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
	query := fmt.Sprintf(`INSERT INTO %s.tasks (id, title, description, status, user_id, created_at, updated_at)
		 VALUES (?, ?, ?, ?, ?, ?, ?)`, r.DBName)

	return r.DB.Exec(
		ctx,
		query,
		t.ID, t.Title, t.Description, t.Status, t.UserID, t.CreatedAt, t.UpdatedAt,
	)
}
