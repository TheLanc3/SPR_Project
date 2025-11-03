package models

type Product struct {
	Id          int64    `gorm:"primaryKey;autoIncrement;unique;column:id"`
	SupplierId  int64    `gorm:"column:supplier_id"`
	Name        string   `gorm:"column:name"`
	Description string   `gorm:"column:description"`
	Price       int      `gorm:"column:price"`
	Quantity    int      `gorm:"column:quantity"`
	Supplier    Supplier `gorm:"foreignKey:SupplierId;references:Id"`
}
