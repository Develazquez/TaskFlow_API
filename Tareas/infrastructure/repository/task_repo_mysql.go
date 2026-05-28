package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tu-usuario/taskflow/Tareas/domain/entities"
	"github.com/tu-usuario/taskflow/Tareas/domain/repository"
	"github.com/tu-usuario/taskflow/core"
)

type taskRepoMySQL struct {
	db *sql.DB
}

func NewTaskRepositoryMySQL(db *sql.DB) repository.TaskRepository {
	return &taskRepoMySQL{
		db: db,
	}
}

func (r *taskRepoMySQL) Save(ctx context.Context, task *entities.Task) error {
	query := `INSERT INTO tasks (user_id, project_id, title, description, due_date, completed) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := r.db.ExecContext(ctx, query, task.UserID, task.ProjectID, task.Title, task.Description, task.DueDate, task.Completed)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	task.ID = int(id)
	return nil
}

func (r *taskRepoMySQL) FindAllByUserID(ctx context.Context, userID int, filters repository.TaskFilters) ([]*entities.Task, error) {
	query := `SELECT id, user_id, project_id, title, description, due_date, completed, created_at, updated_at FROM tasks WHERE user_id = ?`
	args := []interface{}{userID}

	if filters.ProjectID != nil {
		query += ` AND project_id = ?`
		args = append(args, *filters.ProjectID)
	}

	if filters.Completed != nil {
		query += ` AND completed = ?`
		args = append(args, *filters.Completed)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*entities.Task
	for rows.Next() {
		var t entities.Task
		if err := rows.Scan(&t.ID, &t.UserID, &t.ProjectID, &t.Title, &t.Description, &t.DueDate, &t.Completed, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, nil
}

func (r *taskRepoMySQL) FindByIDAndUserID(ctx context.Context, id, userID int) (*entities.Task, error) {
	query := `SELECT id, user_id, project_id, title, description, due_date, completed, created_at, updated_at FROM tasks WHERE id = ? AND user_id = ? LIMIT 1`
	row := r.db.QueryRowContext(ctx, query, id, userID)

	var t entities.Task
	err := row.Scan(&t.ID, &t.UserID, &t.ProjectID, &t.Title, &t.Description, &t.DueDate, &t.Completed, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.ErrTaskNotFound
		}
		return nil, err
	}
	return &t, nil
}

func (r *taskRepoMySQL) Update(ctx context.Context, task *entities.Task) error {
	query := `UPDATE tasks SET project_id = ?, title = ?, description = ?, due_date = ?, completed = ? WHERE id = ? AND user_id = ?`
	_, err := r.db.ExecContext(ctx, query, task.ProjectID, task.Title, task.Description, task.DueDate, task.Completed, task.ID, task.UserID)
	return err
}

func (r *taskRepoMySQL) Delete(ctx context.Context, id, userID int) error {
	query := `DELETE FROM tasks WHERE id = ? AND user_id = ?`
	_, err := r.db.ExecContext(ctx, query, id, userID)
	return err
}
