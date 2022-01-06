package repository

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"selfManager/constants"
)

var (
	TDB      *gorm.DB
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	if err := setupDB(); err != nil {
		log.Fatal(err)
	}
	if err := prepareTestDatabase(); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func setupDB() error {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Println(err)
	}
	envconfig.Process("", &env)

	dsn := fmt.Sprintf(constants.DSN, env.Host, env.User, env.Pass, env.DB, env.Port)
	TDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	migrate()
	fixtures, err = newFixtures()
	if err != nil {
		return err
	}

	return nil
}

func migrate() {
	m := TDB.Migrator()
	err := m.AutoMigrate(&Task{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("table create")
}

func newFixtures() (*testfixtures.Loader, error) {
	DB, _ := TDB.DB()
	return testfixtures.New(
		testfixtures.Database(DB),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("../../testdata/fixtures"),
	)
}

func prepareTestDatabase() error {
	if err := fixtures.Load(); err != nil {
		return err
	}
	return nil
}

func TestFetchTaskList(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want *[]Task
	}{
		{
			name: "成功/全件取得",
			err:  nil,
			want: success["全件取得"],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchTaskList(TDB)
			assert.NoError(t, err, "")
			assert.Equal(t, tt.want, got)
		})
	}
}
