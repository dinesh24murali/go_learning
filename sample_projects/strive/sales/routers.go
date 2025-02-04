package sales

import (
	"strive/common"

	"github.com/gin-gonic/gin"
)

func SalesRegister(router *gin.RouterGroup) {

	db := common.GetDB()

	salesRepo := NewSalesRepository(db)
	salesService := NewSalesService(salesRepo)
	salesHandler := NewSalesHandler(salesService)

	router.POST("/", salesHandler.CreateSale)
	router.GET("/user/:userID", salesHandler.GetSalesByUser)
}
