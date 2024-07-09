package main

import (
	"errors"
)

func getTodoByIdHelper(id string) (*todo, error) {

	for i, todo := range todoList {
		if todo.Id == id {
			return &todoList[i], nil
		}
	}

	return nil, errors.New("todo details not found")
}
