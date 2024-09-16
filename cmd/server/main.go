package main

import (
	"user-management-app/config"
	"user-management-app/controllers"
	"user-management-app/repositories"
	"user-management-app/routes"
	"user-management-app/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	client := config.ConnectDB()
	defer client.Disconnect(nil)

	userRepo := repositories.UserRepository{Collection: config.GetCollection(client, "users")}
	userService := services.UserService{UserRepo: &userRepo}
	userController := controllers.UserController{UserService: &userService}

	router := gin.Default()
	routes.SetupRoutes(router, &userController)

	router.Run(":8080")
}
