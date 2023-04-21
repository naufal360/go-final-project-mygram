package repository

import "go-final-project-mygram/models"

type PhotosInterface interface {
	CreatePhoto(req models.Photos) (res models.Photos, err error)
	GetAllPhotos() (res []models.Photos, err error)
	GetPhotoById(id string) (res models.Photos, err error)
	UpdatePhotoById(id string, req models.Photos) (res models.Photos, err error)
	DeletePhotoById(id string) (err error)
}

func (r Repository) CreatePhoto(req models.Photos) (res models.Photos, err error) {
	err = r.db.Debug().Create(&req).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return
}

func (r Repository) GetAllPhotos() (res []models.Photos, err error) {
	err = r.db.Debug().Preload("Comment").Model(&res).Find(&res).Error
	if err != nil {
		return res, err
	}
	return
}

func (r Repository) GetPhotoById(id string) (res models.Photos, err error) {
	err = r.db.Debug().Preload("Comment").Model(&res).First(&res, "id = ?", id).Error
	if err != nil {
		return res, err
	}
	return
}

func (r Repository) UpdatePhotoById(id string, req models.Photos) (res models.Photos, err error) {
	err = r.db.Debug().Preload("Comment").Model(&res).Where("id = ?", id).Updates(&req).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return
}

func (r Repository) DeletePhotoById(id string) (err error) {
	var photo models.Photos
	err = r.db.Debug().Model(&photo).Delete(&photo, "id = ?", id).Error
	if err != nil {
		return err
	}
	return
}
