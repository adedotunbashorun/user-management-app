package routes

import (
	"user-management-app/controllers"
	"user-management-app/middleware" // Add this line to import the middleware package

	"github.com/gin-gonic/gin"
)

func SetupUsersRoutes(router *gin.Engine, userController *controllers.UserController) {
	// Create a route group with the base path /api/v1
	api := router.Group("/api/v1")
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)
	}

	// Authenticated routes
	auth := api.Group("/user")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/:id", userController.GetUser)
		auth.PUT("/:id", userController.UpdateUser)
	}
}
