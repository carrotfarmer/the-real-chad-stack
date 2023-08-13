package user

import (
	"gorm.io/gorm"

	"carrotfarmer/chad-stack/models/todo"
)

type User struct {
	gorm.Model
	Name    string
	Picture string
	Email   string
	Todos   []todo.Todo
}
