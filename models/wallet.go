// models/user.go
package models

type Wallet struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	UserId  uint    `gorm:"not null;unique_index"`
	Balance float64 `gorm:"not null"`
}
