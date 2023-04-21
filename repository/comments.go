package repository

import "go-final-project-mygram/models"

type CommentsInterface interface {
	CreateComment(req models.Comments) (res models.Comments, err error)
	GetAllComments() (res []models.Comments, err error)
	GetCommentById(id string) (res models.Comments, err error)
	UpdateCommentById(id string, req models.Comments) (res models.Comments, err error)
	DeleteCommentById(id string) (err error)
}

func (r Repository) CreateComment(req models.Comments) (res models.Comments, err error) {
	err = r.db.Debug().Create(&req).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return
}

func (r Repository) GetAllComments() (res []models.Comments, err error) {
	err = r.db.Debug().Model(&res).Find(&res).Error
	if err != nil {
		return res, err
	}
	return
}

func (r Repository) GetCommentById(id string) (res models.Comments, err error) {
	err = r.db.Debug().Model(&res).First(&res, "id = ?", id).Error
	if err != nil {
		return res, err
	}
	return
}

func (r Repository) UpdateCommentById(id string, req models.Comments) (res models.Comments, err error) {
	err = r.db.Debug().Model(&res).Where("id = ?", id).Updates(&req).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return
}

func (r Repository) DeleteCommentById(id string) (err error) {
	var comment models.Comments
	err = r.db.Debug().Model(&comment).Delete(&comment, "id = ?", id).Error
	if err != nil {
		return err
	}
	return
}
