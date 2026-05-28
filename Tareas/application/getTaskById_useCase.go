package application

import (
	"context"

	"github.com/tu-usuario/taskflow/Tareas/domain/entities"
	"github.com/tu-usuario/taskflow/Tareas/domain/repository"
)

type GetTaskByIdUseCase struct {
	taskRepo repository.TaskRepository
}

func NewGetTaskByIdUseCase(repo repository.TaskRepository) *GetTaskByIdUseCase {
	return &GetTaskByIdUseCase{taskRepo: repo}
}

func (u *GetTaskByIdUseCase) Execute(ctx context.Context, id, userID int) (*entities.Task, error) {
	return u.taskRepo.FindByIDAndUserID(ctx, id, userID)
}
