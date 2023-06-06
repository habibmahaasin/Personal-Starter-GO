package service

import (
	"GuppyTech/modules/v1/utilities/user/models"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (n *service) Login(input models.LoginInput) (models.User, error) {
	email := input.Email
	password := input.Password
	user, _ := n.repository.GetUserByEmail(email)

	if user.User_id == "" {
		return user, errors.New("Email yang Anda masukan salah")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println(err)
		return user, errors.New("Password yang Anda masukan salah/tidak terdaftar")
	}
	return user, nil
}
