package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"selfManager/domain/structs"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", os.Getenv("HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	DB = db
}

func Migrate() {
	m := DB.Migrator()
	err := m.AutoMigrate(&structs.Task{})
	if err != nil {
		log.Println(err)
	}
	
	fmt.Println("table create")
}

func FetchTaskList() (*[]structs.Task, error) {
	var tasks []structs.Task
	if err := DB.Order("updated_at desc").Find(&tasks).Error; err != nil {
		return nil, err
	}

	return &tasks, nil
}

func CreateTask(request *structs.Task) error {
	if err := DB.Create(request); err != nil {
		return err
	}
	fmt.Println("success create task")

	return nil
}

func FetchTask(id string) (*structs.Task, error) {
	var task structs.Task
	if err := DB.First(&task, id); err != nil {
		return nil, err
	}

	return &task, nil
}

func UpdateTask(task *structs.Task, req *structs.Task) error {
	if err := DB.Model(task).Updates(req); err != nil {
		return err
	}
	fmt.Println("success update task")

	return nil
}

func DeleteTask(id string) error {
	var task structs.Task
	if err := DB.Delete(&task, id); err != nil {
		return err
	}
	fmt.Println("success delete task")

	return nil
}
