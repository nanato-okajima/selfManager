package task

import (
	"log"
	"net/http"
	"selfManager/constants"
	"text/template"

	"github.com/selfManager/constants"
	"github.com/selfManager/domain/repository"
)

func TaskListHandler(w http.ResponseWriter, r *http.Request) {
	tasks := repository.FetchTaskList()
	tmp := template.Must(template.ParseFiles(constants.GetTaskDirPath("list")))
	if err := tmp.ExecuteTemplate(w, "list.html", tasks); err != nil {
		log.Println(err)
	}
}
