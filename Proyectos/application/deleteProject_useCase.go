package application

import (
	"context"
	"github.com/tu-usuario/taskflow/Proyectos/domain/repository"
)

type DeleteProjectUseCase struct {
	projectRepo repository.ProjectRepository
}

func NewDeleteProjectUseCase(repo repository.ProjectRepository) *DeleteProjectUseCase {
	return &DeleteProjectUseCase{projectRepo: repo}
}

func (u *DeleteProjectUseCase) Execute(ctx context.Context, id, userID int) error {
	// Verificar existencia
	_, err := u.projectRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return err
	}

	return u.projectRepo.Delete(ctx, id, userID)
}
