package controllers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jieqiboh/sothea_backend/entities"
	"github.com/jieqiboh/sothea_backend/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Success - 200 OK
// Bad Request - 400 Bad Request
// Unauthorized - 401 Unauthorized
func TestLogin_Success(t *testing.T) {
	var mockUsecase mocks.LoginUseCase

	user := entities.User{
		Username: "admin",
		Password: "admin",
	}

	token := `"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTc1MjMyMjMsInVzZXJuYW1lIjoiYWRtaW4ifQ.54SU6KNtfcwWWmQOniTBraN0Svo1WTQzz_3Y-W6N24A"`
	mockUsecase.On("Login", context.Background(), user).Return(token, nil)
	router := gin.Default()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"username":"admin","password":"admin"}`))

	NewLoginHandler(router, &mockUsecase)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// Parse the response body
	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the token
	assert.Equal(t, token, resp["token"])
}

func TestLogin_Failure_ValidationError(t *testing.T) {
	var mockUsecase mocks.LoginUseCase

	user := entities.User{
		Username: "admin",
		Password: "admin",
	}

	token := `"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTc1MjMyMjMsInVzZXJuYW1lIjoiYWRtaW4ifQ.54SU6KNtfcwWWmQOniTBraN0Svo1WTQzz_3Y-W6N24A"`
	mockUsecase.On("Login", context.Background(), user).Return(token, nil)
	router := gin.Default()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"username":"admin"}`)) // missing password field

	NewLoginHandler(router, &mockUsecase)

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestLogin_Failure_JSONError(t *testing.T) {
	var mockUsecase mocks.LoginUseCase

	user := entities.User{
		Username: "admin",
		Password: "admin",
	}

	token := `"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTc1MjMyMjMsInVzZXJuYW1lIjoiYWRtaW4ifQ.54SU6KNtfcwWWmQOniTBraN0Svo1WTQzz_3Y-W6N24A"`
	mockUsecase.On("Login", context.Background(), user).Return(token, nil)
	router := gin.Default()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"username": 5, "password": "test"}`)) // Wrong data type results in JSON marshalling error

	NewLoginHandler(router, &mockUsecase)

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestLogin_Failure_Unauthorized(t *testing.T) {
	var mockUsecase mocks.LoginUseCase

	user := entities.User{
		Username: "admin",
		Password: "wrongpassword",
	}

	mockUsecase.On("Login", context.Background(), user).Return("", entities.ErrLoginFailed)
	router := gin.Default()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"username": "admin", "password": "wrongpassword"}`)) // Wrong password

	NewLoginHandler(router, &mockUsecase)

	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}
