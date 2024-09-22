package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
	"time"
	JWT_ENTITY "user-service/entity/jwt"
	"user-service/entity/users"
	pb "user-service/proto"
	"user-service/repository"
)

type Service struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
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
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		log.Printf("CreateUser: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot create users")
	}

	// TODO: send email verification

	return &pb.CreateUserResponse{UserId: user.ID}, nil
}

func (s *Service) ValidateUser(ctx context.Context, req *pb.ValidateUserRequest) (*pb.ValidateUserResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "users not found")
		}
		log.Printf("ValidateUser: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot get user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid password")
	}

	claims := JWT_ENTITY.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * 24 * time.Second)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create token")
	}

	return &pb.ValidateUserResponse{JwtToken: token}, nil
}

func (s *Service) GetUserDetail(ctx context.Context, req *pb.GetUserDetailRequest) (*pb.GetUserDetailResponse, error) {
	user, err := s.repo.GetUserByID(req.GetUserId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "users not found")
		}
		log.Printf("GetUserDetail: %v", err)
		return nil, status.Errorf(codes.Internal, "cannot get user")
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

func (s *Service) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*emptypb.Empty, error) {
	// TODO: Verify email
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
