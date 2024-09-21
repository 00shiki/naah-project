package main

import (
	"fmt"
	"log"
	"net"
	"order-service/config"
	"order-service/database"
	"order-service/handler"
	"order-service/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	config.InitConfig()
	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	db := database.GetDB()
	defer db.Close()

	cartHandler := handler.NewCartHandler(db)
	deliveryHandler := handler.NewDeliveryHandler(db)

	// s := grpc.NewServer()
	s := grpc.NewServer()

	pb.RegisterCartServiceServer(s, cartHandler)
	pb.RegisterDeliveryServiceServer(s, deliveryHandler)

	reflection.Register(s)

	port := config.PORT
	if port == "" {
		port = "50054"
	}
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("RUNNING ORDER SERVICE ON PORT: %s", port)

	if err = s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
