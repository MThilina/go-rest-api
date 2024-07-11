package controller

import (
	"go/go-crud/service"

	"github.com/gin-gonic/gin"
)

func RootPath(context *gin.Context) {
	service.RootPathService(context)
}

func FetchUser(context *gin.Context) {
	service.FetchUserService(context)
}

func AddUser(context *gin.Context) {
	service.AddUserService(context)

}

func FetchUserById(context *gin.Context) {
	service.FetchUserByIdServiceService(context)
}

func UpdateUserById(context *gin.Context) {
	service.UpdateUserByIdService(context)

}
