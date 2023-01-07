package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/izon-masuara/app"
	"github.com/izon-masuara/database"
	"github.com/izon-masuara/helpers"
)

func UserRegister(ctx *gin.Context) {
	var payload *app.PostRequest

	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	fileImageName := fmt.Sprintf("%s-%s", time.Now(), payload.Photo.Filename)

	dir, _ := os.Getwd()
	if err := ctx.SaveUploadedFile(&payload.Photo, dir+"/images/"+fileImageName); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	photo := app.Photo{
		Filename: fileImageName,
		Size:     int(payload.Photo.Size),
		Mimetype: payload.Photo.Header.Get("Content-Type"),
	}

	resultPhoto := database.Db.Create(&photo)
	if resultPhoto.Error != nil {
		os.Remove(dir + "/images/" + fileImageName)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resultPhoto.Error.Error())
		return
	}

	hashPass, _ := helpers.HashPassword(payload.Password)
	resIsEmail := govalidator.IsEmail(payload.Email)
	if !resIsEmail {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Faild format email")
		return
	}

	user := app.User{
		Username: payload.Username,
		Password: hashPass,
		Email:    payload.Email,
		Photo_Id: photo.ID,
	}

	resultUser := database.Db.Create(&user)
	if resultUser.Error != nil {
		os.Remove(dir + "/images/" + fileImageName)
		database.Db.Unscoped().Delete(&app.Photo{}, photo.ID)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resultUser.Error.Error())
		return
	}

	ctx.JSON(http.StatusOK, "Success Register User")
}
