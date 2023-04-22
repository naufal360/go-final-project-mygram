package services

import "go-final-project-mygram/models"

type SocialMediasInterface interface {
	CreateSocialMedia(req models.SocialMedias) (res models.SocialMedias, err error)
	GetAllSocialMedias() (res []models.SocialMedias, err error)
	GetSocialMediaById(id string) (res models.SocialMedias, err error)
	UpdateSocialMediaById(id string, req models.SocialMedias) (res models.SocialMedias, err error)
	DeleteSocialMediaById(id string) (err error)
}

func (s *Services) CreateSocialMedia(req models.SocialMedias) (res models.SocialMedias, err error) {
	res, err = s.repo.CreateSocialMedia(req)
	if err != nil {
		return res, err
	}
	return
}

func (s *Services) GetAllSocialMedias() (res []models.SocialMedias, err error) {
	res, err = s.repo.GetAllSocialMedias()
	if err != nil {
		return res, err
	}
	return
}

func (s *Services) GetSocialMediaById(id string) (res models.SocialMedias, err error) {
	res, err = s.repo.GetSocialMediaById(id)
	if err != nil {
		return res, err
	}
	return
}

func (s *Services) UpdateSocialMediaById(id string, req models.SocialMedias) (res models.SocialMedias, err error) {
	res, err = s.repo.UpdateSocialMediaById(id, req)
	if err != nil {
		return res, err
	}
	return
}

func (s *Services) DeleteSocialMediaById(id string) (err error) {
	err = s.repo.DeleteSocialMediaById(id)
	if err != nil {
		return err
	}
	return
}
