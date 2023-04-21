package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   "Bad Request",
		"message": err.Error(),
	})
}

func AbortAuth(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error":   "Unauthorized",
		"message": err.Error(),
	})
}

func AbortBadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error":   "Bad Request",
		"message": "invalid parameter",
	})
}

func AbortNotFound(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"error":   "Data Not Found",
		"message": "data doesn't exist",
	})
}

func AbortUnzuthorized(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error":   "Unauthorized",
		"message": "you are not allowed to access this data",
	})
}
