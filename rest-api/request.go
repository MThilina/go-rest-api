package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log"
)

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todoList)
}

func addTodo(context *gin.Context) {
	var newTodo todo //create a Todo object (from todo struct)

	// Bind json to the object and check if any error
	if error := context.BindJSON(&newTodo); error != nil {
		log.Printf("Fatal error occured:%v", error)
		return
	}

	todoList = append(todoList, newTodo) // append the new todo to exisiting list

	context.IndentedJSON(http.StatusCreated, newTodo) // show newly created object as a json to client

}

func getTodoById(context *gin.Context) {
	todoId := context.Param("todo-id")
	todo, error := getTodoByIdHelper(todoId)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, error.Error())
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func updateTodoById(context *gin.Context) {
	todoId := context.Param("todo-id")
	todo, error := getTodoByIdHelper(todoId)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, error.Error())
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)

}
