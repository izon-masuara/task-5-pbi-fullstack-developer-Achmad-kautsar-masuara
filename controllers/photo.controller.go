package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izon-masuara/app"
	"github.com/izon-masuara/models"
)

func UploadPhoto(ctx *gin.Context) {
	var payload app.PostPhoto

	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	payloadPhoto := app.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoUrl,
		UserID:   1,
	}

	if err := models.InsertPhoto(payloadPhoto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "success upload photo")

}
