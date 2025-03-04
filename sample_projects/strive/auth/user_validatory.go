package auth

import (
	"strive/common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RegisterUserValidator struct {
	Email    uuid.UUID `form:"email" json:"email" binding:"required,email"`
	Password uint      `form:"password" json:"password" binding:"password"`
}

func (p *RegisterUserValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, p)
	if err != nil {
		return err
	}
	return nil
}

func NewRegisterUserValidator() RegisterUserValidator {
	registerUserValidator := RegisterUserValidator{}
	return registerUserValidator
}
