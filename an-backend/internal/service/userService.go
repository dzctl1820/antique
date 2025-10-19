package service

import (
	"an-backend/models"
	"an-backend/utils"
	"errors"
	"gorm.io/gorm"
)

type DataDB struct {
	DB *gorm.DB
}

func (s *DataDB) create(user *models.User) error {
	return s.DB.Create(user).Error
}

func (s *DataDB) getById(id uint) (*models.User, error) {
	var user models.User
	err := s.DB.First(&user, id).Error
	return &user, err
}

func (s *DataDB) getByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (s *DataDB) LoginByEmail(req *models.LoginRequestByEmail) (*models.LoginResponse, error) {
	user, err := s.getByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if req.Password != user.Password {
		return nil, errors.New("password is incorrect")
	}
	token, err := utils.GenerateJwtToken(user.Email)
	if err != nil {
		return nil, err
	}
	return &models.LoginResponse{Token: token}, nil
}
