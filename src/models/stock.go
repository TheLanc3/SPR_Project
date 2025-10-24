package models

type Stock struct {
	Id      int64  `gorm:"primaryKey;autoIncrement;unique;column:id"`
	Name    string `gorm:"column:name"`
	Phone   string `gorm:"column:phone"`
	Email   string `gorm:"column:email"`
	Address string `gorm:"column:address"`
}
