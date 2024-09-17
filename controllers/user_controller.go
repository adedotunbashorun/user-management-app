package controllers

import (
	"net/http"
	"user-management-app/models"
	"user-management-app/services"
	"user-management-app/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type UserController struct {
	UserService *services.UserService
}

func (uc *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password encryption failed"})
		return
	}
	user.Password = hash

	err = uc.UserService.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (uc *UserController) Login(c *gin.Context) {
	var loginData models.User
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := uc.UserService.Login(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) GetUser(c *gin.Context) {
	userId := c.Param("id")

	user, err := uc.UserService.GetUserFromDB(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found " + userId})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	userId := c.Param("id")

	var updatedUserData struct {
		Username string `json:"username"`
		Name     string `json:"name"`
	}

	if err := c.ShouldBindJSON(&updatedUserData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user := bson.M{
		"username": updatedUserData.Username,
		"name":     updatedUserData.Name,
	}

	if _, err := uc.UserService.UpdateUserInDB(userId, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
