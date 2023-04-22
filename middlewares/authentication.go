package middlewares

import (
	"go-final-project-mygram/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			helpers.AbortAuth(c, err)
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
