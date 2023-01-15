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
		return errors.New(resultUser.Error.Error())
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

func EditUserWithId(id string, payload app.PostRequest) error {
	var user app.User

	founded := database.Db.Where("id = ?", id).Find(&user).Scan(&user)
	if founded.RowsAffected == 0 {
		return errors.New("data not found")
	}

	newPass, err := helpers.HashPassword(payload.Password)
	if err != nil {
		return errors.New("try again")
	}

	var plainPayload = app.User{
		Username: payload.Username,
		Password: newPass,
		Email:    payload.Email,
		Photo:    user.Photo,
	}

	database.Db.Where("id = ?", id).UpdateColumns(plainPayload)

	return nil
}

func DestroyUser(id string) error {
	res := database.Db.Unscoped().Where("id = ?", id).Delete(&app.User{})
	if res.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func FindByUserId(id float64) error {
	res := database.Db.Where("id = ?", id).Find(&app.User{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
