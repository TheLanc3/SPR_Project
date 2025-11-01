package models

import (
	"spr-project/enums"
	"time"
)

type Shipment struct {
	Id         int64        `gorm:"primaryKey;autoIncrement;unique;column:id"`
	ProductId  int64        `gorm:"column:product_id"`
	Quantity   int          `gorm:"column:quantity"`
	SupplierId int64        `gorm:"column:supplier_id"`
	CreatedAt  time.Time    `gorm:"column:created_at;serializer:unixtime;autoCreateTime"`
	Status     enums.Status `gorm:"column:status;default:1"`
	UpdatedAt  time.Time    `gorm:"column:created_at;serializer:unixtime;autoUpdateTime"`
	Product    Product      `gorm:"foreignKey:ProductId;references:Id"`
}
