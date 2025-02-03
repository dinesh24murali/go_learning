# Question:

```
For a schema like the following, how would you implement SOLID principles?

package auth

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                uuid.UUID  gorm:"type:uuid;default:gen_random_uuid()"
	FirstName         string     gorm:"column:first_name"
	LastName          string     gorm:"column:last_name"
	Email             string     gorm:"column:email;unique_index"
	Phone             string     gorm:"column:phone;unique_index"
	Image             *string    gorm:"column:image;default:null"
	Password          *string    gorm:"column:password"
	Status            UserStatus gorm:"column:status;not null;default:0"
	Role              UserRole   gorm:"column:role;not null;default:0"
	RefreshToken      *string    gorm:"column:refresh_token;default:null"
	VerificationToken *string    gorm:"column:verification_token;default:null"
	Addresses         []Address
	Sales             []Sale
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Address struct {
	ID           uuid.UUID gorm:"type:uuid;default:gen_random_uuid()"
	AddressLine1 string    gorm:"column:address_line1"
	AddressLine2 string    gorm:"column:address_line2;default:null"
	State        string    gorm:"column:state"
	City         string    gorm:"column:city"
	Pincode      string    gorm:"column:pincode"
	Phone        string    gorm:"column:phone"
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       uuid.UUID
	Sales        []Sale
}

type Category struct {
	ID        uuid.UUID gorm:"type:uuid;default:gen_random_uuid()"
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string gorm:"column:name"
	Products  []Product
}

type Product struct {
	ID              uuid.UUID gorm:"type:uuid;default:gen_random_uuid()"
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string gorm:"column:name"
	Description     string gorm:"column:description"
	Stock           uint
	Price           float64 gorm:"type:decimal(10,2);check:price >= 0;not null;default:0"
	Count           uint    gorm:"default:1"
	DiscountPercent uint8   gorm:"default:0"
	IsAvailable     bool    gorm:"default:true"
	ImageUrl        *string gorm:"default:null"
	CategoryID      uuid.UUID
	Category        Category
	SaleDetails     []SaleDetails
}

type Sale struct {
	ID               uuid.UUID gorm:"type:uuid;default:gen_random_uuid()"
	CreatedAt        time.Time
	UpdatedAt        time.Time
	UserID           uuid.UUID
	User             User
	AddressID        uuid.UUID
	Address          Address
	Date             time.Time
	DiscountAmount   uint gorm:"default:0"
	PackagingCharges uint gorm:"default:0"
	InvoiceNetAmount uint
	Tax              uint
	Status           SalesStatus gorm:"default:0"
	SaleDetails      []SaleDetails
}

type SaleDetails struct {
	ID        uuid.UUID gorm:"type:uuid;default:gen_random_uuid()"
	CreatedAt time.Time
	UpdatedAt time.Time
	SaleID    uuid.UUID
	Sale      Sale
	ProductID uuid.UUID
	Product   Product
	Quantity  uint
}

You can start simply by adding functions for the product and category. Use generics reuse the database functions used between different models.  
```

To implement **SOLID** principles in your schema, I'll start by adding **generic database functions** for `Product` and `Category` models while ensuring:  

- **S**ingle Responsibility Principle (**SRP**) – Each function does only one thing.  
- **O**pen/Closed Principle (**OCP**) – Code is open for extension but closed for modification.  
- **L**iskov Substitution Principle (**LSP**) – Generic repository should work for any model.  
- **I**nterface Segregation Principle (**ISP**) – Only relevant functions should be implemented per entity.  
- **D**ependency Inversion Principle (**DIP**) – Use interfaces instead of concrete implementations.  

---

### **Step 1: Define a Generic Repository Interface**  
This ensures **DIP** by abstracting the database operations.  

```go
package repository

import (
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Create(entity *T) error
	FindAll() ([]T, error)
	FindByID(id string) (*T, error)
	Update(entity *T) error
	Delete(id string) error
}
```

---

### **Step 2: Implement a Generic GORM Repository**  
This implements **SRP** and **LSP** by working for any model.  

```go
package repository

import (
	"gorm.io/gorm"
)

type GormRepository[T any] struct {
	db *gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}

func (r *GormRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *GormRepository[T]) FindAll() ([]T, error) {
	var entities []T
	err := r.db.Find(&entities).Error
	return entities, err
}

func (r *GormRepository[T]) FindByID(id string) (*T, error) {
	var entity T
	err := r.db.First(&entity, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GormRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *GormRepository[T]) Delete(id string) error {
	return r.db.Delete(new(T), "id = ?", id).Error
}
```

---

### **Step 3: Create Product and Category Repositories**  
These repositories extend the generic one and adhere to **OCP** (can add more domain-specific functions).  

#### **Product Repository**
```go
package repository

import (
	"gorm.io/gorm"
	"your_project/auth"
)

type ProductRepository struct {
	*GormRepository[auth.Product]
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		GormRepository: NewGormRepository[auth.Product](db),
	}
}

func (r *ProductRepository) FindByCategory(categoryID string) ([]auth.Product, error) {
	var products []auth.Product
	err := r.db.Where("category_id = ?", categoryID).Find(&products).Error
	return products, err
}
```

#### **Category Repository**
```go
package repository

import (
	"gorm.io/gorm"
	"your_project/auth"
)

type CategoryRepository struct {
	*GormRepository[auth.Category]
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		GormRepository: NewGormRepository[auth.Category](db),
	}
}
```

---

### **Step 4: Use the Repositories in a Service Layer**
This keeps business logic separate, following **SRP**.

#### **Product Service**
```go
package service

import (
	"your_project/auth"
	"your_project/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *auth.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) GetAllProducts() ([]auth.Product, error) {
	return s.repo.FindAll()
}
```

#### **Category Service**
```go
package service

import (
	"your_project/auth"
	"your_project/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]auth.Category, error) {
	return s.repo.FindAll()
}
```

---

### **Step 5: Dependency Injection in Handlers**
This ensures **DIP**, making it easy to swap implementations (e.g., mocking for tests).

#### **Product Handler**
```go
package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"your_project/auth"
	"your_project/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product auth.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
```

---

### **How SOLID Principles Are Applied**
✅ **SRP (Single Responsibility Principle):**  
- Each layer (repository, service, handler) has its own responsibility.  

✅ **OCP (Open/Closed Principle):**  
- The generic repository can be extended without modifying existing code.  

✅ **LSP (Liskov Substitution Principle):**  
- `GormRepository[T]` can be replaced with any struct (`Product`, `Category`) without breaking functionality.  

✅ **ISP (Interface Segregation Principle):**  
- `Repository[T]` only defines essential functions for database interactions.  

✅ **DIP (Dependency Inversion Principle):**  
- Services depend on repository **interfaces**, making them testable.  

---

### **Next Steps**
- Would you like to implement dependency injection in `main.go`?
- Do you need to add unit tests for repositories/services? 🚀