package application

import (
	"context"
	"time"

	"github.com/tu-usuario/taskflow/Tareas/domain/entities"
	"github.com/tu-usuario/taskflow/Tareas/domain/repository"
)

type UpdateTaskUseCase struct {
	taskRepo repository.TaskRepository
}

func NewUpdateTaskUseCase(repo repository.TaskRepository) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{taskRepo: repo}
}

func (u *UpdateTaskUseCase) Execute(ctx context.Context, id, userID int, projectID *int, title, description string, dueDate *time.Time, completed bool) (*entities.Task, error) {
	task, err := u.taskRepo.FindByIDAndUserID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	task.ProjectID = projectID
	task.Title = title
	task.Description = description
	task.DueDate = dueDate
	task.Completed = completed

	if err := u.taskRepo.Update(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}
