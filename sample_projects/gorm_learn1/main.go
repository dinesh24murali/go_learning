package main

import "fmt"

func main() {
	db := GetDatabase()
	fmt.Println(db)

	db.AutoMigrate()

	db.Create(&Product{Code: "D42", Price: 100})

	// var product Product
	// db.First(&product, 1)

	// fmt.Println(product)
}
