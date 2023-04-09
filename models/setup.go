package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Faield  to connect db")
	}
	fmt.Printf("init connection succefully")
	DB = database
}

func DBMigrate() {
	DB.AutoMigrate(&Coin{})

}
