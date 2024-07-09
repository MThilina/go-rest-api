package main

import (
	"github.com/gin-gonic/gin"
)

// Best practice -
// struct variable should be capitalize since they are used to serialization
// json filed can be used all should be simple and fields should be mention inside quotes and all mention inside tilda
type todo struct {
	Id        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todoList = []todo{
	{"1", "Clean Room", true},
	{"2", "Wash Car", false},
	{"3", "Organize Book Shelf", false},
	{"4", "Bring Vegitables", true},
	{"5", "Learn new skills", true},
}

func main() {
	// server configuration
	router := gin.Default()
	// get todo list
	router.GET("/todo", getTodos)
	router.POST("/todo", addTodo)
	router.GET("/todo/:todo-id", getTodoById)
	router.PATCH("/todo/:todo-id", updateTodoById)
	router.Run("localhost:9090") // should be dynamically input

}
