package router

import (
	"github.com/gin-gonic/gin"
	"github.com/izon-masuara/controllers"
	"github.com/izon-masuara/middlewares"
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
	user.PUT("/:userId", controllers.EditUser)
	user.DELETE("/:userId", controllers.DeleteUser)

	router.Use(middlewares.Authentication)

	router.POST("/photos", controllers.UploadPhoto)
	router.GET("/photos", controllers.GetPhotos)
	router.PUT("/:photoId", controllers.EditPhoto)
	router.DELETE("/:photoId", controllers.DeletePhotoById)
}
