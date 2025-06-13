package test

import (
	"bytes"
	"mediportal/config"
	"mediportal/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// Global shared router
var sharedRouter *gin.Engine

func TestMain(m *testing.M) {
	// Load env
	if err := godotenv.Load("../.env"); err != nil {
		_ = godotenv.Load(".env")
	}

	// DB connect
	config.ConnectDatabase()

	// Setup shared router
	sharedRouter = gin.Default()
	routes.RegisterRoutes(sharedRouter)

	// Run tests
	os.Exit(m.Run())
}

func TestSignup(t *testing.T) {
	reqBody := `{"email": "test@example.com", "password": "password123", "role": "receptionist"}`
	req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	sharedRouter.ServeHTTP(w, req)

	if w.Code != http.StatusCreated && w.Code != http.StatusInternalServerError {
		t.Errorf("Unexpected status code: got %d", w.Code)
	}
}

func TestLogin(t *testing.T) {
	reqBody := `{"email": "test@example.com", "password": "password123"}`
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	sharedRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
