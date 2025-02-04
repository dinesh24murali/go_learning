package inventory

import (
	"strive/common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductModelValidator struct {
	Product struct {
		Name            string    `form:"name" json:"name" binding:"required,min=4,max=255"`
		Description     string    `form:"description" json:"description" binding:"required,min=4,max=1000"`
		Price           float64   `form:"price" json:"price" binding:"required,min=8,max=255"`
		Count           uint      `form:"count" json:"count" binding:"required,number"`
		DiscountPercent uint8     `form:"discount_percent" json:"discount_percent" binding:"required,number,min=0,max=100"`
		IsAvailable     bool      `form:"is_available" json:"is_available" binding:"boolean"`
		CategoryID      uuid.UUID `form:"category_id" json:"category_id" binding:"required,uuid"`
	} `json:"product"`
	productModel common.Product `json:"-"`
}

func (p *ProductModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, p)
	if err != nil {
		return err
	}
	p.productModel.Name = p.Product.Name
	p.productModel.Description = p.Product.Description
	p.productModel.Price = p.Product.Price
	p.productModel.Count = p.Product.Count
	p.productModel.DiscountPercent = p.Product.DiscountPercent
	p.productModel.IsAvailable = p.Product.IsAvailable
	p.productModel.CategoryID = p.Product.CategoryID

	return nil
}

// You can put the default value of a Validator here
func NewProductModelValidator() ProductModelValidator {
	productModelValidator := ProductModelValidator{}
	// productModelValidator.Product.Description ="-"
	return productModelValidator
}
