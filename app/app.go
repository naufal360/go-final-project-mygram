package app

import (
	"fmt"
	"go-final-project-mygram/config"
	"go-final-project-mygram/repository"
	"go-final-project-mygram/routes"
	"go-final-project-mygram/services"

	"os"

	_ "go-final-project-mygram/docs"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApp() {
	repo := repository.NewRepository(config.GORM.DB)
	app := services.NewServices(repo)
	routes.RegisterAPI(router, app)

	port := os.Getenv("APP_PORT")
	router.Use(gin.Recovery())
	router.Run(fmt.Sprintf(":%s", port))
}
