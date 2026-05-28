package repository

import (
	"context"
	"github.com/tu-usuario/taskflow/Tareas/domain/entities"
)

type TaskFilters struct {
	ProjectID *int
	Completed *bool
}

type TaskRepository interface {
	Save(ctx context.Context, task *entities.Task) error
	FindAllByUserID(ctx context.Context, userID int, filters TaskFilters) ([]*entities.Task, error)
	FindByIDAndUserID(ctx context.Context, id, userID int) (*entities.Task, error)
	Update(ctx context.Context, task *entities.Task) error
	Delete(ctx context.Context, id, userID int) error
}
