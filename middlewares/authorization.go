package middlewares

import (
	"go-final-project-mygram/config"
	"go-final-project-mygram/helpers"
	"go-final-project-mygram/models"
	"go-final-project-mygram/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var repo = repository.NewRepository(config.GORM.DB)
		photoId := c.Param("photoId")
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := userData["id"].(string)
		Photo := models.Photos{}

		Photo, err := repo.GetPhotoById(photoId)
		if err != nil {
			helpers.AbortNotFound(c, err)
			return
		}

		if Photo.UserId != userId {
			helpers.AbortUnzuthorized(c, err)
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var repo = repository.NewRepository(config.GORM.DB)
		socialMediaId := c.Param("socialmediaId")
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := userData["id"].(string)
		SocialMedia := models.SocialMedias{}

		SocialMedia, err := repo.GetSocialMediaById(socialMediaId)
		if err != nil {
			helpers.AbortNotFound(c, err)
			return
		}

		if SocialMedia.UserId != userId {
			helpers.AbortUnzuthorized(c, err)
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var repo = repository.NewRepository(config.GORM.DB)
		commentId := c.Param("commentId")
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := userData["id"].(string)
		Comment := models.Comments{}

		Comment, err := repo.GetCommentById(commentId)
		if err != nil {
			helpers.AbortNotFound(c, err)
			return
		}

		if Comment.UserId != userId {
			helpers.AbortUnzuthorized(c, err)
			return
		}

		c.Next()
	}
}
