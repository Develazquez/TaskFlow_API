package repository

import (
	"context"
	"github.com/tu-usuario/taskflow/Proyectos/domain/entities"
)

type ProjectRepository interface {
	Save(ctx context.Context, project *entities.Project) error
	FindAllByUserID(ctx context.Context, userID int) ([]*entities.Project, error)
	FindByIDAndUserID(ctx context.Context, id, userID int) (*entities.Project, error)
	Update(ctx context.Context, project *entities.Project) error
	Delete(ctx context.Context, id, userID int) error
}
