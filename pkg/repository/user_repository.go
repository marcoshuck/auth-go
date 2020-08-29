package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/marcoshuck/auth-go/pkg/model"
)

type UserRepository interface {
	Get(uuid string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(user model.User) (*model.User, error)
}

type gormUserRepository struct {
	db *gorm.DB
}

func (g *gormUserRepository) Get(uuid string) (*model.User, error) {
	var user model.User
	err := g.db.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *gormUserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := g.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *gormUserRepository) Create(user model.User) (*model.User, error) {
	err := g.db.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{}
}
