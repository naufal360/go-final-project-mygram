package repository

import "go-final-project-mygram/models"

type UserInterface interface {
	RegisterUsers(req models.Users) (res models.Users, err error)
	LoginUsers(req models.Users) (res models.Users, err error)
}

func (r Repository) RegisterUsers(req models.Users) (res models.Users, err error) {
	err = r.db.Debug().Create(&req).Take(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repository) LoginUsers(req models.Users) (res models.Users, err error) {
	err = r.db.Debug().Where("email = ?", req.Email).Take(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
