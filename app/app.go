package app

import (
	"fmt"
	"go-final-project-mygram/config"
	"go-final-project-mygram/repository"
	"go-final-project-mygram/routes"
	"go-final-project-mygram/services"

	"os"

	"github.com/gin-gonic/gin"
	// _ "go-rest-api-mock/docs"
)

var router = gin.New()

// @title Product with User API
// @version 1.0
// @description This is an api service for exercising
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE=2.0.html
// @host localhost:8080
// @BasePath /
func StartApp() {
	repo := repository.NewRepository(config.GORM.DB)
	app := services.NewServices(repo)
	routes.RegisterAPI(router, app)

	port := os.Getenv("APP_PORT")
	router.Use(gin.Recovery())
	router.Run(fmt.Sprintf(":%s", port))
}
