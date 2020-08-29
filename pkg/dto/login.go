package dto

// Login has the needed information to login an user.
type Login struct {
	// Email is the user's email.
	Email string `json:"email"`
	// Password is the plain-text user's password.
	Password string `json:"password"`
}
