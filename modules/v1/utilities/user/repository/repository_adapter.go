package repository

import (
	"Batumbuah/modules/v1/utilities/user/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserByEmail(email string) (models.User, error)
	CreateUser(user *models.User) error
	UserCheckIn(userID, image, note string) error
	GetLastCheckInTime(userID string) (models.CheckInLog, error)
	GetCheckInLogs(userID string) ([]models.CheckInLog, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
