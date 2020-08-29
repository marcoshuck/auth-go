package model

import (
	"github.com/jinzhu/gorm"
	"github.com/marcoshuck/auth-go/pkg/dto"
)

// User represents an user in the system.
type User struct {
	gorm.Model
	// UUID is used to identify users across the system.
	UUID string `json:"uuid"`
	// FirstName is the user's first name.
	FirstName string `json:"first_name"`
	// LastName is the user's last name.
	LastName string `json:"last_name"`
	// Email is the user's email.
	Email string `json:"email" gorm:"unique"`
	// Username is the user's identifier.
	Username string `json:"username" gorm:"unique"`
	// Password is the encrypted user's password.
	Password string `json:"password"`
}

// ToPublic extracts the information that can be publicly shared.
func (u User) ToPublic() dto.PublicUser {
	return dto.PublicUser{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
	}
}
