package handler

import (
	"log"
	"net/http"
	"reflect"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	"selfManager/constants"
	"selfManager/usecase/task"
)

type TaskHandler interface {
	Index(http.ResponseWriter, *http.Request)
	New(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Show(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type taskHandler struct {
	taskUseCase task.TaskUseCase
}

func NewTaskHandler(tu task.TaskUseCase) TaskHandler {
	return &taskHandler{
		taskUseCase: tu,
	}
}

func (th *taskHandler) Index(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.taskUseCase.FetchList()
	if err != nil {
		log.Printf("%+v", err)
	}

	tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("list"), constants.GetHeaderTemplate(), constants.GetFooterTemplate()))
	if err := tmp.ExecuteTemplate(w, "list.gtpl", tasks); err != nil {
		log.Printf("%+v", err)
	}
}

func (th *taskHandler) New(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("create"), constants.GetHeaderTemplate(), constants.GetFooterTemplate()))
	if err := tmp.ExecuteTemplate(w, "create.gtpl", nil); err != nil {
		log.Println(err)
	}
}

func (th *taskHandler) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	var req task.Request
	decorder := schema.NewDecoder()
	decorder.RegisterConverter(time.Time{}, parseToDateTime)
	if err := decorder.Decode(&req, r.PostForm); err != nil {
		log.Printf("%+v", err)
	}

	if err := th.taskUseCase.Create(&req); err != nil {
		log.Printf("%+v", err)
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)

}

// TaskHandler is Update task handler
func (th *taskHandler) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	task, err := th.taskUseCase.Find(vars["id"])
	if err != nil {
		log.Printf("%+v", err)
	}

	tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("edit"), constants.GetHeaderTemplate(), constants.GetFooterTemplate()))
	if err := tmp.ExecuteTemplate(w, "edit.gtpl", task); err != nil {
		log.Fatal(err)
	}
}

func (th *taskHandler) Update(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("%+v", err)
	}

	var req task.Request
	decorder := schema.NewDecoder()
	decorder.RegisterConverter(time.Time{}, parseToDateTime)
	if err := decorder.Decode(&req, r.PostForm); err != nil {
		log.Printf("%+v", err)
	}

	vars := mux.Vars(r)
	if err := th.taskUseCase.Update(vars["id"], &req); err != nil {
		log.Printf("%+v", err)
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func (th *taskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := th.taskUseCase.Delete(vars["id"]); err != nil {
		log.Printf("%+v", err)
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func parseToDateTime(str string) reflect.Value {
	t, err := time.Parse(constants.DATE_FORMAT, str)
	if err != nil {
		log.Println(err)
	}

	return reflect.ValueOf(t)
}
