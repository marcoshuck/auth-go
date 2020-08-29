package dto

// Register has the needed information to register an user.
type Register struct {
	// FirstName is the user's first name.
	FirstName string `json:"first_name"`
	// LastName is the user's last name.
	LastName string `json:"last_name"`
	// Email is the user's email.
	Email string `json:"email"`
	// Username is the unique user's identifier.
	Username string `json:"username"`
	// Password is the plain-text user's password.
	Password string `json:"password"`
}
