package application

import (
	"context"

	"github.com/tu-usuario/taskflow/Usuarios/domain/entities"
	"github.com/tu-usuario/taskflow/Usuarios/domain/repository"
	"github.com/tu-usuario/taskflow/Usuarios/domain/services"
	"github.com/tu-usuario/taskflow/core"
)

type RegisterUserUseCase struct {
	userRepo        repository.UserRepository
	tokenManager    services.TokenManager
	passwordService services.PasswordService
}

func NewRegisterUserUseCase(
	userRepo repository.UserRepository,
	tokenManager services.TokenManager,
	passwordService services.PasswordService,
) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		userRepo:        userRepo,
		tokenManager:    tokenManager,
		passwordService: passwordService,
	}
}

func (u *RegisterUserUseCase) Execute(ctx context.Context, name, email, password string) (*entities.User, string, error) {
	existingUser, _ := u.userRepo.FindByEmail(ctx, email)
	if existingUser != nil {
		return nil, "", core.ErrUserAlreadyExists
	}

	hash, err := u.passwordService.HashPassword(password)
	if err != nil {
		return nil, "", core.ErrInternalServer
	}

	user := &entities.User{
		Name:         name,
		Email:        email,
		PasswordHash: hash,
	}

	if err := u.userRepo.Save(ctx, user); err != nil {
		return nil, "", err
	}

	token, err := u.tokenManager.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, "", core.ErrInternalServer
	}

	return user, token, nil
}
