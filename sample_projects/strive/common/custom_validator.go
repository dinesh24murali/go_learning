package common

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const passwordRegex = `^(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$`

// Custom password validator
func passwordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	re := regexp.MustCompile(passwordRegex)
	return re.MatchString(password)
}

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", passwordValidator)
	}
}
