package application

import (
	"context"
	"github.com/tu-usuario/taskflow/Proyectos/domain/entities"
	"github.com/tu-usuario/taskflow/Proyectos/domain/repository"
)

type CreateProjectUseCase struct {
	projectRepo repository.ProjectRepository
}

func NewCreateProjectUseCase(repo repository.ProjectRepository) *CreateProjectUseCase {
	return &CreateProjectUseCase{projectRepo: repo}
}

func (u *CreateProjectUseCase) Execute(ctx context.Context, userID int, name, description string) (*entities.Project, error) {
	project := &entities.Project{
		UserID:      userID,
		Name:        name,
		Description: description,
	}

	if err := u.projectRepo.Save(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}
