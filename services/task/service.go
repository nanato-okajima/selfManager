package task

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
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
	tasks := repository.FetchTaskList()
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
	status, err := strconv.ParseInt(r.FormValue("status"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	dueDatetime, err := time.Parse(constants.DATE_FORMAT, r.FormValue("due-datetime"))
	if err != nil {
		fmt.Println(err)
	}
	pr := structs.Task{
		Name:        r.FormValue("name"),
		Status:      status,
		DueDatetime: dueDatetime,
	}
	repository.CreateTask(&pr)
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

// TaskHandler is Update task handler
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	task := repository.FetchTask(id)
	tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("edit"), constants.GetHeaderTemplate(), constants.GetFooterTemplate()))
	if err := tmp.ExecuteTemplate(w, "edit.gtpl", task); err != nil {
		log.Println(err)
	}
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	decorder := schema.NewDecoder()
	// status, err := strconv.ParseInt(r.FormValue("status"), 10, 64)
	// if err != nil {
	// 	log.Println(err)
	// }
	decorder.RegisterConverter(time.Time{}, ParseToDateTime)
	vars := mux.Vars(r)
	id := vars["id"]
	// id := strings.TrimPrefix(r.URL.Path, "/task/")
	task := repository.FetchTask(id)
	var req structs.Task
	fmt.Printf("%#v", r.PostForm)
	err = decorder.Decode(&req, r.PostForm)
	if err != nil {
		fmt.Println(err)
	}
	repository.UpdateTask(task, &req)
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func ParseToDateTime(str string) reflect.Value {
	t, err := time.Parse(constants.DATE_FORMAT, str)
	if err != nil {
		log.Println(err)
	}
	return reflect.ValueOf(t)
}
