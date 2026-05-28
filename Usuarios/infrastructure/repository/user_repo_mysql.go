package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tu-usuario/taskflow/Usuarios/domain/entities"
	"github.com/tu-usuario/taskflow/Usuarios/domain/repository"
	"github.com/tu-usuario/taskflow/core"
)

type userRepoMySQL struct {
	db *sql.DB
}

func NewUserRepositoryMySQL(db *sql.DB) repository.UserRepository {
	return &userRepoMySQL{
		db: db,
	}
}

func (r *userRepoMySQL) Save(ctx context.Context, user *entities.User) error {
	query := `INSERT INTO users (name, email, password_hash) VALUES (?, ?, ?)`
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

func (r *userRepoMySQL) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := `SELECT id, name, email, password_hash, created_at, updated_at FROM users WHERE email = ? LIMIT 1`
	row := r.db.QueryRowContext(ctx, query, email)

	var user entities.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}
