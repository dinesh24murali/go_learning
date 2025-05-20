package common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid()"`
	FirstName         string     `gorm:"column:first_name"`
	LastName          string     `gorm:"column:last_name"`
	Email             string     `gorm:"column:email;unique"`
	Phone             string     `gorm:"column:phone;unique"`
	Image             *string    `gorm:"column:image;default:null"`
	Password          string     `gorm:"column:password"`
	Status            UserStatus `gorm:"column:status;not null;default:0"`
	Role              UserRole   `gorm:"column:role;not null;default:0"`
	RefreshToken      *string    `gorm:"column:refresh_token;default:null"`
	VerificationToken *string    `gorm:"column:verification_token;default:null"`
	Addresses         []Address
	Sales             []Sale
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Address struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	AddressLine1 string    `gorm:"column:address_line1"`
	AddressLine2 string    `gorm:"column:address_line2;default:null"`
	State        string    `gorm:"column:state"`
	City         string    `gorm:"column:city"`
	Pincode      string    `gorm:"column:pincode"`
	Phone        string    `gorm:"column:phone"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       uuid.UUID `gorm:"column:user_id"`
	Sales        []Sale
}

type Category struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"column:name"`
	Products  []Product
}

type Product struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string  `gorm:"column:name"`
	Description     string  `gorm:"column:description"`
	Stock           uint    `gorm:"default:0"`
	Price           float64 `gorm:"type:decimal(10,2);check:price >= 0;not null;default:0"`
	Count           uint    `gorm:"default:1"`
	DiscountPercent uint8   `gorm:"default:0"`
	IsAvailable     bool    `gorm:"default:true"`
	IsDeleted       bool    `gorm:"column:is_deleted;default:false"`
	ImageUrl        *string `gorm:"default:null"`
	CategoryID      uuid.UUID
	Category        Category
}

type SalesAddress struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	AddressLine1 string    `gorm:"column:address_line1"`
	AddressLine2 string    `gorm:"column:address_line2;default:null"`
	State        string    `gorm:"column:state"`
	City         string    `gorm:"column:city"`
	Pincode      string    `gorm:"column:pincode"`
	Phone        string    `gorm:"column:phone"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       uuid.UUID `gorm:"column:user_id"`
}

type Sale struct {
	gorm.Model
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	UserID           uuid.UUID
	User             User `gorm:"foreignKey:UserID"`
	SalesAddressID   uuid.UUID
	SalesAddress     SalesAddress `gorm:"foreignKey:SalesAddressID"`
	Date             time.Time
	DiscountAmount   uint `gorm:"default:0"`
	PackagingCharges uint `gorm:"default:0"`
	InvoiceNetAmount uint
	Tax              uint
	Status           SalesStatus   `gorm:"default:0"`
	SaleDetails      []SaleDetails `gorm:"foreignKey:SaleID;references:ID"`
}

type SaleDetails struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	SaleID      uuid.UUID
	Sale        Sale `gorm:"foreignKey:SaleID;references:ID"`
	ProductID   uuid.UUID
	ProductName string  `gorm:"column:product_name"`
	Product     Product `gorm:"foreignKey:ProductID"`
	Quantity    int
}
