package db

import (
	"fmt"

	"iamajraj/sql-practice/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Database connection failed")
		return
	}

	DB.AutoMigrate(models.User{})

	fmt.Println("Database connection opened")
}
