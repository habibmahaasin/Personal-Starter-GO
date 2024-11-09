package models

import "time"

// User model definition
type User struct {
    UserID      string           `gorm:"primaryKey;column:user_id"`
    FullName    string           `gorm:"column:full_name"`
    Email       string           `gorm:"column:email"`
    RoleID      int64            `gorm:"column:role_id"`
    Password    string           `gorm:"column:password"`
    Address     string           `gorm:"column:address"`
    DateCreated time.Time        `gorm:"column:date_created"`
    DateUpdated time.Time        `gorm:"column:date_updated"`
    DailyCheckInLogs []DailyCheckInLog `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE;"`
}

type DailyCheckInLog struct {
    ID          string    `gorm:"primaryKey;column:id"`
    UserID      string    `gorm:"column:user_id"`
    Date        time.Time `gorm:"column:date"`
    Time        time.Time `gorm:"column:time"`
    Note        string    `gorm:"column:note"`
    DateCreated time.Time `gorm:"column:date_created"`
    DateUpdated time.Time `gorm:"column:date_updated"`
}

