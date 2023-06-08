package repository

import (
	"GuppyTech/modules/v1/utilities/user/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserByEmail(email string) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
