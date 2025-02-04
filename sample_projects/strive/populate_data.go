package main

import (
	"strive/common"
)

func setupData() {
	db := common.Init()

	common.MigrateModels(db)

	// db.Create(&auth.Category{Name: "Home & Kitchen"})

	// Read Category
	//
	// var category auth.Category
	// db.First(&category, "ID = ?", "597d8a8c-5a3e-4b3a-aa41-9e8d39b54fdc")

	// Add Products
	//
	// var temp uuid.UUID = uuid.MustParse("597d8a8c-5a3e-4b3a-aa41-9e8d39b54fdc")
	// db.Create(&auth.Product{
	// 	Name:            "Bowl",
	// 	Description:     "Used to hold things",
	// 	Stock:           100,
	// 	Price:           100,
	// 	Count:           1,
	// 	DiscountPercent: 0,
	// 	IsAvailable:     true,
	// 	// CategoryID:      temp,
	// 	Category: category,
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
	// var user auth.User
	// db.First(&user, "ID = ?", "24522356-a867-4024-9e8a-f945d95982a6")
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

	// var address auth.Address
	// db.First(&address, "ID = ?", "35d9604a-6606-4de3-a388-918b4fd67563")

	// Add Sale
	//
	// db.Create(&auth.Sale{
	// 	User:             user,
	// 	Address:          address,
	// 	Date:             time.Now(),
	// 	InvoiceNetAmount: 100,
	// 	Tax:              18,
	// })

	// var sale auth.Sale
	// db.First(&sale, "ID = ?", "b936a017-11cc-4204-a9f9-31029d0aabab")

	// var product auth.Product
	// db.First(&product, "ID = ?", "aca0cc7b-7ca9-41f1-8d4c-29052589520d")

	// Add Sales Details
	//
	// db.Create(&auth.SaleDetails{
	// 	Product:  product,
	// 	Sale:     sale,
	// 	Quantity: 2,
	// })
}
