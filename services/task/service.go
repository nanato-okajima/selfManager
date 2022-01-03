package task

import (
	"log"
	"net/http"
	"reflect"
	"text/template"
	"time"

	"github.com/gorilla/schema"
	"github.com/pkg/errors"

	"selfManager/constants"
	"selfManager/domain/repository"
	"selfManager/domain/structs"
)

// TaskListHandler is select task
func taskList(w http.ResponseWriter) error {
	tasks, err := repository.FetchTaskList()
	if err != nil {
		return err
	}

	tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("list"), constants.GetHeaderTemplate(), constants.GetFooterTemplate()))
	if err := tmp.ExecuteTemplate(w, "list.gtpl", tasks); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func newTask(w http.ResponseWriter) error {
	tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("create"), constants.GetHeaderTemplate(), constants.GetFooterTemplate()))
	if err := tmp.ExecuteTemplate(w, "create.gtpl", nil); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// TaskCreateHandler is task insert db
func createTask(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return errors.WithStack(err)
	}

	decorder := schema.NewDecoder()
	decorder.RegisterConverter(time.Time{}, parseToDateTime)

	var req structs.Task
	if err := decorder.Decode(&req, r.PostForm); err != nil {
		return errors.WithStack(err)
	}

	if err := repository.CreateTask(&req); err != nil {
		return err
	}

	return nil
}

// TaskHandler is Update task handler
func task(w http.ResponseWriter, id string) error {
	task, err := repository.FetchTask(id)
	if err != nil {
		return err
	}

	tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("edit"), constants.GetHeaderTemplate(), constants.GetFooterTemplate()))
	if err := tmp.ExecuteTemplate(w, "edit.gtpl", task); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func updateTask(r *http.Request, id string) error {
	if err := r.ParseForm(); err != nil {
		return errors.WithStack(err)
	}

	decorder := schema.NewDecoder()
	decorder.RegisterConverter(time.Time{}, parseToDateTime)

	task, err := repository.FetchTask(id)
	if err != nil {
		return err
	}

	var req structs.Task
	if err = decorder.Decode(&req, r.PostForm); err != nil {
		return errors.WithStack(err)
	}

	if err = repository.UpdateTask(task, &req); err != nil {
		return err
	}

	return nil
}

func deleteTask(id string) error {
	if err := repository.DeleteTask(id); err != nil {
		return err
	}

	return nil
}

func parseToDateTime(str string) reflect.Value {
	t, err := time.Parse(constants.DATE_FORMAT, str)
	if err != nil {
		log.Println(err)
	}

	return reflect.ValueOf(t)
}
