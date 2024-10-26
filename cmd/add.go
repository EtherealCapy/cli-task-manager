package cmd

import (
	"fmt"
	"strconv"
	"tareas-cli/tareas"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [Titulo] [Prioridad]",
	Short: "Agrega un nuevo pendiente, rango de prioridad (1-3)",
	Long:  "Este comando agrega un nuevo pendiente a la lista",
	Run: func(cmd *cobra.Command, args []string) {

		var tarea tareas.Tarea

		if args[0] == "" || args[0] == " " {
			fmt.Println(" [!] Ingrese un título para la tarea")
			return
		}

		tarea.Titulo = args[0]
		tarea.Completado = false
		tarea.Fecha = time.Now().Format("Mon, 02 Jan 2006")
		prioridad, errConv := strconv.Atoi(args[1])

		if errConv != nil {
			fmt.Println(" [!] Error parámetro de prioridad incorrecto")
			return
		}

		if prioridad < 1 || prioridad > 3 {
			fmt.Println(" [!] Rango de prioridad incorrecto")
			return
		}

		tarea.Prioridad = prioridad

		if err := tareas.AddTarea(tarea); err != nil {
			fmt.Println("[!] Error al añadir la tarea: ", err)
		} else {
			fmt.Println(" [*] " + tarea.Titulo + " añadida a la lista de tareas")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
