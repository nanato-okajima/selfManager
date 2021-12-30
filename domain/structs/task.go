package structs

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string
	DueDatetime time.Time
	Status      int64
}
