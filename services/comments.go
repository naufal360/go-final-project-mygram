package services

import "go-final-project-mygram/models"

type CommentsInterface interface {
	CreateComment(req models.Comments) (res models.Comments, err error)
	GetAllComments() (res []models.Comments, err error)
	GetCommentById(id string) (res models.Comments, err error)
	UpdateCommentById(id string, req models.Comments) (res models.Comments, err error)
	DeleteCommentById(id string) (err error)
}

func (s *Services) CreateComment(req models.Comments) (res models.Comments, err error) {
	res, err = s.repo.CreateComment(req)
	if err != nil {
		return res, err
	}
	return
}

func (s *Services) GetAllComments() (res []models.Comments, err error) {
	res, err = s.repo.GetAllComments()
	if err != nil {
		return res, err
	}
	return
}

func (s *Services) GetCommentById(id string) (res models.Comments, err error) {
	res, err = s.repo.GetCommentById(id)
	if err != nil {
		return res, err
	}
	return
}

func (s *Services) UpdateCommentById(id string, req models.Comments) (res models.Comments, err error) {
	res, err = s.repo.UpdateCommentById(id, req)
	if err != nil {
		return res, err
	}
	return
}

func (s *Services) DeleteCommentById(id string) (err error) {
	err = s.repo.DeleteCommentById(id)
	if err != nil {
		return err
	}
	return
}
