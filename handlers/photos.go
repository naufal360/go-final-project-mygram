package handlers

import (
	"fmt"
	"go-final-project-mygram/helpers"
	"go-final-project-mygram/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h *HttpServer) CreatePhotos(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)
	Photo := models.Photos{}
	userId := userData["id"].(string)

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserId = userId

	fmt.Println(userData)

	data, err := h.service.CreatePhoto(Photo)
	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success create photo"
	helpers.SuccessCreated(c, msg, data)
}

func (h *HttpServer) GetAllPhotos(c *gin.Context) {

	data, err := h.service.GetAllPhotos()

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success get all photos"
	helpers.SuccessOk(c, msg, data)
}

func (h *HttpServer) GetPhotoById(c *gin.Context) {
	photoId := c.Param("photoId")

	data, err := h.service.GetPhotoById(photoId)

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success get photo by id"
	helpers.SuccessOk(c, msg, data)
}

func (h *HttpServer) UpdatePhotoById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photos{}
	photoId := c.Param("photoId")
	userId := userData["id"].(string)

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserId = userId

	Photo.Id = photoId

	data, err := h.service.UpdatePhotoById(photoId, Photo)

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success update photo"
	helpers.SuccessUpdated(c, msg, data)
}

func (h *HttpServer) DeletePhotoById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	Photo := models.Photos{}
	photoId := c.Param("photoId")
	userId := userData["id"].(string)

	Photo.UserId = userId

	Photo.Id = photoId

	err := h.service.DeletePhotoById(photoId)

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success delete photo"
	helpers.SuccessDeleted(c, msg)
}
