package auth

import (
	"strive/common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddressValidator struct {
	AddressLine1 string         `form:"address_line1" json:"address_line1" binding:"required,min=5,max=200"`
	AddressLine2 string         `form:"address_line2" json:"address_line2" binding:"max=200"`
	State        string         `form:"state" json:"state" binding:"required,min=2,max=150"`
	City         string         `form:"city" json:"city" binding:"required,min=2,max=150"`
	Pincode      string         `form:"pincode" json:"pincode" binding:"required,min=5,max=10"`
	Phone        string         `form:"phone" json:"phone" binding:"required,min=10,max=15"`
	address      common.Address `json:"-"`
}

func (a *AddressValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, a)
	if err != nil {
		return err
	}
	userID, _ := c.Get("userID")
	userIDTyped := uuid.MustParse(userID.(string))
	a.address.AddressLine1 = a.AddressLine1
	a.address.AddressLine2 = a.AddressLine2
	a.address.State = a.State
	a.address.City = a.City
	a.address.Pincode = a.Pincode
	a.address.Phone = a.Phone
	a.address.UserID = userIDTyped

	return nil
}

func NewAddressValidator() AddressValidator {
	addressValidator := AddressValidator{}
	return addressValidator
}
