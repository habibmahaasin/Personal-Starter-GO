package view

import (
	"Go_Starter/modules/v1/utilities/user/repository"
	"Go_Starter/modules/v1/utilities/user/service"

	"gorm.io/gorm"
)

type userView struct {
	userService service.Service
}

func NewUserView(userService service.Service) *userView {
	return &userView{userService}
}

func View(db *gorm.DB) *userView {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	View := NewUserView(Service)
	return View
}
