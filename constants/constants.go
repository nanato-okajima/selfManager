package constants

const (
	TEMPLATES_DIR  = "templates/"
	TASK_DIR       = "tasks/"
	HTML_EXTENSION = ".html"
)

func GetTaskDirPath(filename string) string {
	return TEMPLATES_DIR + TASK_DIR + filename + HTML_EXTENSION
}
