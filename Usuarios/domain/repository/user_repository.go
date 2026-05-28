package repository

import (
	"context"
	"github.com/tu-usuario/taskflow/Usuarios/domain/entities"
)

type UserRepository interface {
	Save(ctx context.Context, user *entities.User) error
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
}
