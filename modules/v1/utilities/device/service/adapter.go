package service

import (
	"GuppyTech/modules/v1/utilities/device/models"
	"GuppyTech/modules/v1/utilities/device/repository"
)

type Service interface {
	GetLatestCon(token string) (models.Received, error)
	GetDatafromWebhook(sensorData string, antaresDeviceID string) (models.ConnectionDat, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}
