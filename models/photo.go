package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photos struct {
	Id        string `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"not null" json:"title" form:"title" valid:"required~Your title is required"`
	Caption   string `gorm:"type:varchar(255)" json:"caption" form:"caption" valid:"optional,type(string)~Your caption type must be string"`
	PhotoUrl  string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Your photo url is required"`
	UserId    string
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	User      *Users
	Comment   []Comments `gorm:"foreignKey:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment"`
}

func (p *Photos) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	p.Id = "p-" + uuid.New().String()
	err = nil
	return
}
