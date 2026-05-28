package services

type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
