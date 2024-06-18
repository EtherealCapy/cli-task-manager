package cmd

import (
	"fmt"
	"strconv"
	"tareas-cli/tareas"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "rm [indice]",
	Short: "Recibe un id válido y lo elimina del archivo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Seleccione un índice válido")
			return
		}

		if err := tareas.RemTarea(index - 1); err != nil {
			fmt.Println("La tarea no existe")
		} else {
			fmt.Println("Tarea eliminada")
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
