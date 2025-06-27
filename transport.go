package main

import (
	"go-post-api/model"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func getValidationErrors(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			field := fieldErr.Field()
			switch fieldErr.Tag() {
			case "required":
				errors[field] = field + " is required"
			case "email":
				errors[field] = "Invalid email format"
			case "pan":
				errors[field] = "Invalid PAN format"
			case "mobile":
				errors[field] = "Mobile must be a 10 digit number"
			default:
				errors[field] = "Invalid value"
			}
		}
	}
	return errors
}

type Transport struct {
	service Service
}

func NewTransport(service Service) *Transport {
	return &Transport{service: service}
}

func (h *Transport) BuildRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	userGroup.POST("", h.CreateUser)
}

func (h *Transport) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := h.service.CreateUser(user); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{"errors": getValidationErrors(validationErrors)})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
