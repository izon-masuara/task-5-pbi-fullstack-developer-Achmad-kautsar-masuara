package app

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint
	Username string
	Email    *string
	Password string
}

type Pohoto struct {
}
