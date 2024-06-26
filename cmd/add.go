package cmd

import (
	"fmt"
	"strconv"
	"tareas-cli/tareas"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [Titulo] [Prioridad]",
	Short: "Agrega un nuevo pendiente, rango de prioridad (1-3)",
	Long:  "Este comando agrega un nuevo pendiente a la lista",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			fmt.Println("Especifique una prioridad a su tarea")
			return
		}

		var tarea tareas.Tarea
		tarea.Titulo = args[0]
		tarea.Completado = false
		prioridad, errConv := strconv.Atoi(args[1])

		if errConv != nil {
			fmt.Println("Error parámetro de prioridad incorrecto")
			return
		}

		if prioridad < 1 || prioridad > 3 {
			fmt.Println("Rango de prioridad incorrecto")
			return
		}

		tarea.Prioridad = prioridad

		if err := tareas.AddTarea(tarea); err != nil {
			fmt.Println("Error al añadir la tarea: ", err)
		} else {
			fmt.Println(tarea.Titulo + " añadida a la lista de tareas")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
