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
	var sale common.Sale
	if err := c.ShouldBindJSON(&sale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.salesService.CreateSale(&sale); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sale"})
		return
	}

	c.JSON(http.StatusCreated, sale)
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
