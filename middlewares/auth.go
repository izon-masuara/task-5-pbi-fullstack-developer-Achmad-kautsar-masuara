package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/izon-masuara/helpers"
	"github.com/izon-masuara/models"
)

func Authentication(ctx *gin.Context) {
	token := ctx.GetHeader("access_token")
	res, err := helpers.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "you do not have access")
		return
	}
	userId := res.(jwt.MapClaims)["UserID"].(float64)
	if err := models.FindByUserId(userId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "you do not have access")
		return
	}
	ctx.Set("userId", userId)
	ctx.Next()
}
