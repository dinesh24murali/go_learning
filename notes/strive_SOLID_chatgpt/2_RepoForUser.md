# Question:

```
Can you give me a similar implementation for Sales, SalesDetails, Users, and Address. Let's assume that the generic Repository we created above is already in the system.
```

Here’s how you can implement repositories and services for `Sales`, `SaleDetails`, `Users`, and `Addresses`, assuming the **generic repository** is already in place.

---

## **1️⃣ Sales Repository & Service**

### **Sales Repository**
```go
package repository

import (
	"gorm.io/gorm"
	"your_project/auth"
)

type SalesRepository struct {
	*GormRepository[auth.Sale]
}

func NewSalesRepository(db *gorm.DB) *SalesRepository {
	return &SalesRepository{
		GormRepository: NewGormRepository[auth.Sale](db),
	}
}

func (r *SalesRepository) FindByUserID(userID string) ([]auth.Sale, error) {
	var sales []auth.Sale
	err := r.db.Where("user_id = ?", userID).Find(&sales).Error
	return sales, err
}
```

### **Sales Service**
```go
package service

import (
	"your_project/auth"
	"your_project/repository"
)

type SalesService struct {
	repo *repository.SalesRepository
}

func NewSalesService(repo *repository.SalesRepository) *SalesService {
	return &SalesService{repo: repo}
}

func (s *SalesService) CreateSale(sale *auth.Sale) error {
	return s.repo.Create(sale)
}

func (s *SalesService) GetAllSales() ([]auth.Sale, error) {
	return s.repo.FindAll()
}

func (s *SalesService) GetSalesByUser(userID string) ([]auth.Sale, error) {
	return s.repo.FindByUserID(userID)
}
```

---

## **2️⃣ SaleDetails Repository & Service**

### **SaleDetails Repository**
```go
package repository

import (
	"gorm.io/gorm"
	"your_project/auth"
)

type SaleDetailsRepository struct {
	*GormRepository[auth.SaleDetails]
}

func NewSaleDetailsRepository(db *gorm.DB) *SaleDetailsRepository {
	return &SaleDetailsRepository{
		GormRepository: NewGormRepository[auth.SaleDetails](db),
	}
}

func (r *SaleDetailsRepository) FindBySaleID(saleID string) ([]auth.SaleDetails, error) {
	var details []auth.SaleDetails
	err := r.db.Where("sale_id = ?", saleID).Find(&details).Error
	return details, err
}
```

### **SaleDetails Service**
```go
package service

import (
	"your_project/auth"
	"your_project/repository"
)

type SaleDetailsService struct {
	repo *repository.SaleDetailsRepository
}

func NewSaleDetailsService(repo *repository.SaleDetailsRepository) *SaleDetailsService {
	return &SaleDetailsService{repo: repo}
}

func (s *SaleDetailsService) CreateSaleDetail(detail *auth.SaleDetails) error {
	return s.repo.Create(detail)
}

func (s *SaleDetailsService) GetSaleDetailsBySale(saleID string) ([]auth.SaleDetails, error) {
	return s.repo.FindBySaleID(saleID)
}
```

---

## **3️⃣ Users Repository & Service**

### **Users Repository**
```go
package repository

import (
	"gorm.io/gorm"
	"your_project/auth"
)

type UserRepository struct {
	*GormRepository[auth.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		GormRepository: NewGormRepository[auth.User](db),
	}
}

func (r *UserRepository) FindByEmail(email string) (*auth.User, error) {
	var user auth.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByPhone(phone string) (*auth.User, error) {
	var user auth.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
```

### **Users Service**
```go
package service

import (
	"your_project/auth"
	"your_project/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user *auth.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUserByEmail(email string) (*auth.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) GetUserByPhone(phone string) (*auth.User, error) {
	return s.repo.FindByPhone(phone)
}
```

---

## **4️⃣ Address Repository & Service**

### **Address Repository**
```go
package repository

import (
	"gorm.io/gorm"
	"your_project/auth"
)

type AddressRepository struct {
	*GormRepository[auth.Address]
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{
		GormRepository: NewGormRepository[auth.Address](db),
	}
}

func (r *AddressRepository) FindByUserID(userID string) ([]auth.Address, error) {
	var addresses []auth.Address
	err := r.db.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}
```

### **Address Service**
```go
package service

import (
	"your_project/auth"
	"your_project/repository"
)

type AddressService struct {
	repo *repository.AddressRepository
}

func NewAddressService(repo *repository.AddressRepository) *AddressService {
	return &AddressService{repo: repo}
}

func (s *AddressService) AddAddress(address *auth.Address) error {
	return s.repo.Create(address)
}

func (s *AddressService) GetAddressesByUser(userID string) ([]auth.Address, error) {
	return s.repo.FindByUserID(userID)
}
```

---

## **💡 Next Steps**
- Would you like me to add **handlers** for these services?  
- Do you need **unit tests** for repositories and services?  
- Do you want **dependency injection** setup in `main.go`? 🚀