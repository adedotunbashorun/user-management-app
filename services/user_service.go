package services

import (
	"errors"
	"user-management-app/models"
	"user-management-app/repositories"
	"user-management-app/utils"

	"go.mongodb.org/mongo-driver/bson"
)

const userNotFoundErr = "user not found"

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
		return models.User{}, errors.New(userNotFoundErr)
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return models.User{}, errors.New("invalid password")
	}
	return user, nil
}

func (us *UserService) GetUserFromDB(id string) (models.UserResponse, error) {
	user, err := us.UserRepo.GetUserByID(id)
	if err != nil {
		return models.UserResponse{}, errors.New(userNotFoundErr)
	}

	userResponse := models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
	}

	return userResponse, nil
}

func (us *UserService) UpdateUserInDB(id string, updateData bson.M) (models.User, error) {
	return us.UserRepo.UpdateUserByID(id, updateData)
}
