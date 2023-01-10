package app

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Filename string `gorm:"type:varchar(255);not null"`
	Size     int    `gorm:"not null"`
	Mimetype string `gorm:"not null"`
}

type User struct {
	gorm.Model
	Username string `gorm:"varchar(255),not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Photo_id int
	Photo    Photo `gorm:"foreignKey:Photo_id;constraint:OnDelete:CASCADE;"`
}
