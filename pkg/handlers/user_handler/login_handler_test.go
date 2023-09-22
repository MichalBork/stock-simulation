package user_handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"stock-simulation/pkg/model"
	"testing"
)

type MockUserRepository struct{}

func (m *MockUserRepository) FindByUsername(username string) (*model.User, error) {
	return &model.User{
		ID:       1,
		Username: "TestUser2",
		Password: "TestPassword",
		Email:    "test@example.com",
	}, nil
}

func TestLoginHandler(t *testing.T) {
	user := &model.User{
		Username: "TestUser",
		Password: "TestPassword",
	}

	reqBody, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqBody))
	rr := httptest.NewRecorder()

	h := &Handler{
		UserRepository: &MockUserRepository{},
	}

	h.LoginHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)

	if _, ok := response["token"]; !ok {
		t.Errorf("handler returned unexpected body: token key not found")
	}
}
