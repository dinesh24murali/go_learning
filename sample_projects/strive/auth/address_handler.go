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
	validator := NewAddressValidator()
	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("error", err))
		return
	}

	if err := h.addressService.AddAddress(&validator.address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add address"})
		return
	}

	c.JSON(http.StatusCreated, validator.address)
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

func (h *AddressHandler) UpdateAddress(c *gin.Context) {
	validator := NewAddressValidator()
	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("error", err))
		return
	}

	addressID := c.Param("id")
	if err := h.addressService.UpdateAddress(addressID, &validator.address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update address"})
		return
	}

	c.JSON(http.StatusOK, validator.address)
}

func (h *AddressHandler) DeleteAddress(c *gin.Context) {
	addressID := c.Param("id")
	if err := h.addressService.DeleteAddress(addressID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
}
