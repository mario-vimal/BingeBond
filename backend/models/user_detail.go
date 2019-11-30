package models

import (
	"time"
)

type UserDetail struct {
	UserId        int32  `gorm:"primary_key;AUTO_INCREMENT;not null" json:"id" binding:"required"`
	Email         string `gorm:"not null;type:varchar(100);unique" json:"email" binding:"required"`
	FullName      string `gorm:"not null" json:"fullname" binding:"required"`
	UserToken     string
	Password      string    `gorm:"not null" json:"password" binding:"required"`
	ContactNumber string    `gorm:"not null;unique" json:"contactnumber" binding:"required"`
	CreatedAt     time.Time `gorm:"not null" json:"created_at" binding:"required"`
	UpdatedAt     time.Time `gorm:"not null" json:"updated_at" binding:"required"`
}

type Loginmodel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
