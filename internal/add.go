package internal

import (
	"fmt"
	"strconv"
	"tareas-cli/pkg"
	"tareas-cli/pkg/models"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [ Title ] [ Priority ]",
	Short: "Add new task. Priority range: 1-3",
	Long:  "Add new task to the JSON file",
	Run: func(cmd *cobra.Command, args []string) {
		add(cmd, args)
	},
}

func init() {
	addCmd.Flags().StringP("limit", "l", "", "Limit date for the task (YYYY-MM-DD)")
	rootCmd.AddCommand(addCmd)
}

/*
Función que añade una tarea a la lista de tareas
*/
func add(cmd *cobra.Command, args []string) {
	var task models.Task

	if args[0] == "" || args[0] == " " {
		fmt.Println(" [!] Task title is required")
		return
	}

	const (
		MinPriority = 1
		MaxPriority = 3
	)

	priority, errConv := strconv.Atoi(args[1])
	if errConv != nil || priority < MinPriority || priority > MaxPriority {
		fmt.Printf(" [!] Priority must be a number between %d and %d\n", MinPriority, MaxPriority)
		return
	}

	duedate, err := cmd.Flags().GetString("limit")
	if err != nil {
		fmt.Println(" [!] Error getting due date:", err)
		return
	}

	task = models.Task{
		Title:     args[0],
		Completed: false,
		Date:      time.Now().Format("Mon, 02 Jan 2006"),
		Priority:  priority,
		Limit:     duedate,
	}
	
	if priority < 1 || priority > 3 {
		fmt.Println(" [!] Wrong priority range")
		return
	}

	task.Priority = priority

	if duedate == "" {
		task.Limit = "No limit date"
	} else {

		errAddDuedate := addDuedate(duedate, &task)

		if errAddDuedate != nil {
			fmt.Println(" [!] Error adding duedate: ", errAddDuedate)
			return
		}
	}

	if err := pkg.AddTarea(task); err != nil {
		fmt.Println("[!] Wrong title : ", err)
	} else {
		fmt.Println(" [*] " + task.Title + " added!")
	}
}

/*
Función que define las flags del comando
*/
func addDuedate(date string, task *models.Task) error {

	if isValid := checkDateFormat(date); !isValid {
		return fmt.Errorf("wrong date format, must be YYYY-MM-DD")
	}

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return fmt.Errorf("error parsing date: %w", err)
	}

	formattedDate := parsedDate.Format("Mon, 02 Jan 2006")

	task.Limit = formattedDate
	fmt.Println("Limit date set to:", date)

	return nil
}

/*
Función que comprueba el formato de la fecha
*/
func checkDateFormat(date string) bool {
	if _, err := time.Parse("2006-01-02", date); err != nil {
		return false
	}
	return true
}
