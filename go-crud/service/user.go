package service

import (
	"go/go-crud/config"
	"go/go-crud/exception"
	"go/go-crud/modal"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootPathService(context *gin.Context) {
	context.String(http.StatusOK, "Wellcome to CRUD application ")
}

func FetchUserService(context *gin.Context) {
	newUser := []modal.User{}
	config.DB.Find(&newUser)
	context.IndentedJSON(http.StatusOK, newUser)
}

func AddUserService(context *gin.Context) {
	var newUser modal.User

	// Bind json to the object and check if any error
	if error := context.BindJSON(&newUser); error != nil {
		log.Printf("Fatal error occured:%v", error)
		context.Set("error", error)
		return
	}

	if error_persisit := config.DB.Create(&newUser); error_persisit.Error != nil {
		context.Set("error_persisit", error_persisit)
		exception.ErrorHandeler(context)
	} else {
		context.IndentedJSON(http.StatusCreated, newUser) // show newly created object as a json to client

	}
}

func FetchUserByIdServiceService(context *gin.Context) {
	var newUser modal.User
	config.DB.Where("id = ?", context.Param("user-id")).First(&newUser)
	context.IndentedJSON(http.StatusOK, newUser)

}

func UpdateUserByIdService(context *gin.Context) {
	var newUser modal.User
	config.DB.Where("id = ?", context.Param("user-id")).First(&newUser)

	// Bind json to the object and check if any error
	if error := context.BindJSON(&newUser); error != nil {
		log.Printf("Fatal error occured:%v", error)
		context.Set("error", error)
		return
	}

	if error_persisit := (config.DB.Save(&newUser)); error_persisit.Error != nil {
		context.Set("error_persisit", error_persisit)
		exception.ErrorHandeler(context)
	} else {
		context.IndentedJSON(http.StatusCreated, newUser) // show newly created object as a json to client
	}
}
