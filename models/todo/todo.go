package todo

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Text       string
	IsComplete bool
	UserId     uint
}
