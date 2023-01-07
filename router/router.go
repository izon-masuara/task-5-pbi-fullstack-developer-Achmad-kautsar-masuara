package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/izon-masuara/controllers"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func Routers(router *gin.Engine) {
	// router to handle user endpoints
	user := router.Group("/users")
	user.POST("/register", controllers.UserRegister)
	user.GET("/login", func(ctx *gin.Context) {
		dir, _ := os.Getwd()
		ctx.JSON(http.StatusOK, dir)
	})
	user.PUT("/:userId")
	user.DELETE("/:userId")
}
