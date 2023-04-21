package handlers

import (
	"fmt"
	"go-final-project-mygram/helpers"
	"go-final-project-mygram/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h *HttpServer) CreateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedias{}
	userId := userData["id"].(string)

	if contentType == appJson {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserId = userId

	fmt.Println(userData)

	data, err := h.service.CreateSocialMedia(SocialMedia)
	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success create social media"
	helpers.SuccessCreated(c, msg, data)
}

func (h *HttpServer) GetAllSocialMedias(c *gin.Context) {

	data, err := h.service.GetAllSocialMedias()

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success get all social medias"
	helpers.SuccessOk(c, msg, data)
}

func (h *HttpServer) GetSocialMediaById(c *gin.Context) {
	SocialMediaId := c.Param("socialmediaId")

	data, err := h.service.GetSocialMediaById(SocialMediaId)

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success get social media by id"
	helpers.SuccessOk(c, msg, data)
}

func (h *HttpServer) UpdateSocialMediaById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedias{}
	socialMediaId := c.Param("socialmediaId")
	userId := userData["id"].(string)

	if contentType == appJson {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserId = userId

	SocialMedia.Id = socialMediaId

	data, err := h.service.UpdateSocialMediaById(socialMediaId, SocialMedia)

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success update social media"
	helpers.SuccessUpdated(c, msg, data)
}

func (h *HttpServer) DeleteSocialMediaById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	SocialMedia := models.SocialMedias{}
	socialMediaId := c.Param("socialmediaId")
	userId := userData["id"].(string)

	SocialMedia.UserId = userId

	SocialMedia.Id = socialMediaId

	err := h.service.DeleteSocialMediaById(socialMediaId)

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success delete social media"
	helpers.SuccessDeleted(c, msg)
}
