package main

import "fmt"

type Logger struct {
	logs []string
}

func NewLogger() Logger {
	return Logger{
		logs: make([]string, 0, 1000),
	}
}

func (log *Logger) Info(info string) {
	logItem := "INFO: " + info
	log.logs = append(log.logs, logItem)
}

func (log *Logger) Debug(info string) {
	logItem := "DEBUG: " + info
	log.logs = append(log.logs, logItem)
}

func (log *Logger) Error(info string) {
	logItem := "ERROR: " + info
	log.logs = append(log.logs, logItem)
}

func (log *Logger) DrainToConsole() {
	for _, li := range log.logs {
		fmt.Println(li)
	}
}
