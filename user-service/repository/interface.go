package repository

import (
	"database/sql"
	"user-service/entity/users"
)

type Repository interface {
	Reader
	Writer
}

type Reader interface {
	GetUserByID(userID int64) (*users.User, error)
	GetUserByEmail(email string) (*users.User, error)
}

type Writer interface {
	CreateUser(user *users.User) error
	UpdateUser(user *users.User) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
