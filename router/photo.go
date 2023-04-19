package router

import (
	"project-mygram/controller"
	"project-mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func PhotoRouter(e *gin.Engine, h controller.PhotoController) {
	photoRoutes := e.Group("/photos")
	{
		photoRoutes.Use(middlewares.Authentication())
		photoRoutes.POST("/", middlewares.CheckAuthorization(), h.CreatePhoto)
		photoRoutes.GET("/", middlewares.CheckAuthorization(), h.GetAll)
		photoRoutes.PUT("/:id", middlewares.PhotoAuthorization(), h.UpdatePhoto)
		photoRoutes.GET("/:id", middlewares.CheckAuthorization(), h.GetOne)
		photoRoutes.DELETE("/:id", middlewares.PhotoAuthorization(), h.DeletePhoto)
	}
}
