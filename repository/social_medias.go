package repository

import "go-final-project-mygram/models"

type SocialMediasInterface interface {
	CreateSocialMedia(req models.SocialMedias) (res models.SocialMedias, err error)
	GetAllSocialMedias() (res []models.SocialMedias, err error)
	GetSocialMediaById(id string) (res models.SocialMedias, err error)
	UpdateSocialMediaById(id string, req models.SocialMedias) (res models.SocialMedias, err error)
	DeleteSocialMediaById(id string) (err error)
}

func (r Repository) CreateSocialMedia(req models.SocialMedias) (res models.SocialMedias, err error) {
	err = r.db.Debug().Create(&req).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return
}

func (r Repository) GetAllSocialMedias() (res []models.SocialMedias, err error) {
	err = r.db.Debug().Model(&res).Find(&res).Error
	if err != nil {
		return res, err
	}
	return
}

func (r Repository) GetSocialMediaById(id string) (res models.SocialMedias, err error) {
	err = r.db.Debug().Model(&res).First(&res, "id = ?", id).Error
	if err != nil {
		return res, err
	}
	return
}

func (r Repository) UpdateSocialMediaById(id string, req models.SocialMedias) (res models.SocialMedias, err error) {
	err = r.db.Debug().Model(&res).Where("id = ?", id).Updates(&req).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return
}

func (r Repository) DeleteSocialMediaById(id string) (err error) {
	var socialMedia models.SocialMedias
	err = r.db.Model(&socialMedia).Delete(&socialMedia, "id = ?", id).Error
	if err != nil {
		return err
	}
	return
}
