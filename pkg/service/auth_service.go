package service

import (
	"errors"
	"github.com/marcoshuck/auth-go/pkg/dto"
	"github.com/marcoshuck/auth-go/pkg/jwt"
	"github.com/marcoshuck/auth-go/pkg/model"
	"github.com/marcoshuck/auth-go/pkg/password"
)

type AuthService interface {
	Login(input dto.Login) (string, int64, error)
	Register(input dto.Register) (*model.User, error)
}

type authService struct {
	userService UserService
	secretKey   []byte
}

func (a *authService) Login(input dto.Login) (string, int64, error) {
	user, err := a.userService.GetByEmail(input.Email)
	if err != nil {
		return "", 0, err
	}

	ok := password.Compare(input.Password, user.Password)
	if !ok {
		return "", 0, errors.New("invalid credentials")
	}

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

func (a *authService) Register(input dto.Register) (*model.User, error) {
	pass, err := password.Generate(input.Password)
	if err != nil {
		return nil, err
	}

	user := dto.CreateUser{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Username:  input.Username,
		Password:  *pass,
	}

	return a.userService.Create(user)
}

func NewAuthService(userService UserService, secretKey []byte) AuthService {
	return &authService{
		userService: userService,
		secretKey:   secretKey,
	}
}
