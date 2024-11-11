package service

import (
	"Batumbuah/modules/v1/utilities/user/models"
	"errors"
	"fmt"
	"time"

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

	existingUser, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return err
	}

	if existingUser.UserID != "" {
		return errors.New("email already exists")
	}

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

func (s *service) CheckIn(userID, image, note string) error {
	lastCheckIn, err := s.repository.GetLastCheckInTime(userID)
	if err != nil {
		if err.Error() == "no check-in found for the user" {
			return s.repository.UserCheckIn(userID, image, note)
		}
		return err
	}

	if time.Since(lastCheckIn.DateCreated) < 7*24*time.Hour {
		return errors.New("check-in allowed only once every 7 days")
	}

	return s.repository.UserCheckIn(userID, image, note)
}

func (s *service) GetCheckInLogs(userID string) ([]models.CheckInLog, error) {
	return s.repository.GetCheckInLogs(userID)
}

func (s *service) GetUserStats(userID string) (models.UserStats, error) {
	return s.repository.GetUserStats(userID)
}

func (s *service) UpdatePreTestStatus(userID string, email string, status bool) error {
	err := s.repository.UpdateTestInformation(models.TestInformation{
		UserID: userID,
		Email:  email,
	})
	
	if err != nil {
		return fmt.Errorf("failed to create test information: %w", err)
	}

	err = s.repository.UpdateUserStats(userID, models.UserStats{
		UserID:      userID,
		IsPreTested: status,
		DateUpdated: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to update user stats: %w", err)
	}

	return nil
}
