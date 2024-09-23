package users

import (
	"api-gateway/entity/users"
	pb "api-gateway/proto"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type UserService struct {
	client pb.UserServiceClient
}

func NewUserService(client pb.UserServiceClient) *UserService {
	return &UserService{
		client: client,
	}
}

func (us *UserService) RegisterUser(user *users.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.CreateUserRequest{
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		BirthDate: timestamppb.New(user.BirthDate),
		Address:   user.Address,
		ContactNo: user.ContactNo,
		Role:      int64(user.Role),
	}
	var res, err = us.client.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	user.ID = res.GetUserId()
	return nil
}

func (us *UserService) LoginUser(user *users.User) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.ValidateUserRequest{
		Email:    user.Email,
		Password: user.Password,
	}
	res, err := us.client.ValidateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &res.JwtToken, err
}

func (us *UserService) GetUserDetail(userID int64) (*users.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.GetUserDetailRequest{
		UserId: userID,
	}
	res, err := us.client.GetUserDetail(ctx, req)
	if err != nil {
		return nil, err
	}
	user := &users.User{
		ID:        userID,
		Email:     res.GetEmail(),
		FirstName: res.GetFirstName(),
		LastName:  res.GetLastName(),
		BirthDate: res.GetBirthDate().AsTime(),
		Address:   res.GetAddress(),
		ContactNo: res.GetContactNo(),
	}
	return user, nil
}

func (us *UserService) VerifyEmail(userID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.VerifyEmailRequest{
		UserId: userID,
	}
	_, err := us.client.VerifyEmail(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
