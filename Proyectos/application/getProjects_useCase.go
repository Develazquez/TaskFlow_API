package application

import (
	"context"
	"github.com/tu-usuario/taskflow/Proyectos/domain/entities"
	"github.com/tu-usuario/taskflow/Proyectos/domain/repository"
)

type GetProjectsUseCase struct {
	projectRepo repository.ProjectRepository
}

func NewGetProjectsUseCase(repo repository.ProjectRepository) *GetProjectsUseCase {
	return &GetProjectsUseCase{projectRepo: repo}
}

func (u *GetProjectsUseCase) Execute(ctx context.Context, userID int) ([]*entities.Project, error) {
	return u.projectRepo.FindAllByUserID(ctx, userID)
}
