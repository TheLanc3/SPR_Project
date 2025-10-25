package models

type Shipment struct {
	Id         int64   `gorm:"primaryKey;autoIncrement;unique;column:id"`
	ProductId  int64   `gorm:"column:product_id"`
	Quantity   int     `gorm:"column:quantity"`
	SupplierId int64   `gorm:"column:supplier_id"`
	Product    Product `gorm:"foreignKey:ProductId;references:Id"`
}
