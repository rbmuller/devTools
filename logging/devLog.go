package devLog

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	TestLogger    *log.Logger
)

const projectdir = "project-name"

func init() {

	projectName := regexp.MustCompile(`^(.*` + projectdir + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	today := time.Now().Format("2006-01-02")

	file, err := os.OpenFile(fmt.Sprintf("./%s/execution_log_%s.txt", rootPath, today), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	TestLogger = log.New(file, "TEST: ", log.Ldate|log.Ltime|log.Lshortfile)
}
