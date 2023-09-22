package user_handler

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"stock-simulation/pkg/model"
)

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	RegisterUser(user *model.User) error
}

type TokenService interface {
	GenerateToken(user *model.User) (string, error)
}

type HashService interface {
	HashPassword(password string) string
	VerifyPassword(hashedPassword, password string) bool
}

type Handler struct {
	UserRepository UserRepository
	TokenService   TokenService
}

func (h *Handler) decodeUser(r *http.Request) (*model.User, error) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user, err := h.decodeUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	databaseUser, err := h.UserRepository.FindByUsername(user.Username)
	if err != nil || databaseUser == nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// To sprawdzenie zapobiega odwołaniu się do nil
	if databaseUser == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Verify the password
	if !VerifyPassword(databaseUser.Password, user.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	tokenStr, err := h.TokenService.GenerateToken(databaseUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenStr})
}

func VerifyPassword(password string, passwordRequest string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(passwordRequest))
	return err == nil
}
