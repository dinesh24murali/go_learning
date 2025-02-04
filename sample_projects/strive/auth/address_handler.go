package auth

import (
	"net/http"

	"strive/common"

	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	addressService *AddressService
}

func NewAddressHandler(service *AddressService) *AddressHandler {
	return &AddressHandler{addressService: service}
}

func (h *AddressHandler) AddAddress(c *gin.Context) {
	var address common.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.addressService.AddAddress(&address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add address"})
		return
	}

	c.JSON(http.StatusCreated, address)
}

func (h *AddressHandler) GetAddressesByUser(c *gin.Context) {
	userID := c.Param("userID")
	addresses, err := h.addressService.GetAddressesByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch addresses"})
		return
	}

	c.JSON(http.StatusOK, addresses)
}
