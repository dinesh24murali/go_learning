# Question:

```
Sales, SaleDetails, Users, and Addresses
```

Here are the **HTTP handlers** for `Sales`, `SaleDetails`, `Users`, and `Addresses`. These handlers follow **RESTful API** design and assume you're using `gin-gonic/gin` for routing.

---

## **1️⃣ Sales Handler**
```go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"your_project/auth"
	"your_project/service"
)

type SalesHandler struct {
	salesService *service.SalesService
}

func NewSalesHandler(salesService *service.SalesService) *SalesHandler {
	return &SalesHandler{salesService: salesService}
}

func (h *SalesHandler) CreateSale(c *gin.Context) {
	var sale auth.Sale
	if err := c.ShouldBindJSON(&sale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.salesService.CreateSale(&sale); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sale"})
		return
	}

	c.JSON(http.StatusCreated, sale)
}

func (h *SalesHandler) GetSalesByUser(c *gin.Context) {
	userID := c.Param("userID")
	sales, err := h.salesService.GetSalesByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sales"})
		return
	}

	c.JSON(http.StatusOK, sales)
}
```

---

## **2️⃣ SaleDetails Handler**
```go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"your_project/auth"
	"your_project/service"
)

type SaleDetailsHandler struct {
	saleDetailsService *service.SaleDetailsService
}

func NewSaleDetailsHandler(service *service.SaleDetailsService) *SaleDetailsHandler {
	return &SaleDetailsHandler{saleDetailsService: service}
}

func (h *SaleDetailsHandler) CreateSaleDetail(c *gin.Context) {
	var detail auth.SaleDetails
	if err := c.ShouldBindJSON(&detail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.saleDetailsService.CreateSaleDetail(&detail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sale detail"})
		return
	}

	c.JSON(http.StatusCreated, detail)
}

func (h *SaleDetailsHandler) GetSaleDetailsBySale(c *gin.Context) {
	saleID := c.Param("saleID")
	details, err := h.saleDetailsService.GetSaleDetailsBySale(saleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sale details"})
		return
	}

	c.JSON(http.StatusOK, details)
}
```

---

## **3️⃣ Users Handler**
```go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"your_project/auth"
	"your_project/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user auth.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.RegisterUser(&user); err != nil {
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

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserByPhone(c *gin.Context) {
	phone := c.Param("phone")
	user, err := h.userService.GetUserByPhone(phone)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
```

---

## **4️⃣ Address Handler**
```go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"your_project/auth"
	"your_project/service"
)

type AddressHandler struct {
	addressService *service.AddressService
}

func NewAddressHandler(service *service.AddressService) *AddressHandler {
	return &AddressHandler{addressService: service}
}

func (h *AddressHandler) AddAddress(c *gin.Context) {
	var address auth.Address
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
```

---

## **📌 Setting Up Routes in `main.go`**
```go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"your_project/handler"
	"your_project/repository"
	"your_project/service"
)

func main() {
	db, err := gorm.Open(postgres.Open("your_connection_string"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	// Repositories
	salesRepo := repository.NewSalesRepository(db)
	saleDetailsRepo := repository.NewSaleDetailsRepository(db)
	userRepo := repository.NewUserRepository(db)
	addressRepo := repository.NewAddressRepository(db)

	// Services
	salesService := service.NewSalesService(salesRepo)
	saleDetailsService := service.NewSaleDetailsService(saleDetailsRepo)
	userService := service.NewUserService(userRepo)
	addressService := service.NewAddressService(addressRepo)

	// Handlers
	salesHandler := handler.NewSalesHandler(salesService)
	saleDetailsHandler := handler.NewSaleDetailsHandler(saleDetailsService)
	userHandler := handler.NewUserHandler(userService)
	addressHandler := handler.NewAddressHandler(addressService)

	r := gin.Default()

	// Sales routes
	r.POST("/sales", salesHandler.CreateSale)
	r.GET("/sales/user/:userID", salesHandler.GetSalesByUser)

	// Sale details routes
	r.POST("/sale-details", saleDetailsHandler.CreateSaleDetail)
	r.GET("/sale-details/:saleID", saleDetailsHandler.GetSaleDetailsBySale)

	// User routes
	r.POST("/users", userHandler.RegisterUser)
	r.GET("/users/email/:email", userHandler.GetUserByEmail)
	r.GET("/users/phone/:phone", userHandler.GetUserByPhone)

	// Address routes
	r.POST("/addresses", addressHandler.AddAddress)
	r.GET("/addresses/user/:userID", addressHandler.GetAddressesByUser)

	r.Run(":8080")
}
```

---

## **📌 Next Steps**
✅ **Would you like JWT authentication added for secured endpoints?**  
✅ **Need integration tests for these handlers?**  
✅ **Want pagination or filtering support in the repository functions?** 🚀