package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
	"stock-simulation/pkg/model"
	"strconv"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	t.Parallel()

	// Test secret for token generation
	testSecret := "testSecret"

	// Test user
	user := &model.User{
		ID:       1,
		Username: "TestUser2",
		Password: "TestPassword",
		Email:    "test@example.com",
	}

	tokenService := NewTokenService(testSecret)
	token, err := tokenService.GenerateToken(user)

	require.NoError(t, err, "There should be no error while generating token.")

	parsedToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(testSecret), nil
	})

	require.NoError(t, err, "There should be no error when parsing the token.")
	require.NotNil(t, parsedToken, "The parsed token should not be nil.")
	require.True(t, parsedToken.Valid, "The parsed token should be valid.")

	claims, ok := parsedToken.Claims.(*jwt.StandardClaims)

	require.True(t, ok, "The token's claims should be of type StandardClaims.")
	require.Equal(t, strconv.Itoa(user.ID), claims.Subject, "The token subject should equal the user's id.")
}
