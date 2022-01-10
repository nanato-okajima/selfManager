package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string
	Status      int64
	DueDatetime time.Time `schema:"due-datetime"`
}
