package sales

import (
	"net/http"

	"strive/common"

	"github.com/gin-gonic/gin"
)

type SalesHandler struct {
	salesService *SalesService
}

func NewSalesHandler(salesService *SalesService) *SalesHandler {
	return &SalesHandler{salesService: salesService}
}

func (h *SalesHandler) CreateSale(c *gin.Context) {

	salesValidator := NewSalesValidator()

	if err := salesValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := h.salesService.CreateSale(&salesValidator.salesData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sale"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Sales created successfully"})
}

func (h *SalesHandler) GetSalesByUser(c *gin.Context) {
	userID := c.Param("userID")
	sales, err := h.salesService.GetSalesByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sales"})
		return
	}

	c.JSON(http.StatusOK, sales)
}
