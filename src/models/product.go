package models

type Product struct {
	Id          int64  `gorm:"primaryKey;autoIncrement;unique;column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Price       int    `gorm:"column:price"`
	Quantity    int    `gorm:"column:quantity"`
}
