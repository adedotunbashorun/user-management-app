package routes

import (
	"user-management-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
}
