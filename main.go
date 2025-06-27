package main

import (
	"go-post-api/middleware"
	fieldValidator "go-post-api/validator"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LatencyLogger())

	validate := validator.New()
	fieldValidator.RegisterCustomValidations(validate)

	repo := NewInMemoryRepo()
	service := NewUserService(repo, validate)
	transport := NewTransport(service)

	transport.BuildRoutes(r)

	r.Run(":8080")
}
