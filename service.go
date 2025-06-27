package main

import (
	"go-post-api/model"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateUser(user model.User) error
}

type UserService struct {
	repo      Repository
	validator *validator.Validate
}

func NewUserService(repo Repository, v *validator.Validate) Service {
	return &UserService{repo: repo, validator: v}
}

func (s *UserService) CreateUser(user model.User) error {
	if err := s.validator.Struct(user); err != nil {
		return err
	}
	return s.repo.CreateUser(user)
}
