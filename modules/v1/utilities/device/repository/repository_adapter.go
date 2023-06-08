package repository

import (
	"GuppyTech/app/config"
	"GuppyTech/modules/v1/utilities/device/models"

	"gorm.io/gorm"
)

type Repository interface {
	BindSensorData(Device_id string, input models.ConnectionDat) (error, error)
	GetAllDevices() ([]models.Device, error)
	GetDeviceByAntares(antaresDeviceID string) (models.Device, error)
	GetDeviceHistory() ([]models.DeviceHistory, error)
	Control(id string, power string, mode string) error
	PostControlAntares(antares_id string, token string, power string, mode string) error
}

type repository struct {
	db   *gorm.DB
	conf config.Conf
}

func NewRepository(db *gorm.DB, conf config.Conf) *repository {
	return &repository{db, conf}
}
