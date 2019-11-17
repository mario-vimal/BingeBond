package models

import (
	"time"
)

type UserDetail struct {
	UserId        int32  `gorm:"primary_key;AUTO_INCREMENT;not null" json:"user_id"`
	Email         string `gorm:"not null;type:varchar(100);unique"`
	FullName      string `gorm:"not null"`
	UserToken     string
	Password      string    `gorm:"not null"`
	ContactNumber string    `gorm:"not null;unique"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
}
