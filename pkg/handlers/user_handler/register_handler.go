package user_handler

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type ReceivedData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user, err := h.decodeUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Sprawdź, czy użytkownik o podanej nazwie użytkownika już istnieje
	existingUser, err := h.UserRepository.FindByUsername(user.Username)
	if err == nil && existingUser != nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// Haszowanie hasła - zakładam, że posiadasz odpowiednią funkcję
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	err = h.UserRepository.RegisterUser(user)
	if err != nil {
		http.Error(w, "Unable to register user", http.StatusInternalServerError)
		return
	}

	// Po rejestracji możesz natychmiast zalogować użytkownika i wygenerować dla niego token
	tokenStr, err := h.TokenService.GenerateToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenStr})
}
