package app

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"varchar(50),not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Photo    Photo  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:UserID"`
}

type Photo struct {
	gorm.Model
	Title    string `gorm:"varchar(100),not null"`
	Caption  string `gorm:"varchar(255),not null"`
	PhotoUrl string `gorm:"varchar(100),not null"`
	UserID   int
}
