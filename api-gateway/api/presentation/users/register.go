package users

type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
	Address   string `json:"address" validate:"required"`
	ContactNo string `json:"contact_no" validate:"required"`
	Admin     bool   `json:"admin"`
}

type RegisterResponse struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
	Address   string `json:"address" validate:"required"`
	ContactNo string `json:"contact_no" validate:"required"`
}
