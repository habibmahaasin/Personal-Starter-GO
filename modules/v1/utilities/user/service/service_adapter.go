package service

import (
	"Batumbuah/modules/v1/utilities/user/models"
	"Batumbuah/modules/v1/utilities/user/repository"
)

type Service interface {
	Login(input models.LoginInput) (models.User, error)
	Register(fullName, email, password, address string, roleID int) error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}
