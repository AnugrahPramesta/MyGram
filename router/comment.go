package router

import (
	"project-mygram/controller"
	"project-mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func CommentRouter(e *gin.Engine, h controller.CommentController) {
	commentRoutes := e.Group("/comments")
	{
		commentRoutes.Use(middlewares.Authentication())
		commentRoutes.POST("/", h.CreateComment)
		commentRoutes.GET("/", middlewares.CheckAuthorization(), h.GetAll)
		commentRoutes.PUT("/:id", middlewares.CommentAuthorization(), h.UpdateComment)
		commentRoutes.GET("/:id", middlewares.CheckAuthorization(), h.GetOne)
		commentRoutes.DELETE("/:id", middlewares.CommentAuthorization(), h.DeleteComment)
	}
}
