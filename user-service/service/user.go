package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
	"time"
	"user-service/entity/email"
	JWT_ENTITY "user-service/entity/jwt"
	"user-service/entity/users"
	"user-service/pkg/rabbitmq"
	pb "user-service/proto"
	"user-service/repository"
)

type Service struct {
	repo     repository.Repository
	rabbitMQ *rabbitmq.RabbitMQClient
}

func NewUserService(repo repository.Repository, rabbitMQ *rabbitmq.RabbitMQClient) *Service {
	return &Service{
		repo:     repo,
		rabbitMQ: rabbitMQ,
	}
}

func (s *Service) CreateUser(_ context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot hash password")
	}

	user := &users.User{
		Email:     req.GetEmail(),
		Password:  string(passwordHash),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		BirthDate: req.GetBirthDate().AsTime(),
		Address:   req.GetAddress(),
		ContactNo: req.GetContactNo(),
		Role:      int(req.GetRole()),
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		log.Printf("CreateUser: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot create users")
	}

	data := email.Email{
		To:              user.Email,
		Subject:         "Email Verification",
		Body:            "",
		Type:            "verification",
		VerificationURL: fmt.Sprintf("%s/verify?user_id=%d", os.Getenv("USER_ROUTE"), user.ID),
	}
	err = s.rabbitMQ.Push("email_queue", data)
	if err != nil {
		log.Printf("CreateUser: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot push email")
	}

	return &pb.CreateUserResponse{UserId: user.ID}, nil
}

func (s *Service) ValidateUser(_ context.Context, req *pb.ValidateUserRequest) (*pb.ValidateUserResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "users not found")
		}
		log.Printf("ValidateUser: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot get users")
	}

	if !user.Verified {
		return nil, status.Errorf(codes.Unauthenticated, "user is not verified")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid password")
	}

	claims := JWT_ENTITY.Claims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * 24 * time.Hour)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Printf("ValidateUser: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot create token")
	}

	return &pb.ValidateUserResponse{JwtToken: token}, nil
}

func (s *Service) GetUserDetail(_ context.Context, req *pb.GetUserDetailRequest) (*pb.GetUserDetailResponse, error) {
	user, err := s.repo.GetUserByID(req.GetUserId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "users not found")
		}
		log.Printf("GetUserDetail: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot get users")
	}

	return &pb.GetUserDetailResponse{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		BirthDate: timestamppb.New(user.BirthDate),
		Address:   user.Address,
		ContactNo: user.ContactNo,
	}, nil
}

func (s *Service) VerifyEmail(_ context.Context, req *pb.VerifyEmailRequest) (*emptypb.Empty, error) {
	user, err := s.repo.GetUserByID(req.GetUserId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "users not found")
		}
		log.Printf("VerifyEmail: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot get users")
	}
	user.Verified = true
	err = s.repo.UpdateUser(user)
	if err != nil {
		log.Printf("UpdateUser: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot update user")
	}
	return nil, nil
}
