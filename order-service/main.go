package main

import (
	"fmt"
	"log"
	"net"
	"order-service/config"
	"order-service/database"
	"order-service/handler"
	"order-service/pb"
	"order-service/utils"

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
	rmq, err := utils.NewRabbitMQClient(config.RMQ_URL)
	if err != nil {
		log.Fatal(err)
	}

	cartHandler := handler.NewCartHandler(db)
	deliveryHandler := handler.NewDeliveryHandler(db, rmq)
	orderHandler := handler.NewOrderHandler(db, rmq)
	voucherHandler := handler.NewVoucherHandler(db)

	s := grpc.NewServer()

	pb.RegisterCartServiceServer(s, cartHandler)
	pb.RegisterDeliveryServiceServer(s, deliveryHandler)
	pb.RegisterOrderServiceServer(s, orderHandler)
	pb.RegisterVoucherServiceServer(s, voucherHandler)

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

	// TODO - Buat scheduler untuk update status delivery

	if err = s.Serve(listen); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("EXITING . . .")
}
