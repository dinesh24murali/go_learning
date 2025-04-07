package auth

import (
	"strive/common"

	"github.com/gin-gonic/gin"
)

type RegisterUserValidator struct {
	Phone    string          `form:"phone" json:"phone" binding:"required,min=10,max=15"`
	Email    string          `form:"email" json:"email" binding:"required,email"`
	Password string          `form:"password" json:"password" binding:"password"`
	user     RegisterUserDto `json:"-"`
}

type UpdateUserValidator struct {
	Phone    string          `form:"phone" json:"phone" binding:"required,min=10,max=15"`
	Email    string          `form:"email" json:"email" binding:"required,email"`
	Password string          `form:"password" json:"password" binding:"password"`
	user     RegisterUserDto `json:"-"`
}

func (p *RegisterUserValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, p)
	if err != nil {
		return err
	}
	p.user.Phone = p.Phone
	p.user.Email = p.Email
	p.user.Password = p.Password
	return nil
}

func NewRegisterUserValidator() RegisterUserValidator {
	registerUserValidator := RegisterUserValidator{}
	return registerUserValidator
}

func NewUpdateUserValidator() UpdateUserValidator {
	updateUserValidator := UpdateUserValidator{}
	return updateUserValidator
}
