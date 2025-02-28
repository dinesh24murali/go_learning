package main

import (
	"strive/auth"
	"strive/common"
	"strive/inventory"
	"strive/sales"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupApp() {
	db := common.Init()

	common.MigrateModels(db)

	r := gin.Default()

	common.RegisterValidators()

	envFile, _ := godotenv.Read(".env")

	v1 := r.Group("/api")

	inventory.ProductRegister(v1.Group("/product"))
	auth.UserRegister(v1.Group("/user"))
	auth.AddressRegister(v1.Group("/addresses"))
	sales.SalesRegister(v1.Group("/sales"))

	r.Run(":" + envFile["PORT"])
}
