package main

import (
	"api-gateway/api"
	pb "api-gateway/proto"
	"api-gateway/service/carts"
	"api-gateway/service/deliveries"
	"api-gateway/service/orders"
	"api-gateway/service/users"
	"api-gateway/service/vouchers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {
	_ = godotenv.Load()

	userConn, err := grpc.NewClient(os.Getenv("USER_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed connect to user service: %v", err)
	}

	orderConn, err := grpc.NewClient(os.Getenv("ORDER_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed connect to order service: %v", err)
	}

	userClient := pb.NewUserServiceClient(userConn)
	userService := users.NewUserService(userClient)

	cartClient := pb.NewCartServiceClient(orderConn)
	cartService := carts.NewCartService(cartClient)

	orderClient := pb.NewOrderServiceClient(orderConn)
	orderService := orders.NewOrderService(orderClient)

	deliveryClient := pb.NewDeliveryServiceClient(orderConn)
	deliveryService := deliveries.NewDeliveryService(deliveryClient)

	voucherClient := pb.NewVoucherServiceClient(orderConn)
	voucherService := vouchers.NewVoucherService(voucherClient)

	e := echo.New()

	api.Init(e, userService, cartService, orderService, deliveryService, voucherService)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
