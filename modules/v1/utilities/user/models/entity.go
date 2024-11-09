package models

import "time"

type User struct {
	UserID      string    `json:"user_id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	RoleID      int       `json:"role_id"`
	Password    string    `json:"password"`
	Address     string    `json:"address"`
	DateCreated time.Time `json:"date_created" gorm:"autoCreateTime"`
	DateUpdated time.Time `json:"date_updated" gorm:"autoUpdateTime"`
}
