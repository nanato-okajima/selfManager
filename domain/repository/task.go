package repository

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"selfManager/constants"
)

var db DB
var env Env

func SetupDB(envPath string) error {
	err := godotenv.Load(envPath)
	if err != nil {
		return err
	}
	envconfig.Process("", &env)

	dsn := fmt.Sprintf(constants.DSN, env.Host, env.User, env.Pass, env.DB, env.Port)
	db.client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func Migrate() error {
	m := db.client.Migrator()
	err := m.AutoMigrate(&Task{})
	if err != nil {
		return err
	}

	fmt.Println("table create")
	return nil
}

func (db *DB) FetchTaskList() (*[]Task, error) {
	var tasks []Task
	if err := db.client.Order("updated_at desc").Find(&tasks).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return &tasks, nil
}

func (db DB) CreateTask(request *Task) error {
	if err := db.client.Create(request).Error; err != nil {
		return errors.WithStack(err)
	}
	fmt.Println("success create task")

	return nil
}

func (db *DB) FetchTask(id string) (*Task, error) {
	var task Task
	if err := db.client.First(&task, id).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return &task, nil
}

func (db *DB) UpdateTask(task *Task, req *Task) error {
	if err := db.client.Model(task).Updates(req).Error; err != nil {
		return errors.WithStack(err)
	}
	fmt.Println("success update task")

	return nil
}

func (db *DB) DeleteTask(id string) error {
	var task Task
	if err := db.client.Delete(&task, id).Error; err != nil {
		return errors.WithStack(err)
	}
	fmt.Println("success delete task")

	return nil
}
