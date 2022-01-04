package constants

const (
	TEMPLATES_DIR  = "templates/"
	TASK_DIR       = "tasks/"
	HEADER_FILE    = "_header"
	FOOTER_FILE    = "_footer"
	GTPL_EXTENSION = ".gtpl"
	DATE_FORMAT    = "2006-01-02"
	DSN            = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo"
)

func GetTaskDirPath(filename string) string {
	return TEMPLATES_DIR + TASK_DIR + filename + GTPL_EXTENSION
}

func GetHeaderTemplate() string {
	return TEMPLATES_DIR + HEADER_FILE + GTPL_EXTENSION
}

func GetFooterTemplate() string {
	return TEMPLATES_DIR + FOOTER_FILE + GTPL_EXTENSION
}
