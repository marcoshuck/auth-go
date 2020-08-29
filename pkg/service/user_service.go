package service

import (
	"github.com/google/uuid"
	"github.com/marcoshuck/auth-go/pkg/dto"
	"github.com/marcoshuck/auth-go/pkg/model"
	"github.com/marcoshuck/auth-go/pkg/repository"
)

// UserService is the service used to get and create users.
type UserService interface {
	// Get gets an user with the given UUID.
	Get(uuid string) (*model.User, error)

	// GetByEmail gets an user with the given email.
	GetByEmail(email string) (*model.User, error)

	// Create creates a new user from the given create user dto.
	Create(createUser dto.CreateUser) (*model.User, error)
}

// userService is a UserService implementation.
type userService struct {
	repository repository.UserRepository
}

// Get gets an user with the given UUID.
func (u userService) Get(uuid string) (*model.User, error) {
	user, err := u.repository.Get(uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetByEmail gets an user with the given email.
func (u userService) GetByEmail(email string) (*model.User, error) {
	user, err := u.repository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Create creates a new user from the given create user dto.
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

// NewUserService initializes a new UserService implementation with the given user repository.
func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}
