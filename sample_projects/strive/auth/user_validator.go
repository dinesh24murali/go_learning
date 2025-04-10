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

type UpdateUserValidator struct {
	FirstName string        `form:"first_name" json:"first_name" binding:"min=1,max=200"`
	LastName  string        `form:"last_name" json:"last_name" binding:"min=1,max=200"`
	Status    uint          `form:"status" json:"status" binding:"min=1"`
	user      UpdateUserDto `json:"-"`
}

func (p *UpdateUserValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, p)
	if err != nil {
		return err
	}
	p.user.FirstName = p.FirstName
	p.user.LastName = p.LastName
	p.user.Status = p.Status
	return nil
}

func NewUpdateUserValidator() UpdateUserValidator {
	updateUserValidator := UpdateUserValidator{}
	return updateUserValidator
}
