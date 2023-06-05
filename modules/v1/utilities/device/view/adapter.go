package view

import (
	"GuppyTech/modules/v1/utilities/device/repository"
	"GuppyTech/modules/v1/utilities/device/service"

	"gorm.io/gorm"
)

type deviceView struct {
	deviceService service.Service
}

func NewDeviceView(deviceService service.Service) *deviceView {
	return &deviceView{deviceService}
}

func View(db *gorm.DB) *deviceView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	View := NewDeviceView(Service)
	return View
}
