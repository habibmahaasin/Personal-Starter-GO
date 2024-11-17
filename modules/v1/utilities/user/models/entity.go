package models

import "time"

type User struct {
    UserID      string       `gorm:"primaryKey;column:user_id"`
    FullName    string       `gorm:"column:full_name"`
    Email       string       `gorm:"column:email"`
    RoleID      int64        `gorm:"column:role_id"`
    Password    string       `gorm:"column:password"`
    Address     string       `gorm:"column:address"`
    DateCreated time.Time    `gorm:"column:date_created"`
    DateUpdated time.Time    `gorm:"column:date_updated"`
}