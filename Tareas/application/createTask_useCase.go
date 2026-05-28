package application

import (
	"context"
	"time"

	"github.com/tu-usuario/taskflow/Tareas/domain/entities"
	"github.com/tu-usuario/taskflow/Tareas/domain/repository"
)

type CreateTaskUseCase struct {
	taskRepo repository.TaskRepository
}

func NewCreateTaskUseCase(repo repository.TaskRepository) *CreateTaskUseCase {
	return &CreateTaskUseCase{taskRepo: repo}
}

func (u *CreateTaskUseCase) Execute(ctx context.Context, userID int, projectID *int, title, description string, dueDate *time.Time) (*entities.Task, error) {
	task := &entities.Task{
		UserID:      userID,
		ProjectID:   projectID,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Completed:   false,
	}

	if err := u.taskRepo.Save(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}
