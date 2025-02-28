package auth

import (
	"strive/common"

	"github.com/gin-gonic/gin"
)

func UserRegister(router *gin.RouterGroup) {

	db := common.GetDB()

	userRepo := NewUserRepository(db)
	userService := NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	router.POST("/", userHandler.RegisterUser)
	router.PUT("/", userHandler.UpdateUser)
	router.GET("/email/:email", userHandler.GetUserByEmail)
	router.GET("/phone/:phone", userHandler.GetUserByEmail)
}

func AddressRegister(router *gin.RouterGroup) {

	db := common.GetDB()

	addressRepo := NewAddressRepository(db)
	addressServer := NewAddressService(addressRepo)
	addressHandler := NewAddressHandler(addressServer)

	router.POST("/", addressHandler.AddAddress)
	router.GET("/user/:userID", addressHandler.GetAddressesByUser)
}
