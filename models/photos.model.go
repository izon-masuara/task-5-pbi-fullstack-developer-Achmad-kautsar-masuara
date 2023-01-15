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

func GetPhotos(id float64) (app.Photo, error) {
	var photos app.Photo
	result := database.Db.Where("user_id = ?", id).Find(&photos).Scan(&photos)
	if result.Error != nil {
		return photos, result.Error
	}
	return photos, nil
}

func EditPhoto(userId float64, payload app.PostPhoto, photoId string) error {
	var photo app.Photo
	res := database.Db.Where("user_id = ?", int(userId)).Where("id = ?", photoId).Find(&photo).Statement.Scan(&photo)
	if res.RowsAffected == 0 {
		return errors.New("data not found")
	}
	var newPhoto = app.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoUrl,
	}
	updated := database.Db.Where(&photo).UpdateColumns(newPhoto)
	if updated.RowsAffected == 0 {
		return updated.Error
	}
	return nil
}

func DestroyPhotoById(userId float64, photoId string) error {
	res := database.Db.Where("user_id = ?", int(userId)).Where("id = ?", photoId).Delete(&app.Photo{}).Statement
	if res.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}
