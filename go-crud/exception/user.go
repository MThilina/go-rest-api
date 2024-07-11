package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Global exception handler
func ErrorHandeler(context *gin.Context) {

	marshalErr := context.GetString("error")          // marshalErr!="" jaxonbinding errors are there
	persistentErr := context.GetBool("error_persist") // false - persistant erors are occuring

	// Check if there is an error and modulate response accordingly
	if marshalErr != "" || !persistentErr {
		modulateResponse(context, marshalErr)
	}

}

func modulateResponse(context *gin.Context, marshalErr string) {
	var errorResponse gin.H
	if marshalErr != "" {

		errorResponse = gin.H{
			"message": marshalErr, // directly use the error string
			"code":    http.StatusBadRequest,
		}
	} else {
		errorResponse = gin.H{
			"message": "Database Persistance Error Occured", // directly use the error string
			"code":    http.StatusBadRequest,
		}
	}
	// Simulate an error response

	// Send indented JSON response with status code 400 and related error
	context.IndentedJSON(http.StatusBadRequest, errorResponse)
}
