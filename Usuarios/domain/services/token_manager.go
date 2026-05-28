package services

type TokenManager interface {
	GenerateToken(userID int, email string) (string, error)
}
