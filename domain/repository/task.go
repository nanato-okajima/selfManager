package repository

import "selfManager/domain/model"

type TaskRepository interface {
	Fetch() (*[]model.Task, error)
	Create(*model.Task) error
	Find(id string) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id string) error
}
