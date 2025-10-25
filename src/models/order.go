package models

import (
	"spr-project/enums"
	"time"
)

type Order struct {
	Id         int64        `gorm:"primaryKey;autoIncrement;unique;column:id"`
	CustomerId int64        `gorm:"column:customer_id"`
	Customer   Customer     `gorm:"foreignKey:CustomerId;references:Id"`
	Positions  []Item       `gorm:"foreignKey:OrderId;references:Id"`
	Total      int          `gorm:"column:total"`
	CreatedAt  time.Time    `gorm:"column:created_at;serializer:unixtime;autoCreateTime"`
	Status     enums.Status `gorm:"column:status"`
	UpdatedAt  time.Time    `gorm:"column:created_at;serializer:unixtime;autoUpdateTime"`
}
