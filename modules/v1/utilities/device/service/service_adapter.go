package service

import (
	"GuppyTech/modules/v1/utilities/device/models"
	"GuppyTech/modules/v1/utilities/device/repository"
	myJSON "GuppyTech/pkg/json"
)

type Service interface {
	GetDatafromWebhook(sensorData string, antaresDeviceID string) (models.ConnectionDat, error)
	GetAllDevices() ([]models.Device, error)
	GetDeviceHistory() ([]models.DeviceHistory, string, error)
	Control(id string, power string, mode string) error
	PostControlAntares(antares_id string, token string, power string, mode string) error
	AddDevice(input models.DeviceInput, user_id string) error
	GetDeviceById(u_id string, d_id string) (models.Device, error)
	GetDeviceHistoryById(d_id string) ([]models.DeviceHistory, string, error)
}

type service struct {
	json       myJSON.JSON
	repository repository.Repository
}

func NewService(repository repository.Repository, myJSON myJSON.JSON) *service {
	return &service{repository: repository, json: myJSON}
}
