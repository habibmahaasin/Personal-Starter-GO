package repository

import "Batumbuah/modules/v1/utilities/user/models"

func (r *repository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	return user, err
}

func (r *repository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}