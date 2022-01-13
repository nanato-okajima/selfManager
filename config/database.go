package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"selfManager/constants"
	"selfManager/domain/model"
)

var env Env

type Env struct {
	Host string
	User string `envconfig:"POSTGRES_USER"`
	Pass string `envconfig:"POSTGRES_PASSWORD"`
	DB   string `envconfig:"POSTGRES_DB"`
	Port string `envconfig:"POSTGRES_PORT"`
}

func Connect() *gorm.DB {
	dsn := fmt.Sprintf(constants.DSN, env.Host, env.User, env.Pass, env.DB, env.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(model.Task{}); err != nil {
		return err
	}

	fmt.Println("table create")
	return nil
}

func SetEnv(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	envconfig.Process("", &env)
	return nil
}
