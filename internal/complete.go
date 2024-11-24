package internal

import (
	"fmt"
	"strconv"
	"tareas-cli/pkg"

	"github.com/spf13/cobra"
)

func complete(args []string) {
	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid index: ", err)
		return
	}

	if err := pkg.CompleteTarea(i); err != nil {
		fmt.Println("Error while completing: ", err)
	} else {
		fmt.Println("Task completed!")
	}
}

var completeCmd = &cobra.Command{
	Use:   "complete [ Index ]",
	Short: "Mark as coompleted a task by its index",
	Long:  "Change the status of a task to completed by its index in the JSON task list.",
	Run: func(cmd *cobra.Command, args []string) {
		complete(args)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
