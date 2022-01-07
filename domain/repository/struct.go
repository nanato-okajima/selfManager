package repository

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

type Env struct {
	Host     string
	TestHost string `envconfig:"TEST_HOST"`
	User     string `envconfig:"POSTGRES_USER"`
	Pass     string `envconfig:"POSTGRES_PASSWORD"`
	DB       string `envconfig:"POSTGRES_DB"`
	Port     string `envconfig:"POSTGRES_PORT"`
	TestDB   string `envconfig:"TEST_POSTGRES_DB"`
}

type DBClient interface {
	FetchTaskList() (*[]Task, error)
	CreateTask(db *gorm.DB, request *Task) error
	FetchTask(db *gorm.DB, id string) (*Task, error)
	UpdateTask(db *gorm.DB, task *Task, req *Task) error
	DeleteTask(db *gorm.DB, id string) error
}

type DB struct {
	client *gorm.DB
}
