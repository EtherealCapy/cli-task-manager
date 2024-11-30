package models

type Task struct {
	Title      string
	Priority   int
	Completed  bool
	CompleteAt string
	Date       string
	Limit      string
	ActiveDays int
}
