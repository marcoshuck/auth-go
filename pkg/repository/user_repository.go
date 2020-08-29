package repository

import (
	"github.com/marcoshuck/auth-go/pkg/dto"
	"github.com/marcoshuck/auth-go/pkg/model"
)

type UserRepository interface {
	Get(uuid string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(register dto.Register) (*model.User, error)
}
