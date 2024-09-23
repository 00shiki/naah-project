package users

import "api-gateway/entity/users"

type Service interface {
	RegisterUser(user *users.User) error
	LoginUser(user *users.User) (*string, error)
	GetUserDetail(userID int64) (*users.User, error)
	VerifyEmail(userID int64) error
}
