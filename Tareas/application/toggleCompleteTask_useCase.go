package application

import (
	"context"

	"github.com/tu-usuario/taskflow/Tareas/domain/entities"
	"github.com/tu-usuario/taskflow/Tareas/domain/repository"
)

type ToggleCompleteTaskUseCase struct {
	taskRepo repository.TaskRepository
}

func NewToggleCompleteTaskUseCase(repo repository.TaskRepository) *ToggleCompleteTaskUseCase {
	return &ToggleCompleteTaskUseCase{taskRepo: repo}
}

func (u *ToggleCompleteTaskUseCase) Execute(ctx context.Context, id, userID int) (*entities.Task, error) {
	task, err := u.taskRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	task.Completed = !task.Completed

	if err := u.taskRepo.Update(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}
