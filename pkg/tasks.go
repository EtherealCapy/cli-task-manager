package pkg

import (
	"os"
	"tareas-cli/pkg/models"
	"tareas-cli/pkg/utils"
)

var taskFile = os.ExpandEnv("$HOME/.tasks/tasks.json")
var taskList []models.Task

func AddTarea(t models.Task) error {
	if err := utils.LoadTasks(taskFile, &taskList); err != nil {
		return err
	}

	taskList = append(taskList, t)
	return utils.SaveTasks(taskFile, taskList)
}

func RemTarea(i int) error {
	if err := utils.LoadTasks(taskFile, &taskList); err != nil {
		return err
	}

	if i < 0 || i >= len(taskList) {
		return os.ErrInvalid
	}
	taskList = append(taskList[:i], taskList[i+1:]...)
	return utils.SaveTasks(taskFile, taskList)
}

func ListTareas() []models.Task {
	if err := utils.LoadTasks(taskFile, &taskList); err != nil {
		return nil
	}

	return taskList
}

func CompleteTarea(i int) error {
	if err := utils.LoadTasks(taskFile, &taskList); err != nil {
		return err
	}
	if i >= 0 && i <= len(taskList) {
		taskList[i-1].Completed = true
		return utils.SaveTasks(taskFile, taskList)
	}

	return os.ErrInvalid
}

func UpdateTareas(t []models.Task) error {
	taskList = t
	return utils.SaveTasks(taskFile, taskList)
}
