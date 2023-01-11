package controllers

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/izon-masuara/app"
	"github.com/izon-masuara/helpers"
	"github.com/izon-masuara/models"
)

func UserRegister(ctx *gin.Context) {
	var payload app.PostRequest

	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	resIsEmail := govalidator.IsEmail(payload.Email)
	if !resIsEmail {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Faild format email")
		return
	}

	hashPass, _ := helpers.HashPassword(payload.Password)

	user := app.User{
		Username: payload.Username,
		Password: hashPass,
		Email:    payload.Email,
		Photo:    app.Photo{},
	}

	if err := models.InsertUser(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "Success Register User")
}

func UserLogin(ctx *gin.Context) {
	var payload app.GetLoginRequest

	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	checkEmail := govalidator.IsEmail(payload.Email)
	if !checkEmail {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "check your format email")
		return
	}

	founded, err := models.UserLogin(payload)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	token, err := helpers.GenerateToken(founded)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, token)
}
