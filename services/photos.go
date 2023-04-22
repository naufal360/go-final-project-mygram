package services

import "go-final-project-mygram/models"

type PhotosInterface interface {
	CreatePhoto(req models.Photos) (res models.Photos, err error)
	GetAllPhotos() (res []models.Photos, err error)
	GetPhotoById(id string) (res models.Photos, err error)
	UpdatePhotoById(id string, req models.Photos) (res models.Photos, err error)
	DeletePhotoById(id string) (err error)
}

func (s *Services) CreatePhoto(req models.Photos) (res models.Photos, err error) {
	res, err = s.repo.CreatePhoto(req)
	if err != nil {
		return res, err
	}
	return
}

func (s *Services) GetAllPhotos() (res []models.Photos, err error) {
	res, err = s.repo.GetAllPhotos()
	if err != nil {
		return res, err
	}
	return
}

func (s *Services) GetPhotoById(id string) (res models.Photos, err error) {
	res, err = s.repo.GetPhotoById(id)
	if err != nil {
		return res, err
	}
	return
}

func (s *Services) UpdatePhotoById(id string, req models.Photos) (res models.Photos, err error) {
	res, err = s.repo.UpdatePhotoById(id, req)
	if err != nil {
		return res, err
	}
	return
}

func (s *Services) DeletePhotoById(id string) (err error) {
	err = s.repo.DeletePhotoById(id)
	if err != nil {
		return err
	}
	return
}
