package handlers

import (
	"go-final-project-mygram/helpers"
	"go-final-project-mygram/models"

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
// @Param models.Users body models.Users true "register users"
// @Success 201 {object} object{data=object{id=string, email=string, username=string}, error=boolean, message=string}
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

	newData := gin.H{
		"id":       data.Id,
		"email":    data.Email,
		"username": data.Username,
	}
	msg := "Success create users"
	helpers.SuccessCreated(c, msg, newData)
}

// LoginUsers godoc
// @Summary Post a login users data
// @Description Post details of users corresponding to the input Id
// @Tags users
// @Accept json
// @Produce json
// @Param models.Users body models.Users true "login users"
// @Success 200 {object} object{data=object{token=string}, error=boolean, message=string}
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

	dataToken := gin.H{
		"token": token,
	}
	msg := "Success login user"
	helpers.SuccessOk(c, msg, dataToken)
}
