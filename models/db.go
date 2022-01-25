package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	USER := os.Getenv("POSTGRES_USER")
	PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DBNAME := os.Getenv("POSTGRES_DB")
	PORT  := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=pg_container user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", 
	USER, PASSWORD, DBNAME, PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&UserBook{})

	return db
}