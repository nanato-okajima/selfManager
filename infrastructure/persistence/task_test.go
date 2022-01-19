package persistence

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"selfManager/config"
	"selfManager/domain/model"
	"selfManager/domain/repository"
)

var fixtures *testfixtures.Loader
var tr repository.TaskRepository

func TestMain(m *testing.M) {
	config.SetEnv("../../.env.test")
	tr = NewTaskPersistence(config.Connect())

	var err error
	fixtures, err = newFixtures()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func newFixtures() (*testfixtures.Loader, error) {
	db, _ := config.Connect().DB()
	return testfixtures.New(
		testfixtures.Database(db),
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

func TestFetch(t *testing.T) {
	if err := prepareTestDatabase(); err != nil {
		t.FailNow()
	}

	tests := map[string]struct {
		err  error
		want *[]model.Task
	}{
		"成功/全件取得": {
			err:  nil,
			want: success["全件取得"],
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tr.Fetch()
			if assert.NoError(t, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	if err := prepareTestDatabase(); err != nil {
		t.FailNow()
	}

	tests := map[string]struct {
		err     error
		request *model.Task
	}{
		"成功/新規登録": {
			err: nil,
			request: &model.Task{
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

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := tr.Create(tt.request)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
		})
	}
}

func TestFind(t *testing.T) {
	if err := prepareTestDatabase(); err != nil {
		t.FailNow()
	}

	tests := map[string]struct {
		err  error
		id   string
		want *model.Task
	}{
		"成功/1件取得": {
			err: nil,
			id:  "1",
			want: &model.Task{
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

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tr.Find(tt.id)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	if err := prepareTestDatabase(); err != nil {
		t.FailNow()
	}

	tests := map[string]struct {
		err     error
		request struct {
			task *model.Task
			req  *model.Task
		}
	}{
		"成功/タスク変更": {
			err: nil,
			request: struct {
				task *model.Task
				req  *model.Task
			}{
				req: &model.Task{
					Model: gorm.Model{
						ID: 1,
					},
					Name:        "testhoge",
					Status:      2,
					DueDatetime: time.Date(2022, time.March, 3, 0, 0, 0, 0, time.Local),
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := tr.Update(tt.request.req)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
		})
	}
}

func TestDelete(t *testing.T) {
	if err := prepareTestDatabase(); err != nil {
		log.Fatal(err)
	}

	tests := map[string]struct {
		name string
		err  error
		id   string
	}{
		"成功/タスク削除": {
			err: nil,
			id:  "2",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := tr.Delete(tt.id)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
		})
	}
}
