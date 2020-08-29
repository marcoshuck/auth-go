package model

import (
	"github.com/jinzhu/gorm"
	"github.com/marcoshuck/auth-go/pkg/dto"
)

type User struct {
	gorm.Model
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
}

func (u User) ToPublic() dto.PublicUser {
	return dto.PublicUser{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Username:  u.Username,
	}
}
