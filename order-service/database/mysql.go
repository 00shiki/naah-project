// database.go
package database

import (
	"database/sql"
	"fmt"
	"log"
	"order-service/config"

	_ "github.com/go-sql-driver/mysql"
)

// DB is a global variable to hold the database connection pool
var db *sql.DB

// InitDB initializes the MySQL connection
func InitDB() error {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)

	// Open the MySQL connection
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("FAILED INITATE CONNECTION TO MYSQL DATABASE: %w", err)
	}

	// Ping the database to ensure the connection is established
	if err = db.Ping(); err != nil {
		return fmt.Errorf("FAILED INITATE CONNECTION TO MYSQL DATABASE: %w", err)
	}

	log.Println("SUCCESS INITATE CONNECTION TO MYSQL DATABASE")
	return nil
}

// GetDB returns the database connection pool
func GetDB() *sql.DB {
	return db
}
