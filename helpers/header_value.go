package helpers

import "github.com/gin-gonic/gin"

func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}

var (
	appJson = "application/json"
)

func BindingRequest(c *gin.Context, contentType string, data any) {
	if contentType == appJson {
		c.ShouldBindJSON(&data)
	} else {
		c.ShouldBind(&data)
	}
}
