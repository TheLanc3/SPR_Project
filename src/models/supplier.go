package models

type Supplier struct {
	Id        int64      `gorm:"primaryKey;autoIncrement;unique;column:id"`
	Shipments []Shipment `gorm:"foreignKey:SupplierId;references:Id"`
	Name      string     `gorm:"column:name"`
	Phone     string     `gorm:"column:phone"`
	Email     string     `gorm:"column:email"`
	Address   string     `gorm:"column:address"`
}
