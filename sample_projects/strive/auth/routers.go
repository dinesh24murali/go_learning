package auth

import (
	"strive/common"

	"github.com/gin-gonic/gin"
)

func AuthRegister(router *gin.RouterGroup) {

	db := common.GetDB()

	userRepo := NewUserRepository(db)
	userService := NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	router.POST("/login", userHandler.Login)
	router.POST("/register", userHandler.RegisterUser)
}

func UserRegister(router *gin.RouterGroup) {

	db := common.GetDB()

	userRepo := NewUserRepository(db)
	userService := NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	router.PUT("/", userHandler.UpdateUser)
	router.GET("/me", userHandler.GetUser)
	router.GET("/phone/:phone", userHandler.GetUserByPhone)
}

func AddressRegister(router *gin.RouterGroup) {

	db := common.GetDB()

	addressRepo := NewAddressRepository(db)
	addressServer := NewAddressService(addressRepo)
	addressHandler := NewAddressHandler(addressServer)

	router.POST("/", addressHandler.AddAddress)
	router.GET("/user/:userID", addressHandler.GetAddressesByUser)
	router.PUT("/address/:id", addressHandler.UpdateAddress)
	router.DELETE("/address/:id", addressHandler.UpdateAddress)
}
