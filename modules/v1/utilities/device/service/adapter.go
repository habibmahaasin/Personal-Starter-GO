package service

import (
	"GuppyTech/modules/v1/utilities/device/models"
	"GuppyTech/modules/v1/utilities/device/repository"
)

type Service interface {
	GetLatestCon(token string) (models.Received, error)
	GetDatafromWebhook(sensorData string, antaresDeviceID string) (models.ConnectionDat, error)
	GetAllDevices() ([]models.Device, error)
	GetDeviceHistory() ([]models.DeviceHistory, error)
	Control(id string, power string, mode string) error
	PostControlAntares(antares_id string, token string, power string, mode string) error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}
