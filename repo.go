package main

import "go-post-api/model"

type Repository interface {
	CreateUser(user model.User) error
}

type InMemoryRepo struct {
	users []model.User
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		users: []model.User{},
	}
}

func (r *InMemoryRepo) CreateUser(user model.User) error {
	r.users = append(r.users, user)
	return nil
}
