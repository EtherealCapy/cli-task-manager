package internal

import (
	"fmt"
	"strconv"
	"tareas-cli/pkg"
	"tareas-cli/pkg/models"
	"time"

	"github.com/spf13/cobra"
)

func add(args []string) {
	var task models.Task

	if args[0] == "" || args[0] == " " {
		fmt.Println(" [!] Task title is required")
		return
	}

	task.Title = args[0]
	task.Completed = false
	task.Date = time.Now().Format("Mon, 02 Jan 2006")
	priority, errConv := strconv.Atoi(args[1])

	if errConv != nil {
		fmt.Println(" [!] Wrong priority format. Must be a number between 1 and 3")
		return
	}

	if priority < 1 || priority > 3 {
		fmt.Println(" [!] Wrong priority range")
		return
	}

	task.Priority = priority

	if err := pkg.AddTarea(task); err != nil {
		fmt.Println("[!] Wrong title : ", err)
	} else {
		fmt.Println(" [*] " + task.Title + " added!")
	}
}

var addCmd = &cobra.Command{
	Use:   "add [ Title ] [ Priority ]",
	Short: "Add new task. Priority range: 1-3",
	Long:  "Add new task to the JSON file",
	Run: func(cmd *cobra.Command, args []string) {
		add(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
