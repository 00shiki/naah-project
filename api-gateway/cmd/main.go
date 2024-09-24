package main

import (
	"api-gateway/api"
	pb "api-gateway/proto"
	"api-gateway/service/users"
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
	userClient := pb.NewUserServiceClient(userConn)
	userService := users.NewUserService(userClient)

	e := echo.New()

	api.Init(e, userService)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
