package todo

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	UserID     uint
	Text       string
	IsComplete bool
}
