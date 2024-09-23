package users

import "time"

type User struct {
	ID        int64
	Email     string
	Password  string
	FirstName string
	LastName  string
	BirthDate time.Time
	Address   string
	ContactNo string
	Role
}

type Role int

const (
	Customer Role = iota + 1
	Admin
)
