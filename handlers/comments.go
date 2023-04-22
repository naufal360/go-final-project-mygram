package handlers

import (
	"fmt"
	"go-final-project-mygram/helpers"
	"go-final-project-mygram/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateComents godoc
// @Summary Post a new comments data
// @Description Post details of comments corresponding to the input Id
// @Tags comments
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param models.Comments body models.Comments true "create comments"
// @Success 201 {object} object{data=models.Comments, error=boolean, message=string}
// @Router /comments [post]
func (h *HttpServer) CreateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)
	Comment := models.Comments{}
	userId := userData["id"].(string)

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserId = userId

	fmt.Println(userData)

	data, err := h.service.CreateComment(Comment)
	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success create comment"
	helpers.SuccessCreated(c, msg, data)
}

// GetAllComments godoc
// @Summary Get all comments data
// @Description Get all comments
// @Tags comments
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Success 200 {object} object{data=models.Comments, error=boolean, message=string}
// @Router /comments [get]
func (h *HttpServer) GetAllComments(c *gin.Context) {

	data, err := h.service.GetAllComments()

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success get all comments"
	helpers.SuccessOk(c, msg, data)
}

// GetCommentById godoc
// @Summary Get details of comments corresponding to the input Id
// @Description Get all comments
// @Tags comments
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param Id path string true "ID of the comments"
// @Success 200 {object} object{data=models.Comments, error=boolean, message=string}
// @Router /comments/{commentId} [get]
func (h *HttpServer) GetCommentById(c *gin.Context) {
	commentId := c.Param("commentId")

	data, err := h.service.GetCommentById(commentId)

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success get comment by id"
	helpers.SuccessOk(c, msg, data)
}

// UpdateCommentById godoc
// @Summary Update a comment data
// @Description Update details of comments corresponding to the input Id
// @Tags comments
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param Id path string true "ID of the comments"
// @Success 200 {object} object{data=models.Comments, error=boolean, message=string}
// @Router /comments/{commentId} [put]
func (h *HttpServer) UpdateCommentById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comments{}
	commentId := c.Param("commentId")
	userId := userData["id"].(string)

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserId = userId

	Comment.Id = commentId

	data, err := h.service.UpdateCommentById(commentId, Comment)

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success update comment"
	helpers.SuccessUpdated(c, msg, data)
}

// Delete CommentById godoc
// @Summary Delete details of comments corresponding to the input Id
// @Description Delete a comment
// @Tags comments
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param Id path string true "ID of the comments"
// @Success 200 {object} object{error=boolean, message=string}
// @Router /comments/{commentId} [delete]
func (h *HttpServer) DeleteCommentById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	Comment := models.Comments{}
	commentId := c.Param("commentId")
	userId := userData["id"].(string)

	Comment.UserId = userId

	Comment.Id = commentId

	err := h.service.DeleteCommentById(commentId)

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	msg := "Success delete comment"
	helpers.SuccessDeleted(c, msg)
}
