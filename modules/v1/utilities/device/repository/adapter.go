package repository

import (
	"GuppyTech/modules/v1/utilities/device/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetLatestCon(token string) (models.Received, error)
	BindSensorData(DeviceId string, input models.ConnectionDat) error
	GetAllDevices() ([]models.Device, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
