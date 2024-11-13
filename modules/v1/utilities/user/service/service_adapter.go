package service

import (
	"Batumbuah/modules/v1/utilities/user/models"
	"Batumbuah/modules/v1/utilities/user/repository"
)

type Service interface {
	Login(input models.LoginInput) (models.User, error)
	Register(fullName, email, password, address string, roleID int) error
	RegisterPlant(userID, name, email string) error
	GetPlantByUserID(userID string) ([]models.UserPlant, error)
	GetPlantByID(plantID string) (models.UserPlant, error)
	CheckIn(userID, plantID, image, note string) error 
	GetCheckInLogs(UserPlantID string) ([]models.CheckInLog, error)
	GetPlantStatsById(plantID string) (models.PlantStats, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}
