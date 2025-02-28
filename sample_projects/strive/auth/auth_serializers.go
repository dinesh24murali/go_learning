package auth

import (
	"strive/common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthSerializer struct {
	c *gin.Context
}

type MeResponse struct {
	ID        uuid.UUID         `json:"id"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Email     string            `json:"email"`
	Phone     string            `json:"phone"`
	ImageUrl  *string           `json:"image_url"`
	Status    common.UserStatus `json:"status"`
}

func (a *AuthSerializer) UserResponse() MeResponse {
	userModel := a.c.MustGet("user_model").(common.User)
	user := MeResponse{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
		Phone:     userModel.Phone,
		ImageUrl:  userModel.Image,
		Status:    userModel.Status,
	}
	return user
}
