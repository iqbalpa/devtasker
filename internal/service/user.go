package service

import (
	"devtasker/internal/model"
	"devtasker/internal/repository"
	"devtasker/internal/utils"
	"fmt"
)

type IUserService interface {
	Register(name, username, pass string) (model.User, error)
	Login(username, pass string) (string, error)
}

type UserService struct {
	r repository.IUserRepository
}

func NewUserService(r repository.IUserRepository) *UserService {
	return &UserService{
		r: r,
	}
}

func (us *UserService) Register(name, username, pass string) (model.User, error) {
	fmt.Println(">>>>>>> service auth register")
	hashPass, err := utils.HashPassword(pass)
	fmt.Println(">>>>>>> service auth register level 2")
	if err != nil {
		return model.User{}, err
	}
	fmt.Println(name, username, hashPass)
	u, err := us.r.CreateUser(name, username, hashPass)
	fmt.Println(">>>>>>> service auth register level 3")
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (us *UserService) Login(username, pass string) (string, error) {
	u, err := us.r.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	res := utils.ComparePassword(u.PasswordHash, pass)
	if !res {
		return "", fmt.Errorf("username or password is wrong")
	}
	return "<token>", nil
}
