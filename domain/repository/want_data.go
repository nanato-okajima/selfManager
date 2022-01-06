package repository

import (
	"time"

	"gorm.io/gorm"
)

var success = map[string]*[]Task{
	"全件取得": {
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local),
				UpdatedAt: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local),
				DeletedAt: *new(gorm.DeletedAt),
			},
			Name:        "test1",
			Status:      1,
			DueDatetime: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local),
		},
		{
			Model: gorm.Model{
				ID:        2,
				CreatedAt: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local),
				UpdatedAt: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local),
				DeletedAt: *new(gorm.DeletedAt),
			},
			Name:        "test2",
			Status:      2,
			DueDatetime: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local),
		},
	},
}
