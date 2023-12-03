// models/user.go
package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null;unique_index"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
}
