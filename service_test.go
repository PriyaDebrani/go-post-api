package main

import (
	"testing"

	"go-post-api/model"
	customValidator "go-post-api/validator"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestUserService_CreateUser(t *testing.T) {
	tests := []struct {
		name    string
		user    model.User
		wantErr bool
		Err     error
	}{

		{
			name: "invalid PAN :NEG",
			user: model.User{
				Name:   "User1",
				PAN:    "1234ABCDE",
				Mobile: "1234567765",
				Email:  "user1@google.com",
			},
			wantErr: true,
		},
		{
			name: "invalid mobile number :NEG",
			user: model.User{
				Name:   "User2",
				PAN:    "AAAAA1234F",
				Mobile: "1234",
				Email:  "user2@google.com",
			},
			wantErr: true,
		},
		{
			name: "invalid email :NEG",
			user: model.User{
				Name:   "User3",
				PAN:    "GHJSS1234F",
				Mobile: "111112222",
				Email:  "abc-email.com",
			},
			wantErr: true,
		},
		{
			name: "missing name :NEG",
			user: model.User{
				PAN:    "AAAAA1234F",
				Mobile: "1234432112",
				Email:  "test@example.com",
			},
			wantErr: true,
		},
		{
			name: "successfully add user :POS",
			user: model.User{
				Name:   "Priya",
				PAN:    "ABRDS2134F",
				Mobile: "2121221121",
				Email:  "priya@google.com",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validate := validator.New()
			customValidator.RegisterCustomValidations(validate)

			repo := NewInMemoryRepo()
			service := NewUserService(repo, validate)

			err := service.CreateUser(tt.user)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
