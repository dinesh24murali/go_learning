package common

import "gorm.io/gorm"

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Address{})
	db.AutoMigrate(&Sale{})
	db.AutoMigrate(&SaleDetails{})
}
