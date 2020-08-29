package service

import (
	"github.com/marcoshuck/auth-go/pkg/dto"
	"github.com/marcoshuck/auth-go/pkg/model"
)

type UserService interface {
	Login(input dto.Login) (string, error)
	Register(input dto.Register) (*model.User, error)
}
