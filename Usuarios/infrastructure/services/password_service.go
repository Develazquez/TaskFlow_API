package services

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/tu-usuario/taskflow/Usuarios/domain/services"
)

type passwordService struct{}

func NewPasswordService() services.PasswordService {
	return &passwordService{}
}

func (s *passwordService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *passwordService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
