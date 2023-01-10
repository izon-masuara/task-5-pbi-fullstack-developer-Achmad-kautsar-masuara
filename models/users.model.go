package models

import (
	"errors"
	"os"

	"github.com/izon-masuara/app"
	"github.com/izon-masuara/database"
	"github.com/izon-masuara/helpers"
)

func InsertUser(userPayload app.User, fileImageName string) error {
	resultUser := database.Db.Create(&userPayload)
	if resultUser.Error != nil {
		os.Remove(helpers.Dir + "/images/" + fileImageName)
		database.Db.Unscoped().Delete(&app.Photo{}, userPayload.Photo_id)
		return errors.New("failed to register")
	}
	return nil
}
