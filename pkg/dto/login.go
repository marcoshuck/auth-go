package dto

// Login represents the needed information to login an user in the system.
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
