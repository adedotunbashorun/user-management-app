package services

import (
	"errors"
	"user-management-app/models"
	"user-management-app/repositories"
	"user-management-app/utils"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func (us *UserService) RegisterUser(user models.User) error {
	_, err := us.UserRepo.GetUserByEmail(user.Email)
	if err == nil {
		return errors.New("user already exists")
	}
	return us.UserRepo.CreateUser(user)
}

func (us *UserService) Login(email, password string) (models.User, error) {
	user, err := us.UserRepo.GetUserByEmail(email)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return models.User{}, errors.New("invalid password")
	}
	return user, nil
}
