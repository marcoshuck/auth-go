package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

// Profile represents an user profile.
type Profile struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Claims extends the default claims from the JWT package.
type Claims struct {
	Profile Profile
	jwt.StandardClaims
}
