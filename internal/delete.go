package internal

import (
	"fmt"
	"strconv"
	"tareas-cli/pkg"

	"github.com/spf13/cobra"
)

func remove(args []string) {
	index, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Wrong index:", err)
		return
	}

	if err := pkg.RemTarea(index - 1); err != nil {
		fmt.Println("Error while deleting the task:", err)
	} else {
		fmt.Println("Task deleted")
	}
}

var removeCmd = &cobra.Command{
	Use:   "rm [ Index ]",
	Short: "Reiceve an index and remove the task",
	Long:  "Remove the task from the JSON file with the index passed as argument",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		remove(args)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
