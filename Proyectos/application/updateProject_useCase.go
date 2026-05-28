package application

import (
	"context"
	"github.com/tu-usuario/taskflow/Proyectos/domain/entities"
	"github.com/tu-usuario/taskflow/Proyectos/domain/repository"
)

type UpdateProjectUseCase struct {
	projectRepo repository.ProjectRepository
}

func NewUpdateProjectUseCase(repo repository.ProjectRepository) *UpdateProjectUseCase {
	return &UpdateProjectUseCase{projectRepo: repo}
}

func (u *UpdateProjectUseCase) Execute(ctx context.Context, id, userID int, name, description string) (*entities.Project, error) {
	project, err := u.projectRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	project.Name = name
	project.Description = description

	if err := u.projectRepo.Update(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}
