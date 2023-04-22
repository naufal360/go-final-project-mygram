package handlers

import (
	"fmt"
	"go-final-project-mygram/helpers"
	"go-final-project-mygram/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateSocialMedias godoc
// @Summary Post a new socialmedias data
// @Description Post details of socialmedias corresponding to the input Id
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param models.SocialMedias body models.SocialMedias true "create social medias"
// @Success 201 {object} object{data=models.SocialMedias, error=boolean, message=string}
// @Router /socialmedias [post]
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

// GetAllSocialMedias godoc
// @Summary Get all socialmedias data
// @Description Get all socialmedias
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Success 200 {object} object{data=models.SocialMedias, error=boolean, message=string}
// @Router /socialmedias [get]
func (h *HttpServer) GetAllSocialMedias(c *gin.Context) {

	data, err := h.service.GetAllSocialMedias()

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success get all social medias"
	helpers.SuccessOk(c, msg, data)
}

// GetSocialMediaById godoc
// @Summary Get details of socialmedias corresponding to the input Id
// @Description Get all socialmedias
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param Id path string true "ID of the social medias"
// @Success 200 {object} object{data=models.SocialMedias, error=boolean, message=string}
// @Router /socialmedias/{socialmediaId} [get]
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

// UpdateSocialMediaById godoc
// @Summary Update a socialmedia data
// @Description Update details of socialmedias corresponding to the input Id
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param Id path string true "ID of the social medias"
// @Success 200 {object} object{data=models.SocialMedias, error=boolean, message=string}
// @Router /socialmedias/{socialmediaId} [put]
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

// Delete SocialMediaById godoc
// @Summary Delete details of socialmedias corresponding to the input Id
// @Description Delete a socialmedia
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param Id path string true "ID of the social medias"
// @Success 200 {object} object{error=boolean, message=string}
// @Router /socialmedias/{socialmediaId} [delete]
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
