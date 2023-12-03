// models/user.go
package models

type Account struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	UserId  uint   `gorm:"not null"`
	Account string `gorm:"not null"`
	Bank    string `gorm:"not null"`
	Name    string `gorm:"not null"`
}
