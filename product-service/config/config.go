package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() *sql.DB {
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
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Execute SQL to create tables
	if err := createTables(); err != nil {
		log.Fatal("Failed to create tables:", err)
	}

	log.Println("Successfully connected to the database and created tables")
	return DB
}

func createTables() error {
	tableQueries := []string{
		`CREATE TABLE IF NOT EXISTS vouchers (
            voucher_id VARCHAR(255) PRIMARY KEY,
            discount DECIMAL(10, 2) NOT NULL,
            valid_until DATE,
            used BOOLEAN NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        );`,

		`CREATE TABLE IF NOT EXISTS users (
            user_id INT AUTO_INCREMENT PRIMARY KEY,
            email VARCHAR(255) UNIQUE NOT NULL,
            password_hash VARCHAR(255) NOT NULL,
            first_name VARCHAR(255) NOT NULL,
            last_name VARCHAR(255) NOT NULL,
            birth_date DATE,
            address TEXT,
            contact_no VARCHAR(20),
            role INT NOT NULL,
            verified BOOLEAN DEFAULT FALSE,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        );`,

		`CREATE TABLE IF NOT EXISTS shoe_models (
            model_id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            price INT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        );`,

		`CREATE TABLE IF NOT EXISTS shoe_details (
            shoe_id INT AUTO_INCREMENT PRIMARY KEY,
            model_id INT NOT NULL,
            size INT NOT NULL,
            stock INT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            FOREIGN KEY (model_id) REFERENCES shoe_models (model_id)
        );`,

		`CREATE TABLE IF NOT EXISTS carts (
            cart_id INT AUTO_INCREMENT PRIMARY KEY,
            user_id INT NOT NULL,
            quantity INT NOT NULL,
            shoe_id INT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users (user_id),
            FOREIGN KEY (shoe_id) REFERENCES shoe_details (shoe_id)
        );`,

		`CREATE TABLE IF NOT EXISTS orders (
            order_id INT AUTO_INCREMENT PRIMARY KEY,
            user_id INT,
            voucher_id VARCHAR(255),
            status VARCHAR(20),
            price INT,
            fee INT,
            discount INT,
            total_price INT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            metadata TEXT,
            FOREIGN KEY (user_id) REFERENCES users (user_id),
            FOREIGN KEY (voucher_id) REFERENCES vouchers (voucher_id)
        );`,

		`CREATE TABLE IF NOT EXISTS payments (
            payment_id INT AUTO_INCREMENT PRIMARY KEY,
            order_id INT,
            payment_external_id VARCHAR(36),
            amount INT,
            status VARCHAR(255),
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            metadata TEXT,
            FOREIGN KEY (order_id) REFERENCES orders (order_id)
        );`,

		`CREATE TABLE IF NOT EXISTS deliveries (
            delivery_id INT AUTO_INCREMENT PRIMARY KEY,
            order_id INT,
            track_id VARCHAR(225),
            delivery_date DATETIME,
            arrival_date DATETIME,
            courier_name VARCHAR(50),
            courier_service VARCHAR(100),
            weight_grams INT,
            origin_city_id VARCHAR(50),
            destination_city_id VARCHAR(50),
            delivery_fee INT,
            status VARCHAR(20),
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            metadata TEXT,
            FOREIGN KEY (order_id) REFERENCES orders (order_id)
        );`,

		`CREATE TABLE IF NOT EXISTS order_details (
            order_detail_id INT AUTO_INCREMENT PRIMARY KEY,
            order_id INT NOT NULL,
            shoe_id INT NOT NULL,
            quantity INT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            FOREIGN KEY (order_id) REFERENCES orders (order_id),
            FOREIGN KEY (shoe_id) REFERENCES shoe_details (shoe_id)
        );`,
	}

	for _, query := range tableQueries {
		if _, err := DB.Exec(query); err != nil {
			return err
		}
	}

	return nil
}
