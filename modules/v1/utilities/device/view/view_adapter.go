package view

import (
	"GuppyTech/app/config"
	"GuppyTech/modules/v1/utilities/device/repository"
	"GuppyTech/modules/v1/utilities/device/service"
	myJSON "GuppyTech/pkg/json"

	"gorm.io/gorm"
)

type deviceView struct {
	deviceService service.Service
}

func NewDeviceView(deviceService service.Service) *deviceView {
	return &deviceView{deviceService}
}

func View(db *gorm.DB, conf config.Conf) *deviceView {
	Repository := repository.NewRepository(db, conf)
	Service := service.NewService(Repository, myJSON.Instance())
	View := NewDeviceView(Service)
	return View
}
