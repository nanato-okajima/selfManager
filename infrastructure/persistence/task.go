package persistence

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"selfManager/domain/model"
	"selfManager/domain/repository"
)

type taskPersistence struct {
	con *gorm.DB
}

func NewTaskPersistence(con *gorm.DB) repository.TaskRepository {
	return &taskPersistence{con}
}

func (t *taskPersistence) Fetch() (*[]model.Task, error) {
	var tasks []model.Task
	if err := t.con.Order("updated_at desc").Find(&tasks).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return &tasks, nil
}

func (t *taskPersistence) Create(task *model.Task) error {
	if err := t.con.Create(task).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (t *taskPersistence) Find(id string) (*model.Task, error) {
	var task model.Task
	if err := t.con.First(&task, id).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return &task, nil
}

func (t *taskPersistence) Update(task *model.Task) error {
	if err := t.con.Updates(task).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (t *taskPersistence) Delete(id string) error {
	var task model.Task
	if err := t.con.Delete(&task, id).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
