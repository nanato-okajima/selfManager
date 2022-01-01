package task

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

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

// TaskCreateHandler is task insert db
func TaskCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("create"), constants.GetHeaderTemplate(), constants.GetFooterTemplate()))
		if err := tmp.ExecuteTemplate(w, "create.gtpl", nil); err != nil {
			log.Println(err)
		}
	} else if r.Method == http.MethodPost {
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
		http.Redirect(w, r, "/tasks", 303)
	}
}

// TaskHandler is Update task handler
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := strings.TrimPrefix(r.URL.Path, "/task/")
		task := repository.FetchTask(id)
		tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("edit"), constants.GetHeaderTemplate(), constants.GetFooterTemplate()))
		if err := tmp.ExecuteTemplate(w, "edit.gtpl", task); err != nil {
			log.Println(err)
		}
	} else if r.Method == http.MethodPost {
		status, err := strconv.ParseInt(r.FormValue("status"), 10, 64)
		if err != nil {
			log.Println(err)
		}
		dueDatetime, err := time.Parse(constants.DATE_FORMAT, r.FormValue("due-datetime"))
		if err != nil {
			log.Println(err)
		}
		id := strings.TrimPrefix(r.URL.Path, "/task/")
		task := repository.FetchTask(id)
		repository.UpdateTask(task, &structs.Task{
			Name:        r.FormValue("name"),
			Status:      status,
			DueDatetime: dueDatetime,
		})
		http.Redirect(w, r, "/tasks", 303)
	}
}
