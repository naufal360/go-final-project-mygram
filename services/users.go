package services

import (
	"errors"
	"go-final-project-mygram/helpers"
	"go-final-project-mygram/models"
)

type UserInterface interface {
	RegisterUsers(req models.Users) (res models.Users, err error)
	LoginUsers(req models.Users) (token string, err error)
}

func (s *Services) RegisterUsers(req models.Users) (res models.Users, err error) {
	res.Age = uint(req.Age)
	res, err = s.repo.RegisterUsers(req)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Services) LoginUsers(req models.Users) (token string, err error) {
	password := req.Password
	data, err := s.repo.LoginUsers(req)
	if err != nil {
		return "", err
	}

	comparePassword := helpers.ComparePass(data.Password, password)
	if !comparePassword {
		return "", errors.New("invalid email/password")
	}

	token = helpers.GenerateToken(data.Id, data.Email)

	return token, nil
}
