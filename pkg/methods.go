package pkg

import (
	"errors"
	"os"
	"tareas-cli/pkg/models"
	"tareas-cli/pkg/utils"
)

var taskFile = os.ExpandEnv("$HOME/.tasks/tasks.json")
var taskList []models.Task

/*
Función que devuelve la posición de memoria de una tarea

i: índice de la tarea a recuperar
*/
func GetTarea(i int) (*models.Task, error) {
	var err error

	if err = utils.LoadTasks(taskFile, &taskList); err != nil {
		err = os.ErrNotExist
		return nil, err
	}

	if i < 0 || i >= len(taskList) {
		err = os.ErrInvalid
		return nil, err
	}

	return &taskList[i], nil
}

/*
Función que añade una tarea al archivo JSON

t: tarea a añadir
*/
func AddTarea(t models.Task) error {
	var err error

	if err = utils.LoadTasks(taskFile, &taskList); err != nil {
		return err
	}

	taskList = append(taskList, t)

	if err = utils.SaveTasks(taskFile, taskList); err != nil {
		return err
	}

	return nil
}

/*
Función que elimina una tarea del archivo JSON

i: índice de la tarea a eliminar
*/
func RemTarea(i int) error {
	var err error

	if err = utils.LoadTasks(taskFile, &taskList); err != nil {
		return err
	}

	if i < 0 || i >= len(taskList) {
		return os.ErrInvalid
	}

	taskList = append(taskList[:i], taskList[i+1:]...)

	if err = utils.SaveTasks(taskFile, taskList); err != nil {
		return err
	}

	return nil
}

/*
Función que lista las tareas del archivo JSON
*/
func ListTareas() ([]models.Task, error) {
	if err := utils.LoadTasks(taskFile, &taskList); err != nil {
		return nil, err
	}

	return taskList, nil
}

/*
Función que marca una tarea como completada

i: índice de la tarea a completar
*/
func CompleteTarea(i int) error {
	var err error

	if err = utils.LoadTasks(taskFile, &taskList); err != nil {
		return err
	}

	if i < 0 || i >= len(taskList) {
		return errors.New("task index out of range")
	}

	taskList[i].Completed = true

	if err := utils.SaveTasks(taskFile, taskList); err != nil {
		return errors.New("failed to save tasks")
	}

	return nil
}

/*
Función que actualiza los días activos

i: índice de la tarea a actualizar
days: días activos
*/
func UpdateActiveDays(i int, days int) error {
	var err error

	if days < 0 {
		err = errors.New("days cannot be negative")
		return err
	}

	if err = utils.LoadTasks(taskFile, &taskList); err != nil {
		err = errors.New("failed to load tasks")
		return err
	}

	if i < 0 || i >= len(taskList) {
		err = errors.New("invalid task index")
		return err
	}

	task := &taskList[i]
	task.ActiveDays = days

	if err := utils.SaveTasks(taskFile, taskList); err != nil {
		err = errors.New("failed to save tasks")
		return err
	}

	return nil
}

/*
Función que actualiza la lista de tareas en el archivo JSON

t: lista de tareas a guardar
*/
func UpdateTareas(t []models.Task) error {
	taskList = t

	if err := utils.SaveTasks(taskFile, taskList); err != nil {
		return errors.New("failed to save tasks")
	}

	return nil
}

/*
Función que varia el estado de una tarea

i: índice de la tarea a variar
*/
func ToggleTarea(i int) error {

	if err := utils.LoadTasks(taskFile, &taskList); err != nil {
		return err
	}

	if i < 0 || i >= len(taskList) {
		return errors.New("index out of range")
	}

	task := &taskList[i]
	task.Completed = !task.Completed

	return utils.SaveTasks(taskFile, taskList)
}
