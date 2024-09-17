package routes

import (
	"user-management-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
	// Create a route group with the base path /api/v1
	api := router.Group("/api/v1")
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)
		// api.GET("/users", userController.GetUsers)
	}
}
