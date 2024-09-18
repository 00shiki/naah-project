package main

import (
	"log"
	"order-service/config"
	"order-service/database"
)

func main() {

	config.InitConfig()
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	db := database.GetDB()
	defer db.Close()
}
