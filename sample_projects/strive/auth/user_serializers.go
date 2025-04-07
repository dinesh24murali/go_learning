package auth

import (
	"strive/common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserSerializer struct {
	C    *gin.Context
	User common.User
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

func (a *UserSerializer) UserResponse() MeResponse {
	userModel := a.User
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
