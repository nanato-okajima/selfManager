package task

import (
	"time"

	"github.com/pkg/errors"

	"selfManager/domain/model"
	"selfManager/domain/repository"
)

type TaskUseCase interface {
	FetchList() (*[]model.Task, error)
	Create(*Request) error
	Find(string) (*model.Task, error)
	Update(string, *Request) error
	Delete(string) error
}

type Request struct {
	Name        string
	Status      int64
	DueDatetime time.Time `schema:"due-datetime"`
}

type taskUseCase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUseCase(tr repository.TaskRepository) TaskUseCase {
	return &taskUseCase{
		taskRepository: tr,
	}
}

func (tu *taskUseCase) FetchList() (*[]model.Task, error) {
	tasks, err := tu.taskRepository.Fetch()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return tasks, nil
}

func (tu *taskUseCase) Create(req *Request) error {
	task := model.Task{
		Name:        req.Name,
		Status:      req.Status,
		DueDatetime: req.DueDatetime,
	}
	err := tu.taskRepository.Create(&task)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (tu *taskUseCase) Find(id string) (*model.Task, error) {
	if id == "" {
		return nil, errors.WithStack(errors.New("idが不正です"))
	}

	task, err := tu.taskRepository.Find(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return task, nil
}

func (tu *taskUseCase) Update(id string, req *Request) error {
	if id == "" {
		return errors.WithStack(errors.New("idが不正です"))
	}

	task, err := tu.taskRepository.Find(id)
	if err != nil {
		return errors.WithStack(err)
	}

	task.Name = req.Name
	task.Status = req.Status
	task.DueDatetime = req.DueDatetime

	if err := tu.taskRepository.Update(task); err != nil {
		return err
	}

	return nil
}

func (tu *taskUseCase) Delete(id string) error {
	if id == "" {
		return errors.WithStack(errors.New("idが不正です"))
	}

	if err := tu.taskRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
