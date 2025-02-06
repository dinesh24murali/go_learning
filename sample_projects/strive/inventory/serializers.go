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

func (p *ProductSerializer) ListResponse() []ProductResponse {
	products := p.c.MustGet("products").([]common.Product)
	response := []ProductResponse{}
	for _, v := range products {
		response = append(response, ProductResponse{
			ID:              v.ID,
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
			Name:            v.Name,
			Description:     v.Description,
			Stock:           v.Stock,
			Price:           v.Price,
			Count:           v.Count,
			DiscountPercent: v.DiscountPercent,
			IsAvailable:     v.IsAvailable,
			ImageUrl:        v.ImageUrl,
			CategoryID:      v.CategoryID,
		})
	}
	return response
}
