package sales

import (
	"strive/auth"
	"strive/common"
	"strive/inventory"

	"github.com/gin-gonic/gin"
)

func SalesRegister(router *gin.RouterGroup) {

	db := common.GetDB()

	emailService, _ := common.NewEmailService()
	salesRepo := NewSalesRepository(db)
	userRepo := auth.NewUserRepository(db)
	addressRepo := auth.NewAddressRepository(db)
	productRepo := inventory.NewProductRepository(db)
	salesService := NewSalesService(salesRepo, userRepo, addressRepo, productRepo, emailService)
	salesHandler := NewSalesHandler(salesService)

	router.POST("/", salesHandler.CreateSale, salesHandler.SendEmail)
	router.GET("/user/:userID", salesHandler.GetSalesByUser)
}
