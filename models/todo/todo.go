package todo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Id         uuid.UUID
	Text       string
	IsComplete bool
	UserId     uint
}
