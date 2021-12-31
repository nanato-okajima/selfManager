package task

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/selfManager/constants"
	"github.com/selfManager/domain/repository"
	"github.com/selfManager/domain/structs"
)

func TaskListHandler(w http.ResponseWriter, r *http.Request) {
	tasks := repository.FetchTaskList()
	tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("list"), constants.GetHeaderTemplate(), constants.GetFooterTemplate()))
	if err := tmp.ExecuteTemplate(w, "list.gtpl", tasks); err != nil {
		log.Println(err)
	}
}

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
