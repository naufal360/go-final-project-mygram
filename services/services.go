package services

import "go-final-project-mygram/repository"

type Services struct {
	repo repository.RepositoryInterface
}

type ServicesInterface interface {
	UserInterface
	PhotosInterface
	SocialMediasInterface
	CommentsInterface
}

func NewServices(repo repository.RepositoryInterface) ServicesInterface {
	return &Services{repo: repo}
}
