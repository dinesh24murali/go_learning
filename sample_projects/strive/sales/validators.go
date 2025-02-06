package sales

import (
	"strive/common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductList struct {
	ProductID uuid.UUID `form:"product_id" json:"product_id" binding:"required,uuid"`
	Quantity  uint      `form:"qty" json:"qty" binding:"required,number,min=1"`
}

type SalesValidator struct {
	Sale struct {
		UserID       uuid.UUID      `form:"user_id" json:"user_id" binding:"required,uuid"`
		AddressID    uuid.UUID      `form:"address_id" json:"address_id" binding:"required,uuid"`
		SalesDetails *[]ProductList `json:"products"`
	} `json:"sale"`
	salesData PartialSale `json:"-"`
}

func (p *SalesValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, p)
	if err != nil {
		return err
	}
	p.salesData.UserID = p.Sale.UserID
	p.salesData.AddressID = p.Sale.AddressID
	p.salesData.SalesDetails = p.Sale.SalesDetails

	return nil
}

func NewSalesValidator() SalesValidator {
	productModelValidator := SalesValidator{}
	return productModelValidator
}
