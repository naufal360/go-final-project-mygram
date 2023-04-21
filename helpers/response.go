package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessOk(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": msg,
		"data":    data,
	})
}

func SuccessCreated(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"message": msg,
		"data":    data,
	})
}

func SuccessUpdated(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"message": msg,
		"data":    data,
	})
}

func SuccessDeleted(c *gin.Context, msg string) {
	c.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"message": msg,
	})
}

func CustomBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   true,
		"message": "Bad Request, " + err.Error(),
	})
}

func AbortAuth(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error":   true,
		"message": "Unauthorized, " + err.Error(),
	})
}

func AbortBadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error":   true,
		"message": "Bad Request, invalid parameter",
	})
}

func AbortNotFound(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"error":   true,
		"message": "Data Not Found, data doesn't exist",
	})
}

func AbortUnzuthorized(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error":   "Unauthorized",
		"message": "you are not allowed to access this data",
	})
}
