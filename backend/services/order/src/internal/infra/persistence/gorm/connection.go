package persistence

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while trying to connect to DB: ", err.Error())
	}

	if err := db.AutoMigrate(&Order{}); err != nil {
		log.Fatal("Error while trying to migrate Order: ", err.Error())
	}

	if err := db.AutoMigrate(&OrderItem{}); err != nil {
		log.Fatal("Error while trying to migrate Order Item: ", err.Error())
	}

	return db
}
