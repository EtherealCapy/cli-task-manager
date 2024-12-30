package internal

import (
	"fmt"
	"math"
	"strconv"
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
		updateActiveDays(args)
	},
}

func init() {
	rootCmd.AddCommand(updateActiveCmd)
}

func updateActiveDays(args []string) {

	if args[0] == "" || args[0] == " " {
		fmt.Println(" [!] Task ID is required")
		return
	}

	id, errConv := strconv.Atoi(args[0])

	if errConv != nil {
		fmt.Println(" [!] Wrong ID format. Must be a number")
		return
	}

	id--

	if id < 0 {
		fmt.Println(" [!] Invalid ID: out of range")
		return
	}

	task, err := pkg.GetTarea(id)

	if err != nil {
		fmt.Println(" [!] Task not found")
		return
	}

	activeDays := calculateActiveDays(task)

	if errUpdate := pkg.UpdateActiveDays(id, activeDays); errUpdate != nil {
		fmt.Println(" [!] Error while updating active days")
		return
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
