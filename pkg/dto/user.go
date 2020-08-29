package dto

// CreateUser has the needed information to create a new user.
type CreateUser struct {
	// FirstName is the user's first name.
	FirstName string `json:"first_name" validate:"required"`

	// LastName is the user's last name.
	LastName string `json:"last_name" validate:"required"`

	// Email is the user's email.
	Email string `json:"email"  validate:"required,email"`

	// Username is the unique user's identifier.
	Username string `json:"username" validate:"required,alphanum"`

	// Password is the plain-text user's password.
	Password string `json:"password" validate:"required,base64"`
}

// PublicUser has the user's information that can be publicly shared.
type PublicUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}
