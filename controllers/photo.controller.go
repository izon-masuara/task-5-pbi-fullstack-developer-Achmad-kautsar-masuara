package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izon-masuara/app"
	"github.com/izon-masuara/models"
)

func UploadPhoto(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(float64)
	var payload app.PostPhoto

	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	payloadPhoto := app.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoUrl,
		UserID:   userId,
	}

	if err := models.InsertPhoto(payloadPhoto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "success upload photo")

}

func GetPhotos(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(float64)
	photos, err := models.GetPhotos(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, photos)
}

func EditPhoto(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(float64)
	photoId := ctx.Param("photoId")
	var photo app.PostPhoto
	if err := ctx.ShouldBind(&photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.EditPhoto(userId, photo, photoId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "success updated")
}

func DeletePhotoById(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(float64)
	photoId := ctx.Param("photoId")
	if err := models.DestroyPhotoById(userId, photoId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "success deleted photo")
}
