package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tu-usuario/taskflow/Usuarios/domain/services"
)

type jwtManager struct {
	secretKey string
}

func NewJWTManager(secretKey string) services.TokenManager {
	return &jwtManager{
		secretKey: secretKey,
	}
}

func (m *jwtManager) GenerateToken(userID int, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(m.secretKey))
}
