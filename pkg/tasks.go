package pkg

import (
	"os"
	"tareas-cli/pkg/models"
	"tareas-cli/pkg/utils"
)

var taskFile = os.ExpandEnv("$HOME/.tasks/tasks.json")
var taskList []models.Task

/*
Función que añade una tarea al archivo JSON

t: tarea a añadir
*/
func AddTarea(t models.Task) error {
	if err := utils.LoadTasks(taskFile, &taskList); err != nil {
		return err
	}

	taskList = append(taskList, t)
	return utils.SaveTasks(taskFile, taskList)
}

/*
Función que elimina una tarea del archivo JSON

i: índice de la tarea a eliminar
*/
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

/*
Función que lista las tareas del archivo JSON
*/
func ListTareas() []models.Task {
	if err := utils.LoadTasks(taskFile, &taskList); err != nil {
		return nil
	}

	return taskList
}

/*
Función que marca una tarea como completada

i: índice de la tarea a completar
*/
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

/*
Función que actualiza la lista de tareas en el archivo JSON

t: lista de tareas a guardar
*/
func UpdateTareas(t []models.Task) error {
	taskList = t
	return utils.SaveTasks(taskFile, taskList)
}
