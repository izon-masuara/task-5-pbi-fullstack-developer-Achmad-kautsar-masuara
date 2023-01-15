package router

import (
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
	user.POST("/login", controllers.UserLogin)
	user.PUT("/:userId")
	user.DELETE("/:userId")

	router.POST("/photos", controllers.UploadPhoto)
}
