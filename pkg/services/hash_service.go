package service

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptHashService struct{}

func (s *BcryptHashService) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
