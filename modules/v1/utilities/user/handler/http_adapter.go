package handler

import (
	"Go_Starter/modules/v1/utilities/user/repository"
	"Go_Starter/modules/v1/utilities/user/service"
	jwt "Go_Starter/pkg/jwt"

	"gorm.io/gorm"
)

type userHandler struct {
	userService service.Service
	jwtoken     jwt.JwToken
}

func NewUserHandler(userService service.Service, jwtoken jwt.JwToken) *userHandler {
	return &userHandler{userService, jwtoken}
}

func Handler(db *gorm.DB) *userHandler {
	userRepository := repository.NewRepository(db)
	userService := service.NewService(userRepository)
	userHandler := NewUserHandler(userService, jwt.NewJwToken())

	return userHandler
}
