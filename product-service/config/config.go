package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DatabaseConfig() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading Configuration")
	}
	fmt.Println("Config Done")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// dbConnect := "root:oOxPeoXhtQxVUWIrYMkUnXHIjdbXtQlc@tcp(mysql.railway.internal:3306)/railway"
	// return dbConnect
	// return "mysql://root:oOxPeoXhtQxVUWIrYMkUnXHIjdbXtQlc@viaduct.proxy.rlwy.net:36162/railway"
}
