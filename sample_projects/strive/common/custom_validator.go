package common

import (
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Custom password validator
func passwordValidator(fl validator.FieldLevel) bool {
	fmt.Print("God!!")
	password := fl.Field().String()

	// Check for at least one number
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	// Check for at least one special character
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)

	return hasNumber && hasSpecial
}

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", passwordValidator)
	}
}
