package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=your_username dbname=your_database_name password=your_password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate will create the 'users' table based on the User model
	db.AutoMigrate(&User{})

	fmt.Println("Connected to the database.")
}
