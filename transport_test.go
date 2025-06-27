package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-post-api/middleware"
	fieldValidator "go-post-api/validator"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func setupRouterForTest() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(middleware.LatencyLogger())

	v := validator.New()
	fieldValidator.RegisterCustomValidations(v)

	repo := NewInMemoryRepo()
	svc := NewUserService(repo, v)
	h := NewTransport(svc)

	router.POST("/users", h.CreateUser)
	return router
}

func TestCreateUserHandler(t *testing.T) {
	router := setupRouterForTest()

	tests := []struct {
		name           string
		payload        string
		statusCode     int
		expectedField  string
		expectedErrMsg string
		successMsg     string
	}{
		{
			name: "invalid PAN :NEG",
			payload: `{
				"name": "User1",
				"pan": "1234ABCDE",
				"mobile": "1234567765",
				"email": "user1@google.com"
			}`,
			statusCode:     http.StatusBadRequest,
			expectedField:  "PAN",
			expectedErrMsg: "Invalid PAN format",
		},
		{
			name: "invalid email :NEG",
			payload: `{
				"name": "User2",
				"pan": "AAAAA1234Z",
				"mobile": "1234567890",
				"email": "abc-email.com"
			}`,
			statusCode:     http.StatusBadRequest,
			expectedField:  "Email",
			expectedErrMsg: "Invalid email format",
		},
		{
			name: "invalid mobile number :NEG",
			payload: `{
				"name": "User2",
				"pan": "AAAAA1234Z",
				"mobile": "12345",
				"email": "abc@email.com"
			}`,
			statusCode:     http.StatusBadRequest,
			expectedField:  "Mobile",
			expectedErrMsg: "Mobile must be a 10 digit number",
		},
		{
			name: "missing name :NEG",
			payload: `{
				"pan": "AAAAA1234F",
				"mobile": "1234567890",
				"email": "user1@google.com"
			}`,
			statusCode:     http.StatusBadRequest,
			expectedField:  "Name",
			expectedErrMsg: "Name is required",
		},
		{
			name:           "invalid json :NEG",
			payload:        `{invalid}`,
			statusCode:     http.StatusBadRequest,
			expectedField:  "",
			expectedErrMsg: "Invalid JSON",
		},
		{
			name: "valid json input :POS",
			payload: `{
				"name": "Priya",
				"pan": "ABCDE1234F",
				"mobile": "9876543210",
				"email": "priya@google.com"
			}`,
			statusCode: http.StatusOK,
			successMsg: "User created successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(tt.payload))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.statusCode, resp.Code)

			var responseBody map[string]interface{}
			err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
			assert.NoError(t, err)

			if tt.statusCode == http.StatusOK {
				assert.Equal(t, tt.successMsg, responseBody["message"])
			} else {
				if tt.expectedField != "" {
					errorsMap := responseBody["errors"].(map[string]interface{})
					assert.Contains(t, errorsMap, tt.expectedField)
					assert.Equal(t, tt.expectedErrMsg, errorsMap[tt.expectedField])
				} else {
					assert.Equal(t, tt.expectedErrMsg, responseBody["error"])
				}
			}
		})
	}
}
