package repository

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"selfManager/constants"
)

var DB *gorm.DB

type Env struct {
	Host     string
	TestHost string `envconfig:"TEST_HOST"`
	User     string `envconfig:"POSTGRES_USER"`
	Pass     string `envconfig:"POSTGRES_PASSWORD"`
	DB       string `envconfig:"POSTGRES_DB"`
	Port     string `envconfig:"POSTGRES_PORT"`
	TestDB   string `envconfig:"TEST_POSTGRES_DB"`
}

var env Env

func SetupDB() {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Println(err)
	}
	envconfig.Process("", &env)

	dsn := fmt.Sprintf(constants.DSN, env.Host, env.User, env.Pass, env.DB, env.Port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
}

func Migrate() {
	m := DB.Migrator()
	err := m.AutoMigrate(&Task{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("table create")
}

func FetchTaskList(db *gorm.DB) (*[]Task, error) {
	var tasks []Task
	if err := db.Order("updated_at desc").Find(&tasks).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return &tasks, nil
}

func CreateTask(request *Task) error {
	if err := DB.Create(request).Error; err != nil {
		return errors.WithStack(err)
	}
	fmt.Println("success create task")

	return nil
}

func FetchTask(id string) (*Task, error) {
	var task Task
	if err := DB.First(&task, id).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return &task, nil
}

func UpdateTask(task *Task, req *Task) error {
	if err := DB.Model(task).Updates(req).Error; err != nil {
		return errors.WithStack(err)
	}
	fmt.Println("success update task")

	return nil
}

func DeleteTask(id string) error {
	var task Task
	if err := DB.Delete(&task, id).Error; err != nil {
		return errors.WithStack(err)
	}
	fmt.Println("success delete task")

	return nil
}
