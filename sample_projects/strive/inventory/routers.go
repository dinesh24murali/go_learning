package inventory

import (
	"strive/common"

	"github.com/gin-gonic/gin"
)

func ProductRegister(router *gin.RouterGroup) {

	db := common.GetDB()

	productRepo := NewProductRepository(db)
	productService := NewProductService(productRepo)
	productHandler := NewProductHandler(productService)

	router.POST("/", productHandler.CreateProduct)
	router.GET("/", productHandler.GetProducts)
}
