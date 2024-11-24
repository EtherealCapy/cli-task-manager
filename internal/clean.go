package internal

import (
	"fmt"
	"tareas-cli/pkg"
	"tareas-cli/pkg/models"

	"github.com/spf13/cobra"
)

func clean() {
	taskList := pkg.ListTareas()
	var newTasks []models.Task

	for _, task := range taskList {
		if !task.Completed {
			newTasks = append(newTasks, task)
		}
	}

	err := pkg.UpdateTareas(newTasks)
	if err != nil {
		fmt.Println("Error while updating task list")
		return
	}
	fmt.Println("Task list updated successfully")
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete all completed tasks",
	Long:  "Delete all marked as completed tasks from the JSON file",
	Run: func(cmd *cobra.Command, args []string) {
		clean()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
