package sales

import (
	"fmt"
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
	// salesID, salesErr := h.salesService.CreateSale(&salesValidator.salesData)
	// if salesErr != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sale"})
	// 	return
	// }
	// c.Set("SalesID", salesID)
	c.Set("SalesID", "9a99b94d-251f-481d-9208-5cd9e8456dbf")
	c.JSON(http.StatusCreated, gin.H{"message": "Sales created successfully"})
}

func (h *SalesHandler) SendEmail(c *gin.Context) {
	temp, err := c.Get("SalesID")
	if err == false {
		return
	}
	salesID := temp.(string)
	emailError := h.salesService.SendSalesEmail(salesID)
	if emailError != nil {
		fmt.Println("Failed to send Email")
		fmt.Println(emailError)
		return
	}
	fmt.Println("Email sent successfully")
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
