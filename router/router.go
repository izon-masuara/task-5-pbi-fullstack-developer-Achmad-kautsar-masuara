package router

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func Routers(router *gin.Engine) {
	// router to handle user endpoints
	user := router.Group("/users")
	user.POST("/register")
	user.GET("/login", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello world")
	})
	user.PUT("/:userId")
	user.DELETE("/:userId")
}
