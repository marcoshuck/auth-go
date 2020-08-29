package dto

// CreateUser has the needed information to create a new user.
type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// PublicUser has the user's information that can be publicly shared.
type PublicUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}
