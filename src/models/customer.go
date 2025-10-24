package models

type Customer struct {
	Id      int64  `gorm:"primaryKey;autoIncrement;unique;column:id"`
	Name    string `gorm:"column:name"`
	Phone   string `gorm:"column:phone"`
	Address string `gorm:"column:address"`
}
