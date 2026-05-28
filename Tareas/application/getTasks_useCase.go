package application

import (
	"context"

	"github.com/tu-usuario/taskflow/Tareas/domain/entities"
	"github.com/tu-usuario/taskflow/Tareas/domain/repository"
)

type GetTasksUseCase struct {
	taskRepo repository.TaskRepository
}

func NewGetTasksUseCase(repo repository.TaskRepository) *GetTasksUseCase {
	return &GetTasksUseCase{taskRepo: repo}
}

func (u *GetTasksUseCase) Execute(ctx context.Context, userID int, filters repository.TaskFilters) ([]*entities.Task, error) {
	return u.taskRepo.FindAllByUserID(ctx, userID, filters)
}
