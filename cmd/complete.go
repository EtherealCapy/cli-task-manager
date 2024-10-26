package cmd

import (
	"fmt"
	"strconv"
	"tareas-cli/tareas"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Recibe un id válido y lo marca como completado",
	Run: func(cmd *cobra.Command, args []string) {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Indice inválido: ", err)
			return
		}

		if err := tareas.CompleteTarea(i); err != nil {
			fmt.Println("Error al completar la tarea: ", err)
		} else {
			fmt.Println("Tarea completada")
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
