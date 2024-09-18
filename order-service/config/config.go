package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB_HOST string
var DB_PORT string
var DB_NAME string
var DB_USER string
var DB_PASSWORD string

var PORT string

func InitConfig() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("FAILED TO INITIATE CONFIG")
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	PORT = os.Getenv("PORT")

	log.Println("SUCCESS INITIATE CONFIG")

}
