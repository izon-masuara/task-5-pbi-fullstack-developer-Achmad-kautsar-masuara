package models

import (
	"errors"

	"github.com/izon-masuara/app"
	"github.com/izon-masuara/database"
	"github.com/izon-masuara/helpers"
)

func InsertUser(userPayload app.User) error {
	resultUser := database.Db.Create(&userPayload)
	if resultUser.Error != nil {
		return errors.New("failed to register")
	}
	return nil
}

func UserLogin(payload app.GetLoginRequest) (app.User, error) {
	var users app.User
	res := database.Db.Where("email = ?", payload.Email).Find(&users).Scan(&users)
	errMsg := errors.New("username or password are wrong")
	if res.Error != nil {
		return app.User{}, errMsg
	}

	checkPass := helpers.CheckPasswordHash(payload.Password, users.Password)
	if !checkPass {
		return app.User{}, errMsg
	}

	return users, nil
}
