package models

import "spr-project/enums"

type Supplier struct {
	Id       int64        `gorm:"primaryKey;autoIncrement;unique;column:id"`
	Shipmens Shipment     `gorm:"foreignKey:SupplierId;references:Id"`
	Status   enums.Status `gorm:"column:status"`
	Name     string       `gorm:"column:name"`
	Phone    string       `gorm:"column:phone"`
	Email    string       `gorm:"column:email"`
	Address  string       `gorm:"column:address"`
}
