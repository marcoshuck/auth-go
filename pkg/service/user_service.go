package service

import (
	"github.com/google/uuid"
	"github.com/marcoshuck/auth-go/pkg/dto"
	"github.com/marcoshuck/auth-go/pkg/model"
	"github.com/marcoshuck/auth-go/pkg/repository"
)

type UserService interface {
	Get(uuid string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(createUser dto.CreateUser) (*model.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func (u userService) Get(uuid string) (*model.User, error) {
	user, err := u.repository.Get(uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userService) GetByEmail(email string) (*model.User, error) {
	user, err := u.repository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userService) Create(createUser dto.CreateUser) (*model.User, error) {
	user, err := u.repository.Create(model.User{
		UUID:      uuid.New().String(),
		FirstName: createUser.FirstName,
		LastName:  createUser.LastName,
		Email:     createUser.Email,
		Username:  createUser.Username,
		Password:  createUser.Password,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}
