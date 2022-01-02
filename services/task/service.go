package task

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	"selfManager/constants"
	"selfManager/domain/repository"
	"selfManager/domain/structs"
)

// TaskListHandler is select task
func TaskListHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := repository.FetchTaskList()
	if err != nil {
		log.Println(err)
	}

	tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("list"), constants.GetHeaderTemplate(), constants.GetFooterTemplate()))
	if err := tmp.ExecuteTemplate(w, "list.gtpl", tasks); err != nil {
		log.Println(err)
	}
}

func NewTaskHandler(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("create"), constants.GetHeaderTemplate(), constants.GetFooterTemplate()))
	if err := tmp.ExecuteTemplate(w, "create.gtpl", nil); err != nil {
		log.Println(err)
	}
}

// TaskCreateHandler is task insert db
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	decorder := schema.NewDecoder()
	decorder.RegisterConverter(time.Time{}, ParseToDateTime)

	var req structs.Task
	err = decorder.Decode(&req, r.PostForm)
	if err != nil {
		fmt.Println(err)
	}
	err = repository.CreateTask(&req)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

// TaskHandler is Update task handler
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	task, err := repository.FetchTask(id)
	if err != nil {
		log.Println(err)
	}

	tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("edit"), constants.GetHeaderTemplate(), constants.GetFooterTemplate()))
	if err := tmp.ExecuteTemplate(w, "edit.gtpl", task); err != nil {
		log.Println(err)
	}
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	decorder := schema.NewDecoder()
	decorder.RegisterConverter(time.Time{}, ParseToDateTime)

	vars := mux.Vars(r)
	id := vars["id"]
	task, err := repository.FetchTask(id)
	if err != nil {
		log.Println(err)
	}

	var req structs.Task
	err = decorder.Decode(&req, r.PostForm)
	if err != nil {
		log.Println(err)
	}
	err = repository.UpdateTask(task, &req)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func TaskDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := repository.DeleteTask(id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func ParseToDateTime(str string) reflect.Value {
	t, err := time.Parse(constants.DATE_FORMAT, str)
	if err != nil {
		log.Println(err)
	}

	return reflect.ValueOf(t)
}
