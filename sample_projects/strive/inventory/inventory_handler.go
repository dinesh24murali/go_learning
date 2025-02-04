package inventory

import (
	"net/http"
	"strive/common"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *ProductService
}

func NewProductHandler(service *ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {

	productModelValidator := NewProductModelValidator()
	if err := productModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	// var product common.Product
	// if err := c.ShouldBindJSON(&product); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	if err := h.service.CreateProduct(&productModelValidator.productModel); err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("product_model", productModelValidator.productModel)
	serializer := ProductSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

// func (h *ProductHandler) GetProducts(c *gin.Context) {
// 	products, err := h.service.GetAllProducts()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, products)
// }

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusOK, products)
}
