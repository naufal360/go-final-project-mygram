package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SocialMedias represents the model for socialmedias
type SocialMedias struct {
	Id             string     `gorm:"primaryKey" json:"id"`
	Name           string     `gorm:"not null" json:"name" form:"name" valid:"required~Your name is required"`
	SocialMediaUrl string     `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Your social media url is required"`
	UserId         string     `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	User           *Users     `json:"user"`
}

func (s *SocialMedias) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	s.Id = "s-" + uuid.New().String()
	err = nil
	return
}
