package internal

import (
	"fmt"
	"os"
	"tareas-cli/pkg"
	"tareas-cli/pkg/utils"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all taks",
	Long:  "List all the tasks saved in ~/.tasks/tasks.json",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list() {
	tasks, err := pkg.ListTareas()

	if err != nil {
		fmt.Println("Error while loading tasks")
		return
	}

	if len(tasks) != 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"#", "title", "status", "priority", "created at", "duedate"})

		for i, task := range tasks {
			status := "Pending"

			if task.Completed {
				status = "Completed"
			}

			priority := utils.SetPriority(task.Priority)

			statuscolor := utils.GetColorForStatus(status)
			prioritycolor := utils.GetColorForPriority(priority)

			if task.Limit == "" {
				task.Limit = "No duedate"
			}

			table.Append([]string{
				fmt.Sprintf("%d ", i+1),
				task.Title,
				statuscolor.Sprintf("%s", status),
				prioritycolor.Sprintf("%s", priority),
				task.Date,
				task.Limit,
			})
		}

		table.Render()
	} else {
		fmt.Println("\t")
		fmt.Println(" [!] No pending tasks")
		fmt.Println("\t")
	}
}
