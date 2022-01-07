package repository

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var fixtures *testfixtures.Loader

func TestMain(m *testing.M) {
	var err error
	if err = SetupDB("../../.env.test"); err != nil {
		log.Fatal(err)
	}
	Migrate()

	fixtures, err = newFixtures()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func newFixtures() (*testfixtures.Loader, error) {
	DB, _ := db.client.DB()
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
	if err := prepareTestDatabase(); err != nil {
		t.FailNow()
	}

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
			got, err := db.FetchTaskList()
			if assert.NoError(t, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestCreateTask(t *testing.T) {
	if err := prepareTestDatabase(); err != nil {
		t.FailNow()
	}

	tests := []struct {
		name    string
		err     error
		request *Task
	}{
		{
			name: "成功/新規登録",
			err:  nil,
			request: &Task{
				Model: gorm.Model{
					ID:        3,
					CreatedAt: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local),
					DeletedAt: *new(gorm.DeletedAt),
				},
				Name:        "test3",
				Status:      3,
				DueDatetime: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := db.CreateTask(tt.request)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
		})
	}
}

func TestFetchTask(t *testing.T) {
	if err := prepareTestDatabase(); err != nil {
		t.FailNow()
	}

	tests := []struct {
		name string
		err  error
		id   string
		want *Task
	}{
		{
			name: "成功/1件取得",
			err:  nil,
			id:   "1",
			want: &Task{
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.FetchTask(tt.id)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	if err := prepareTestDatabase(); err != nil {
		t.FailNow()
	}

	tests := []struct {
		name    string
		err     error
		request struct {
			task *Task
			req  *Task
		}
	}{
		{
			name: "成功/タスク変更",
			err:  nil,
			request: struct {
				task *Task
				req  *Task
			}{
				task: &Task{
					Model: gorm.Model{
						ID: 1,
					},
				},
				req: &Task{
					Model: gorm.Model{
						ID: 4,
					},
					Name:        "testhoge",
					Status:      2,
					DueDatetime: time.Date(2022, time.March, 3, 0, 0, 0, 0, time.Local),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := db.UpdateTask(tt.request.task, tt.request.req)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	if err := prepareTestDatabase(); err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name string
		err  error
		id   string
	}{
		{
			name: "成功/タスク削除",
			err:  nil,
			id:   "2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := db.DeleteTask(tt.id)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
		})
	}
}
