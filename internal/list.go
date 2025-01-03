package internal

import (
	"fmt"
	"os"
	"tareas-cli/pkg"
	"tareas-cli/pkg/utils"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var countPending bool
var countTotal bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all the tasks saved in ~/.tasks/tasks.json",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&countPending, "count", "c", false, "Count the number of pending tasks")
	listCmd.Flags().BoolVarP(&countTotal, "count-total", "t", false, "Count the total number of tasks")
}

func list() {
	tasks, err := pkg.ListTareas()

	if err != nil {
		fmt.Println("Error while loading tasks")
		return
	}

	if countPending {
		count := 0
		for _, task := range tasks {
			if !task.Completed {
				count++
			}
		}
		fmt.Printf("Number of pending tasks: %d\n", count)
		return
	}

	if countTotal {
		fmt.Printf("Total number of tasks: %d\n", len(tasks))
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
