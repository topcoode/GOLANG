package main

import (
	"log"
	"os"
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func init() {
	// Open or create log files
	infoFile, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open info log file:", err)
	}

	warningFile, err := os.OpenFile("warning.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open warning log file:", err)
	}

	errorFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open error log file:", err)
	}

	// Initialize loggers
	infoLogger = log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(warningFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// Example usage
	infoLogger.Println("This is an information log.")
	warningLogger.Println("This is a warning log.")
	errorLogger.Println("This is an error log.")
}
