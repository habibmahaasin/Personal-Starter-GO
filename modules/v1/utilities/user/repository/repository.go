package repository

import (
	"Batumbuah/modules/v1/utilities/user/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

func (r *repository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	return user, err
}

func (r *repository) CreateUser(user *models.User) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(user).Error; err != nil {
            return err
        }

        userStats := models.UserStats{
            UserID:           user.UserID,
            TotalCheckIn:     0,
            IsPreTested:      false,
            IsPostTested:     false,
            IsRedeemedReward: false,
			IsAvailableToRedeem: false,
            DateCreated:      time.Now(),
            DateUpdated:      time.Now(),
        }
        if err := tx.Create(&userStats).Error; err != nil {
            return err
        }

        return nil
    })
}


func (r *repository) UserCheckIn(userID, image, note string) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        checkInLog := models.CheckInLog{
            UserID:      userID,
            Image:       image,
            Note:        note,
            DateCreated: time.Now(),
            DateUpdated: time.Now(),
        }
        if err := tx.Create(&checkInLog).Error; err != nil {
            return err
        }

        var userStats models.UserStats
        if err := tx.Where("user_id = ?", userID).First(&userStats).Error; err != nil {
            return err
        }

        userStats.TotalCheckIn++
        if err := tx.Save(&userStats).Error; err != nil {
            return err
        }

        return nil
    })
}

func (r *repository) GetLastCheckInTime(userID string) (models.CheckInLog, error) {
	var checkInLog models.CheckInLog
	err := r.db.Where("user_id = ?", userID).Order("date_created desc").First(&checkInLog).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return checkInLog, errors.New("no check-in found for the user")

		}
		return checkInLog, err
	}
	return checkInLog, nil
}

func (r *repository) GetCheckInLogs(userID string) ([]models.CheckInLog, error) {
    var checkInLogs []models.CheckInLog
    err := r.db.Where("user_id = ?", userID).Find(&checkInLogs).Error
    return checkInLogs, err
}