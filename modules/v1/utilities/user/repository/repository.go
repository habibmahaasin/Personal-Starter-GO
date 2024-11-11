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

        TestInformation := models.TestInformation{
            UserID: user.UserID,
            Email: user.Email,
        }
        if err := tx.Create(&TestInformation).Error; err != nil {
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

        var checkInCount int64
        if err := tx.Model(&models.CheckInLog{}).Where("user_id = ?", userID).Count(&checkInCount).Error; err != nil {
            return err
        }

        if err := tx.Model(&models.UserStats{}).Where("user_id = ?", userID).Update("total_check_in", checkInCount).Error; err != nil {
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

func (r *repository) GetUserStats(userID string) (models.UserStats, error) {
    var userStats models.UserStats
    err := r.db.Where("user_id = ?", userID).First(&userStats).Error
    return userStats, err
}

func (r *repository) UpdateUserStats(userID string, userStats models.UserStats) error {
    return r.db.Model(&models.UserStats{}).Where("user_id = ?", userID).Updates(userStats).Error
}

func (r *repository) UpdateTestInformation(testInformation models.TestInformation) error {
    return r.db.Model(&models.TestInformation{}).Where("user_id = ?", testInformation.UserID).Updates(testInformation).Error
}