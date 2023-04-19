package main

import (
	"fmt"
	"net/http"
	"project-mygram/app"
	"project-mygram/configs"
	"project-mygram/controller"
	"project-mygram/database"
	"project-mygram/helpers"
	routers "project-mygram/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/basic/docs"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDbConnection()

	repo     = app.WiringRepository(db)
	services = app.WiringService(repo)

	userController        controller.UserController        = controller.NewUserController(services)
	socialMediaController controller.SocialMediaController = controller.NewSocialMediaController(services)
	commentController     controller.CommentController     = controller.NewCommentController(services)
	photoController       controller.PhotoController       = controller.NewPhotoController(services)
)

func main() {
	defer database.CloseDbConnection(db)
	config := configs.GetInstance()
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%v", config.Appconfig.Port)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.APIResponse("welcome its server", http.StatusOK, "false"))
	})

	// Route here
	routers.UserRouter(router, userController)
	routers.SocialMediaRouter(router, socialMediaController)
	routers.CommentRouter(router, commentController)
	routers.PhotoRouter(router, photoController)

	listen := fmt.Sprintf(":%v", config.Appconfig.Port)
	router.Run(listen)
}
