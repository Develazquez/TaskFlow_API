package application

import (
	"context"
	"github.com/tu-usuario/taskflow/Proyectos/domain/entities"
	"github.com/tu-usuario/taskflow/Proyectos/domain/repository"
)

type GetProjectByIdUseCase struct {
	projectRepo repository.ProjectRepository
}

func NewGetProjectByIdUseCase(repo repository.ProjectRepository) *GetProjectByIdUseCase {
	return &GetProjectByIdUseCase{projectRepo: repo}
}

func (u *GetProjectByIdUseCase) Execute(ctx context.Context, id, userID int) (*entities.Project, error) {
	return u.projectRepo.FindByIDAndUserID(ctx, id, userID)
}
