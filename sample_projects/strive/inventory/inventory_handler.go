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

	if err := h.service.CreateProduct(&productModelValidator.productModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("product_model", productModelValidator.productModel)
	serializer := ProductSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"data": serializer.Response()})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("ID")

	updateProductModelValidator := NewUpdateProductModelValidator()
	if err := updateProductModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := h.service.UpdateProduct(id, &updateProductModelValidator.productModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("product_model", updateProductModelValidator.productModel)
	serializer := ProductSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"data": serializer.Response()})
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("products", products)
	serializer := ProductSerializer{c}
	c.JSON(http.StatusOK, gin.H{"data": serializer.ListResponse()})
}
