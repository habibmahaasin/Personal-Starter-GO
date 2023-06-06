package repository

import (
	"GuppyTech/modules/v1/utilities/device/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetLatestCon(token string) (models.Received, error)
	BindSensorData(Device_id string, input models.ConnectionDat) error
	GetAllDevices() ([]models.Device, error)
	GetDeviceByAntares(antaresDeviceID string) (models.Device, error)
	GetDeviceHistory() ([]models.DeviceHistory, error)
	Control(id string, power string, mode string) error
	PostControlAntares(antares_id string, token string, power string, mode string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
