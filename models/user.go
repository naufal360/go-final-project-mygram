package models

import (
	"go-final-project-mygram/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Users struct {
	Id          string         `gorm:"primaryKey" json:"id"`
	Username    string         `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email       string         `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password    string         `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age         uint           `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required,range(9|120)~Your age minimal is 9 years old"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
	Photo       []Photos       `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo"`
	SocialMedia []SocialMedias `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media"`
	Comment     []Comments     `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Id = "u-" + uuid.New().String()
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
