package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"tareas-cli/pkg/models"
)

/*
Función que carga las tareas desde un archivo JSON

taskFile: nombre del archivo JSON
tasks: posición de memoria de un slice de tareas
*/
func LoadTasks(taskFile string, tasks *[]models.Task) error {
	file, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("error while trying to open file: %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return fmt.Errorf("error while encoding JSON: %v", err)
	}

	return nil
}

/*
Función que guarda las tareas en un archivo JSON

taskFile: nombre del archivo JSON
taskList: slice de tareas
*/
func SaveTasks(taskFile string, taskList []models.Task) error {
	file, err := os.Create(taskFile)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(taskList)
}
