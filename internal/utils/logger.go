package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	logFile := createLogFile()
	InfoLogger = log.New(logFile, "[INFO]:\t", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(logFile, "[WARN]:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(logFile, "[ERROR]:\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func createLogFile() *os.File {
	absPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get the base dir:", err)
	}
	logDir := filepath.Join(absPath, "log")
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		fmt.Println("Failed to create log directory:", err)
	}
	logPath := filepath.Join(logDir, "log.log")
	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open the log file:", err)
	}
	return logFile
}
