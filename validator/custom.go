package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidations(v *validator.Validate) {
	v.RegisterValidation("pan", func(fl validator.FieldLevel) bool {
		pattern := `^[A-Z]{5}[0-9]{4}[A-Z]$`
		validPAN, _ := regexp.MatchString(pattern, fl.Field().String())
		return validPAN
	})

	v.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		pattern := `^[0-9]{10}$`
		validMobile, _ := regexp.MatchString(pattern, fl.Field().String())
		return validMobile
	})
}
