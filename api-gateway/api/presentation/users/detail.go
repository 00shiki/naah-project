package users

type DetailResponse struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
	Address   string `json:"address"`
	ContactNo string `json:"contact_no"`
}
