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

	// Public routes (no authentication required)
	auth.AuthRegister(v1.Group("/auth")) // For login/register endpoints

	// Protected routes (require authentication)
	protected := v1.Group("")
	protected.Use(common.AuthMiddleware())
	{
		inventory.ProductRegister(protected.Group("/product"))
		auth.UserRegister(protected.Group("/user"))
		auth.AddressRegister(protected.Group("/address"))
		sales.SalesRegister(protected.Group("/sales"))
	}

	// Admin-only routes
	// adminRoutes := protected.Group("")
	// adminRoutes.Use(common.AdminMiddleware())
	// {
	// 	// Add admin-specific routes here
	// }
	r.Run(":" + envFile["PORT"])
}
