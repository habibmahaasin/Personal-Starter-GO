package models

import "time"

// User model definition
type User struct {
    UserID      string        `gorm:"primaryKey;column:user_id"`
    FullName    string        `gorm:"column:full_name"`
    Email       string        `gorm:"column:email"`
    RoleID      int64         `gorm:"column:role_id"`
    Password    string        `gorm:"column:password"`
    Address     string        `gorm:"column:address"`
    DateCreated time.Time     `gorm:"column:date_created"`
    DateUpdated time.Time     `gorm:"column:date_updated"`
    CheckInLogs []CheckInLog  `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE;"`
    UserStats   UserStats     `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE;"`
}

type CheckInLog struct {
    ID          uint      `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
    UserID      string    `gorm:"column:user_id" json:"user_id"`
    Image       string    `gorm:"column:image" json:"image"`
    Note        string    `gorm:"column:note" json:"note"`
    DateCreated time.Time `gorm:"column:date_created" json:"date_created"`
    DateUpdated time.Time `gorm:"column:date_updated" json:"date_updated"`
}

type UserStats struct {
    UserID           string    `gorm:"primaryKey;column:user_id"`
    TotalCheckIn     int       `json:"total_check_in"`
    IsPreTested      bool      `json:"is_pre_tested"`
    IsPostTested     bool      `json:"is_post_tested"`
    IsAvailableToRedeem bool   `json:"is_available_to_redeem"`
    IsRedeemedReward bool      `json:"is_redeemed_reward"`
    DateCreated      time.Time `json:"date_created"`
    DateUpdated      time.Time `json:"date_updated"`
}