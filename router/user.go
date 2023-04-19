package router

import (
	"project-mygram/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(e *gin.Engine, userController controller.UserController) {
	userRoutes := e.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
	}
}
