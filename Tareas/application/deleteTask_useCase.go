package application

import (
	"context"

	"github.com/tu-usuario/taskflow/Tareas/domain/repository"
)

type DeleteTaskUseCase struct {
	taskRepo repository.TaskRepository
}

func NewDeleteTaskUseCase(repo repository.TaskRepository) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{taskRepo: repo}
}

func (u *DeleteTaskUseCase) Execute(ctx context.Context, id, userID int) error {
	_, err := u.taskRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return err
	}

	return u.taskRepo.Delete(ctx, id, userID)
}
