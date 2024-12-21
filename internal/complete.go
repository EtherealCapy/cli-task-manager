package internal

import (
	"fmt"
	"strconv"
	"tareas-cli/pkg"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [ Index ]",
	Short: "Mark as coompleted a task by its index",
	Long:  "Change the status of a task to completed by its index in the JSON task list.",
	Run: func(cmd *cobra.Command, args []string) {
		toggleFlag, _ := cmd.Flags().GetBool("toggle")
		complete(args, toggleFlag)
	},
}

func init() {
	completeCmd.Flags().BoolP("toggle", "t", false, "Toggle complete status")
	rootCmd.AddCommand(completeCmd)
}

func complete(args []string, toggle bool) {
	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid index, try with a valid one")
		return
	}
	
	if toggle {
		if err := pkg.ToggleTarea(i - 1); err != nil {
			fmt.Println("Error while toggling: ", err)
		} else {
			fmt.Println("Task toggled!")
		}
		return
	}
	
	if err := pkg.CompleteTarea(i - 1); err != nil {
		fmt.Println("Error while completing: ", err)
	} else {
		fmt.Println("Task completed!")
	}
}
