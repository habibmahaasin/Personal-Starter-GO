package repository

import (
	"Go_Starter/modules/v1/utilities/user/models"

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
        return nil
    })
}