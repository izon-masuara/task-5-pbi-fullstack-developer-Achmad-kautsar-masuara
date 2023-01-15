package models

import (
	"errors"
	"fmt"

	"github.com/izon-masuara/app"
	"github.com/izon-masuara/database"
)

func InsertPhoto(payload app.Photo) error {
	resultInsert := database.Db.Create(&payload)
	if resultInsert.Error != nil {
		fmt.Println(resultInsert.Error)
		return errors.New("failed upload photo")
	}

	return nil
}
