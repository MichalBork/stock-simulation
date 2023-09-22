package service

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"stock-simulation/pkg/model"
)

type TokenService struct {
	Secret string
}

func NewTokenService(secret string) *TokenService {
	return &TokenService{Secret: secret}
}

func (s *TokenService) GenerateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Subject:   strconv.Itoa(user.ID),
	})
	return token.SignedString([]byte(s.Secret))
}
