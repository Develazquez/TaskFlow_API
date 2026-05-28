package application

import (
	"context"

	"github.com/tu-usuario/taskflow/Usuarios/domain/repository"
	"github.com/tu-usuario/taskflow/Usuarios/domain/services"
	"github.com/tu-usuario/taskflow/core"
)

type LoginUserUseCase struct {
	userRepo        repository.UserRepository
	tokenManager    services.TokenManager
	passwordService services.PasswordService
}

func NewLoginUserUseCase(
	userRepo repository.UserRepository,
	tokenManager services.TokenManager,
	passwordService services.PasswordService,
) *LoginUserUseCase {
	return &LoginUserUseCase{
		userRepo:        userRepo,
		tokenManager:    tokenManager,
		passwordService: passwordService,
	}
}

func (u *LoginUserUseCase) Execute(ctx context.Context, email, password string) (string, error) {
	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", core.ErrInvalidCredentials // Ocultamos si el usuario existe o no por seguridad
	}

	if !u.passwordService.CheckPasswordHash(password, user.PasswordHash) {
		return "", core.ErrInvalidCredentials
	}

	token, err := u.tokenManager.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", core.ErrInternalServer
	}

	return token, nil
}
