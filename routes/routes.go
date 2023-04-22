package routes

import (
	"go-final-project-mygram/handlers"
	"go-final-project-mygram/middlewares"
	"go-final-project-mygram/services"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title MyGram API
// @version 1.0
// @description This is a service for mygram app API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE=2.0.html
// @host localhost:8080
// @BasePath /
func RegisterAPI(r *gin.Engine, app services.ServicesInterface) {
	server := handlers.NewHttpServer(app)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", server.UserRegister)
		userRouter.POST("/login", server.UserLogin)
	}

	photosRouter := r.Group("/photos")
	{
		photosRouter.Use(middlewares.Authentication())
		photosRouter.POST("/", server.CreatePhotos)
		photosRouter.GET("/", server.GetAllPhotos)
		photosRouter.GET("/:photoId", server.GetPhotoById)
		photosRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), server.UpdatePhotoById)
		photosRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), server.DeletePhotoById)
	}

	socialMediasRouter := r.Group("/socialmedias")
	{
		socialMediasRouter.Use(middlewares.Authentication())
		socialMediasRouter.POST("/", server.CreateSocialMedia)
		socialMediasRouter.GET("/", server.GetAllSocialMedias)
		socialMediasRouter.GET("/:socialmediaId", server.GetSocialMediaById)
		socialMediasRouter.PUT("/:socialmediaId", middlewares.SocialMediaAuthorization(), server.UpdateSocialMediaById)
		socialMediasRouter.DELETE("/:socialmediaId", middlewares.SocialMediaAuthorization(), server.DeleteSocialMediaById)
	}

	commentsRouter := r.Group("/comments")
	{
		commentsRouter.Use(middlewares.Authentication())
		commentsRouter.POST("/", server.CreateComment)
		commentsRouter.GET("/", server.GetAllComments)
		commentsRouter.GET("/:commentId", server.GetCommentById)
		commentsRouter.PUT("/:commentId", middlewares.CommentAuthorization(), server.UpdateCommentById)
		commentsRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), server.DeleteCommentById)
	}
}
