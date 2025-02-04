package inventory

import (
	"strive/common"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductSerializer struct {
	c *gin.Context
}

type ProductResponse struct {
	ID              uuid.UUID `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Stock           uint      `json:"stock"`
	Price           float64   `json:"price"`
	Count           uint      `json:"count"`
	DiscountPercent uint8     `json:"discount_percent"`
	IsAvailable     bool      `json:"is_available"`
	ImageUrl        *string   `json:"image_url"`
	CategoryID      uuid.UUID `json:"category_id"`
}

func (p *ProductSerializer) Response() ProductResponse {
	productModel := p.c.MustGet("product_model").(common.Product)
	user := ProductResponse{
		ID:              productModel.ID,
		CreatedAt:       productModel.CreatedAt,
		UpdatedAt:       productModel.UpdatedAt,
		Name:            productModel.Name,
		Description:     productModel.Description,
		Stock:           productModel.Stock,
		Price:           productModel.Price,
		Count:           productModel.Count,
		DiscountPercent: productModel.DiscountPercent,
		IsAvailable:     productModel.IsAvailable,
		ImageUrl:        productModel.ImageUrl,
		CategoryID:      productModel.CategoryID,
	}
	return user
}
