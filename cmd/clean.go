package cmd

import (
	"fmt"
	"tareas-cli/tareas"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Elimina todas las tareas con estado completado",
	Long:  "Este comando elimina todas las tareas que hayan sido completadas",
	Run: func(cmd *cobra.Command, args []string) {
		lista := tareas.ListTareas()
		var nuevasTareas []tareas.Tarea

		for _, tarea := range lista {
			if !tarea.Completado {
				nuevasTareas = append(nuevasTareas, tarea)
			}
		}

		err := tareas.UpdateTareas(nuevasTareas)
		if err != nil {
			fmt.Println("Error al actualizar la lista de tareas")
			return
		}
		fmt.Println("Lista de tareas actualizada")
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
