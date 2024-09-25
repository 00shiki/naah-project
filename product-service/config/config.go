package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"product-service/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	log.Println("Successfully connected to the database")

	if err := DB.AutoMigrate(
		&models.User{},
		&models.ShoeModel{},
		&models.ShoeDetail{},
		&models.Cart{},
		&models.Order{},
		&models.OrderDetail{},
		&models.Payment{},
		&models.Delivery{},
		&models.Voucher{},
	); err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Println("Table already exists, skipping migration for that table.")
		} else {
			log.Fatalf("Failed to auto-migrate models: %v", err)
		}
	}

	log.Println("Successfully auto-migrated models")
	return DB
}
