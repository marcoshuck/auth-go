package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/marcoshuck/auth-go/pkg/model"
)

var (
	// ErrInvalidID is returned when an user id is passed on the create method.
	ErrInvalidID = errors.New("error invalid id")
)

// UserRepository groups a set of methods to manage user models.
type UserRepository interface {
	// Get returns an user from the given UUID.
	Get(uuid string) (*model.User, error)
	// GetByEmail returns an user from the given email.
	GetByEmail(email string) (*model.User, error)
	// Create user creates a new user.
	Create(user model.User) (*model.User, error)
}

// gormUserRepository is a UserRepository implementation using gorm.
type gormUserRepository struct {
	// db has a reference to a gorm database.
	db *gorm.DB
}

// Get returns an user from the given UUID.
func (g *gormUserRepository) Get(uuid string) (*model.User, error) {
	var user model.User
	err := g.db.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail returns an user from the given email.
func (g *gormUserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := g.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create user creates a new user.
func (g *gormUserRepository) Create(user model.User) (*model.User, error) {
	// Check that the user is a new user.
	if user.Model.ID != 0 {
		return nil, ErrInvalidID
	}

	// Save the new user.
	err := g.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// NewUserRepository initializes a new gorm UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{
		db: db,
	}
}
