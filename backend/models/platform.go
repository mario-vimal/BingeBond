package models

import (
	"time"
)

type Platform struct {
	Id            int32  `gorm:"primary_key;AUTO_INCREMENT;not null"`
	UserId        int32  `gorm:"not null" json:"user_id"`
	PatformName   string `gorm:"not null"`
	PlatformToken string
	PlatformId    string `gorm:"not null"`
	IsLoggedIn    bool
	IsFavourite   bool
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
}
