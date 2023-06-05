package handler

import (
	"GuppyTech/modules/v1/utilities/device/repository"
	"GuppyTech/modules/v1/utilities/device/service"

	"gorm.io/gorm"
)

type deviceHandler struct {
	deviceService service.Service
}

func NewDeviceHandler(productService service.Service) *deviceHandler {
	return &deviceHandler{productService}
}

func Handler(db *gorm.DB) *deviceHandler {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	Handler := NewDeviceHandler(Service)
	return Handler
}
