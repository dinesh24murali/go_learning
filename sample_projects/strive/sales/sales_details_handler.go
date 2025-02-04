package sales

import (
	"net/http"

	"strive/common"

	"github.com/gin-gonic/gin"
)

type SaleDetailsHandler struct {
	saleDetailsService *SaleDetailsService
}

func NewSaleDetailsHandler(service *SaleDetailsService) *SaleDetailsHandler {
	return &SaleDetailsHandler{saleDetailsService: service}
}

func (h *SaleDetailsHandler) CreateSaleDetail(c *gin.Context) {
	var detail common.SaleDetails
	if err := c.ShouldBindJSON(&detail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.saleDetailsService.CreateSaleDetail(&detail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sale detail"})
		return
	}

	c.JSON(http.StatusCreated, detail)
}

func (h *SaleDetailsHandler) GetSaleDetailsBySale(c *gin.Context) {
	saleID := c.Param("saleID")
	details, err := h.saleDetailsService.GetSaleDetailsBySale(saleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sale details"})
		return
	}

	c.JSON(http.StatusOK, details)
}
