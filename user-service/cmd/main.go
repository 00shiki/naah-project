package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	pb "user-service/proto"
	"user-service/repository"
	"user-service/service"
)

func main() {
	_ = godotenv.Load()

	address := fmt.Sprintf(":%v", os.Getenv("PORT"))
	lis, err := net.Listen("tcp", address)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	repoUser := repository.NewUserRepository(db)
	s := grpc.NewServer()
	reflection.Register(s)
	us := service.NewUserService(repoUser)
	pb.RegisterUserServiceServer(s, us)
	fmt.Printf("Starting gRPC server on port %v\n", address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
