package handlers

import (
	"go-final-project-mygram/helpers"
	"go-final-project-mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

// RegisterUsers godoc
// @Summary Post a new users data
// @Description Post details of users corresponding to the input Id
// @Tags users
// @Accept json
// @Produce json
// @Param models.User body models.User true "register users"
// @Success 201 {object} models.User
// @Router /users/register [post]
func (h *HttpServer) UserRegister(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.Users{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	data, err := h.service.RegisterUsers(User)
	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       data.Id,
		"email":    data.Email,
		"username": data.Username,
	})
}

// LoginUsers godoc
// @Summary Post a login users data
// @Description Post details of users corresponding to the input Id
// @Tags users
// @Accept json
// @Produce json
// @Param models.User body models.User true "login users"
// @Success 200 {object} object{token=string}
// @Router /users/login [post]
func (h *HttpServer) UserLogin(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.Users{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	token, err := h.service.LoginUsers(User)

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
