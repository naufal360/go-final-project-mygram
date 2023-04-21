package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comments struct {
	Id        string     `gorm:"primaryKey" json:"id"`
	UserId    string     `json:"user_id"`
	PhotoId   string     `json:"photo_id"`
	Message   string     `gorm:"not null" json:"message" form:"message" valid:"required~Your message is required"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	User      *Users     `json:"user"`
	Photo     *Photos    `json:"photo"`
}

func (c *Comments) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	c.Id = "c-" + uuid.New().String()
	err = nil
	return
}
