package main

import (
	"fmt"
	"strive/auth"
	"strive/common"
)

func main() {
	db := common.GetDatabase()
	fmt.Println(db)

	auth.MigrateModels(db)

	// db.Create(&auth.Category{Name: "Home & Kitchen"})

	// Read Category
	//
	// var category auth.Category
	// db.First(&category, "ID = ?", "437a8406-ebff-4faa-89df-0d3a12481115")

	// Add Products
	//
	// var temp uuid.UUID = uuid.MustParse("437a8406-ebff-4faa-89df-0d3a12481115")
	// db.Create(&auth.Product{
	// 	Name:            "Toothbrush",
	// 	Description:     "Used to brush teeth",
	// 	Stock:           100,
	// 	Price:           100,
	// 	Count:           1,
	// 	DiscountPercent: 0,
	// 	IsAvailable:     true,
	// 	CategoryID:      temp,
	// 	Category:      	 category,
	// })

	// Add User
	//
	// password := "Password"
	// db.Create(&auth.User{
	// 	FirstName: "Jim",
	// 	LastName:  "Moriarty",
	// 	Email:     "jim@moriarty.com",
	// 	Phone:     "9876545434",
	// 	Password:  &password,
	// 	Status:    auth.Active,
	// 	Role:      auth.Admin,
	// })

	// Get user
	var user auth.User
	db.First(&user, "ID = ?", "fa6bdb39-aa37-4ec6-8b08-b849f82e3293")
	//
	// Add an address for the user
	//
	// db.Create(&auth.Address{
	// 	AddressLine1: "No. 12 Govinda street",
	// 	State:        "Karnataka",
	// 	City:         "Bangalore",
	// 	Pincode:      "560102",
	// 	Phone:        "3454565676",
	// 	UserID:       user.ID,
	// })

	// Get list of users along with Addresses
	//
	// var users []auth.User
	// err := db.Model(&auth.User{}).Preload("Addresses").Find(&users).Error
	// fmt.Println(err)
	// fmt.Println(users)

	var address auth.Address
	db.First(&address, "ID = ?", "9d0f95dc-25c8-44a0-8786-6ffea53dd88a")

	// Add Sale
	//
	// db.Create(&auth.Sale{
	// 	User:             user,
	// 	Address:          address,
	// 	Date:             time.Now(),
	// 	InvoiceNetAmount: 100,
	// 	Tax:              18,
	// })

	var sale auth.Sale
	db.First(&sale, "ID = ?", "598f4270-3db5-4a25-9d52-4850b4449efa")

	var product auth.Product
	db.First(&product, "ID = ?", "bc547a47-4a87-495f-9b0b-c78fc814e88a")

	// Add Sales Details
	//
	db.Create(&auth.SaleDetails{
		Product:  product,
		Sale:     sale,
		Quantity: 2,
	})
}
