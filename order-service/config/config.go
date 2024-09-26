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
var DELIVERY_API_KEY string

var PORT string
var DELIVERY_COST_URL string
var DELIVERY_PROVINCE_URL string
var DELIVERY_CITY_URL string

var XENDIT_SECRET_KEY string
var XENDIT_INVOICE_URL string

var RMQ_URL string

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
	DELIVERY_COST_URL = os.Getenv("DELIVERY_COST_URL")
	DELIVERY_PROVINCE_URL = os.Getenv("DELIVERY_PROVINCE_URL")
	DELIVERY_CITY_URL = os.Getenv("DELIVERY_CITY_URL")
	DELIVERY_API_KEY = os.Getenv("DELIVERY_API_KEY")
	XENDIT_SECRET_KEY = os.Getenv("XENDIT_SECRET_KEY")
	XENDIT_INVOICE_URL = os.Getenv("XENDIT_INVOICE_URL")
	PORT = os.Getenv("PORT")
	RMQ_URL = os.Getenv("RMQ_URL")

	log.Println("SUCCESS INITIATE CONFIG")

}
