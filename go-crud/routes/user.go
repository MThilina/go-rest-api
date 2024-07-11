package routes

import (
	"go/go-crud/controller"

	"github.com/gin-gonic/gin"
)

// render the root path
func FetchUser(router *gin.Engine) {
	router.GET("/", controller.RootPath)
	router.GET("/user", controller.FetchUser)
	router.POST("/user", controller.AddUser)
	router.GET("/user/:user-id", controller.FetchUserById)
}
