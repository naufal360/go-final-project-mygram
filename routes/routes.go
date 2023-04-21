package routes

import (
	"go-final-project-mygram/handlers"
	"go-final-project-mygram/middlewares"
	"go-final-project-mygram/services"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

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
}
