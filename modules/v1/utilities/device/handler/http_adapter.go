package handler

import (
	"GuppyTech/app/config"
	"GuppyTech/modules/v1/utilities/device/repository"
	"GuppyTech/modules/v1/utilities/device/service"
	myJSON "GuppyTech/pkg/json"

	"gorm.io/gorm"
)

type deviceHandler struct {
	deviceService service.Service
}

func NewDeviceHandler(productService service.Service) *deviceHandler {
	return &deviceHandler{productService}
}

func Handler(db *gorm.DB, conf config.Conf) *deviceHandler {
	Repository := repository.NewRepository(db, conf)
	Service := service.NewService(Repository, myJSON.Instance())
	Handler := NewDeviceHandler(Service)
	return Handler
}
