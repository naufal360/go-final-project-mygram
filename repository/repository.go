package repository

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	UserInterface
	PhotosInterface
	SocialMediasInterface
	CommentsInterface
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
