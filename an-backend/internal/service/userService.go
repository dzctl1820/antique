package service

import (
	"an-backend/models"
	"an-backend/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

type UserDB struct {
	DB *gorm.DB
}

func (s *UserDB) create(user *models.User) error {
	return s.DB.Create(user).Error
}

func (s *UserDB) getById(id uint) (*models.User, error) {
	var user models.User
	err := s.DB.First(&user, id).Error
	return &user, err
}

func (s *UserDB) getByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (s *UserDB) LoginByEmail(req *models.LoginRequestByEmail) (*models.LoginResponse, error) {
	user, err := s.getByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if req.Password != user.Password {
		return nil, errors.New("password is incorrect")
	}
	token, err := utils.GenerateJwtToken(user.ID, user.Email, time.Hour*24)
	if err != nil {
		return nil, err
	}
	return &models.LoginResponse{Token: token}, nil
}

func (s *UserDB) RegisterByEmail(req *models.RegisterRequestByEmail) error {
	user := &models.User{
		Name:     "hero" + fmt.Sprintf(strconv.Itoa(rand.Intn(90000)+10000)),
		Email:    req.Email,
		Password: req.Password,
		Phone:    "12345678901",
	}
	return s.create(user)
}
