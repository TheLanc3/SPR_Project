package database

import (
	"fmt"
	"spr-project/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Database initialization failed")
	}

	err = db.AutoMigrate(
		&models.Customer{},
		&models.Item{},
		&models.Order{},
		&models.Product{},
		&models.Shipment{},
		&models.Stock{},
		&models.Supplier{},
	)
	if err != nil {
		panic("Database initialization failed: ")
	}

	return db
}
