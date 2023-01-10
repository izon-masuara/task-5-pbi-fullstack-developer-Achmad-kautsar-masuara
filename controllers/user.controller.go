package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/izon-masuara/app"
	"github.com/izon-masuara/helpers"
	"github.com/izon-masuara/models"
)

func UserRegister(ctx *gin.Context) {
	var payload *app.PostRequest

	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	resIsEmail := govalidator.IsEmail(payload.Email)
	if !resIsEmail {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Faild format email")
		return
	}

	fileImageName := fmt.Sprintf("%s-%s", time.Now(), payload.Photo.Filename)

	if err := ctx.SaveUploadedFile(&payload.Photo, helpers.Dir+"/images/"+fileImageName); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	photo := app.Photo{
		Filename: fileImageName,
		Size:     int(payload.Photo.Size),
		Mimetype: payload.Photo.Header.Get("Content-Type"),
	}

	hashPass, _ := helpers.HashPassword(payload.Password)

	user := app.User{
		Username: payload.Username,
		Password: hashPass,
		Email:    payload.Email,
		Photo:    photo,
	}

	if err := models.InsertUser(user, fileImageName); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "Success Register User")
}
