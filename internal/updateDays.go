package internal

import (
	"fmt"
	"math"
	"tareas-cli/pkg"
	"tareas-cli/pkg/models"
	"time"

	"github.com/spf13/cobra"
)

var updateActiveCmd = &cobra.Command{
	Use:   "updateActive [ Task ID ]",
	Short: "Update task active days",
	Long:  "Update the active days for a task",
	Run: func(cmd *cobra.Command, args []string) {
		updateActiveDays()
	},
}

func init() {
	rootCmd.AddCommand(updateActiveCmd)
}

func updateActiveDays() {

	tasks, err := pkg.ListTareas()

	if err != nil {
		fmt.Println("Error while loading tasks")
		return
	}

	if len(tasks) != 0 {

		for i := range tasks {

			task, err := pkg.GetTarea(i)

			if err != nil {
				fmt.Println(" [!] Task not found")
				return
			}

			activeDays := calculateActiveDays(task)

			if errUpdate := pkg.UpdateActiveDays(i, activeDays); errUpdate != nil {
				fmt.Println(" [!] Error while updating active days")
				return
			}
		}
	}
}

func calculateActiveDays(task *models.Task) int {
	layout := "Mon, 02 Jan 2006"
	taskDate, err := time.Parse(layout, task.Date)
	if err != nil {
		return 0
	}

	currentTime := time.Now()
	diff := currentTime.Sub(taskDate)

	activeDays := int(math.Floor(diff.Hours() / 24))

	if activeDays < 0 {
		return 0
	}

	return activeDays
}
