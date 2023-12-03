package db

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dsn := "host=localhost port=5432 user=postgres dbname=synergize_tech password=16051994 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to the database.")
}
