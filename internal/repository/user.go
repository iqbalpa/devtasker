package repository

import (
	"devtasker/internal/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(name, username, hashedPass string) (model.User, error)
	GetUserByUsername(username string) (model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(name, username, hashedPass string) (model.User, error) {
	u := model.User{
		Name:         name,
		Username:     username,
		PasswordHash: hashedPass,
	}
	ur.db.Create(&u)
	return u, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	result := ur.db.First(&user, "username = ?", username)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}
