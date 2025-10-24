package models

import "spr-project/enums"

type Order struct {
	Id         int64        `gorm:"primaryKey;autoIncrement;unique;column:id"`
	CustomerId int64        `gorm:"column:customer_id"`
	Positions  string       `gorm:"column:positions"`
	Total      int          `gorm:"column:total"`
	CreatedAt  int64        `gorm:"column:created_at;serializer:unixtime;type:time"`
	Status     enums.Status `gorm:"column:status"`
	UpdatedAt  int64        `gorm:"column:created_at;serializer:unixtime;type:time"`
}
