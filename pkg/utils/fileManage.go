package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"tareas-cli/pkg/models"
)

func LoadTasks(taskFile string, tasks *[]models.Task) error {
	file, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("Error while trying to open file: %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return fmt.Errorf("Error while encoding JSON: %v", err)
	}

	return nil
}

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
