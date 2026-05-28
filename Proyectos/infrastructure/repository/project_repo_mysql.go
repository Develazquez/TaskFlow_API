package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tu-usuario/taskflow/Proyectos/domain/entities"
	"github.com/tu-usuario/taskflow/Proyectos/domain/repository"
	"github.com/tu-usuario/taskflow/core"
)

type projectRepoMySQL struct {
	db *sql.DB
}

func NewProjectRepositoryMySQL(db *sql.DB) repository.ProjectRepository {
	return &projectRepoMySQL{
		db: db,
	}
}

func (r *projectRepoMySQL) Save(ctx context.Context, project *entities.Project) error {
	query := `INSERT INTO projects (user_id, name, description) VALUES (?, ?, ?)`
	result, err := r.db.ExecContext(ctx, query, project.UserID, project.Name, project.Description)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	project.ID = int(id)
	return nil
}

func (r *projectRepoMySQL) FindAllByUserID(ctx context.Context, userID int) ([]*entities.Project, error) {
	query := `SELECT id, user_id, name, description, created_at, updated_at FROM projects WHERE user_id = ?`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*entities.Project
	for rows.Next() {
		var p entities.Project
		if err := rows.Scan(&p.ID, &p.UserID, &p.Name, &p.Description, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, &p)
	}
	return projects, nil
}

func (r *projectRepoMySQL) FindByIDAndUserID(ctx context.Context, id, userID int) (*entities.Project, error) {
	query := `SELECT id, user_id, name, description, created_at, updated_at FROM projects WHERE id = ? AND user_id = ? LIMIT 1`
	row := r.db.QueryRowContext(ctx, query, id, userID)

	var p entities.Project
	err := row.Scan(&p.ID, &p.UserID, &p.Name, &p.Description, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.ErrProjectNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (r *projectRepoMySQL) Update(ctx context.Context, project *entities.Project) error {
	query := `UPDATE projects SET name = ?, description = ? WHERE id = ? AND user_id = ?`
	_, err := r.db.ExecContext(ctx, query, project.Name, project.Description, project.ID, project.UserID)
	return err
}

func (r *projectRepoMySQL) Delete(ctx context.Context, id, userID int) error {
	query := `DELETE FROM projects WHERE id = ? AND user_id = ?`
	_, err := r.db.ExecContext(ctx, query, id, userID)
	return err
}
