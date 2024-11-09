package service

import (
	"Batumbuah/modules/v1/utilities/user/models"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (n *service) Login(input models.LoginInput) (models.User, error) {
	email := input.Email
	password := input.Password
	user, _ := n.repository.GetUserByEmail(email)

	if user.UserID == "" {
		return user, errors.New("email yang Anda masukan salah")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println(err)
		return user, errors.New("password yang Anda masukan salah/tidak terdaftar")
	}
	return user, nil
}

func (s *service) Register(fullName, email, password, address string, roleID int) error {
	// Generate a unique user ID
	userID := uuid.New().String()

	// Hash the password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create the user model
	user := &models.User{
		UserID:   userID,
		FullName: fullName,
		Email:    email,
		Password: string(hashedPassword),
		Address:  address,
		RoleID:   int64(roleID),
	}

	// Save user to the repository
	return s.repository.CreateUser(user)
}
