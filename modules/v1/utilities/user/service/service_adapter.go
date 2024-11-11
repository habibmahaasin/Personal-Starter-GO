package service

import (
	"Batumbuah/modules/v1/utilities/user/models"
	"Batumbuah/modules/v1/utilities/user/repository"
)

type Service interface {
	Login(input models.LoginInput) (models.User, error)
	Register(fullName, email, password, address string, roleID int) error
	CheckIn(userID string, image, note string) error
	GetCheckInLogs(userID string) ([]models.CheckInLog, error)
	GetUserStats(userID string) (models.UserStats, error)
	UpdatePreTestStatus(userID string, email string, status bool) error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}
