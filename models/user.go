// models/user.go
package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique_index"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
}
