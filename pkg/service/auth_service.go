package service

import (
	"errors"
	"github.com/marcoshuck/auth-go/pkg/dto"
	"github.com/marcoshuck/auth-go/pkg/jwt"
	"github.com/marcoshuck/auth-go/pkg/model"
	"github.com/marcoshuck/auth-go/pkg/password"
	"github.com/marcoshuck/auth-go/pkg/validator"
)

// AuthService groups a set of methods to login and register an user into the system.
type AuthService interface {
	// Login checks if the given login credentials are valid and returns a token and the expiration time.
	Login(input dto.Login) (string, int64, error)
	// Register checks if the given sign up credentials are valid and returns the user that was created.
	Register(input dto.Register) (*model.User, error)
}

// authService is a AuthService implementation.
type authService struct {
	// userService is the reference to a UserService used to get and create users.
	userService UserService
	// secretKey is the key used to sign tokens.
	secretKey []byte
	validator validator.Validator
}

// Login checks if the given login credentials are valid and returns a token and the expiration time.
func (a *authService) Login(input dto.Login) (string, int64, error) {
	// Validate input
	err := a.validator.Struct(input)
	if err != nil {
		return "", 0, err
	}

	// Get the user by email.
	user, err := a.userService.GetByEmail(input.Email)
	if err != nil {
		return "", 0, err
	}

	// Compare passwords
	ok := password.Compare(input.Password, user.Password)
	if !ok {
		return "", 0, errors.New("invalid credentials")
	}

	// Sign the token
	token, expiresAt, err := jwt.Encode(jwt.JWT{
		UUID:      user.UUID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, a.secretKey)

	if err != nil {
		return "", 0, err
	}

	return *token, *expiresAt, nil
}

// Register checks if the given sign up credentials are valid and returns the user that was created.
func (a *authService) Register(input dto.Register) (*model.User, error) {
	// Validate input
	err := a.validator.Struct(input)
	if err != nil {
		return nil, err
	}

	// Generate password
	pass, err := password.Generate(input.Password)
	if err != nil {
		return nil, err
	}

	// Create user dto
	user := dto.CreateUser{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Username:  input.Username,
		Password:  *pass,
	}

	// Delegate user creation to user service.
	result, err := a.userService.Create(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// NewAuthService initializes a new authentication service using an UserService and a secretKey.
func NewAuthService(userService UserService, validator validator.Validator, secretKey []byte) AuthService {
	return &authService{
		userService: userService,
		validator:   validator,
		secretKey:   secretKey,
	}
}
