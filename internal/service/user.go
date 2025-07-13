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
	hashPass, err := utils.HashPassword(pass)
	if err != nil {
		return model.User{}, err
	}
	fmt.Println(name, username, hashPass)
	u, err := us.r.CreateUser(name, username, hashPass)
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
	token, err := utils.GenerateToken(u)
	if err != nil {
		return "", err
	}
	return token, nil
}
