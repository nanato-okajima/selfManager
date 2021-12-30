package repository

import "time"

type Task struct {
	Name        string
	DueDatetime time.Time
}
