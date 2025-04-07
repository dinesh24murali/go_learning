package auth

import (
	"net/http"

	"strive/common"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *UserService
}

func NewUserHandler(userService *UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	userRegisterValidator := NewRegisterUserValidator()
	if err := userRegisterValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := h.userService.RegisterUser(&userRegisterValidator.user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user common.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := h.userService.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userSerializer := UserSerializer{C: c, User: *user}

	c.JSON(http.StatusOK, userSerializer.UserResponse())
}

func (h *UserHandler) GetUserByPhone(c *gin.Context) {
	phone := c.Param("phone")
	user, err := h.userService.GetUserByPhone(phone)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	userSerializer := UserSerializer{C: c, User: *user}

	c.JSON(http.StatusOK, userSerializer.UserResponse())
}

func (h *UserHandler) Login(c *gin.Context) {

	// @todo work from here
	user, err := h.userService.GetUserByEmail("email@maildrop.com")
	// ... validate credentials ...

	// Generate JWT token
	token, err := common.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	userSerializer := UserSerializer{C: c, User: *user}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  userSerializer.UserResponse(),
	})
}
