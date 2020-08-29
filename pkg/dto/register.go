package dto

// Register has the needed information to register an user.
type Register struct {
	// FirstName is the user's first name.
	FirstName string `json:"first_name" validate:"required"`

	// LastName is the user's last name.
	LastName string `json:"last_name" validate:"required"`

	// Email is the user's email.
	Email string `json:"email"  validate:"required,email"`

	// Username is the unique user's identifier.
	Username string `json:"username" validate:"required,alphanum"`

	// Password is the plain-text user's password.
	Password string `json:"password" validate:"required"`
}
