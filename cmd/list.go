package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strings"
	"tareas-cli/tareas"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar todas las tareas existentes",
	Long:  "Este comando muestra todas las tareas",
	Run: func(cmd *cobra.Command, args []string) {

		tareas, err := tareas.ListTareas()
		if err != nil {
			fmt.Println("Error al cargar las tareas")
			return
		}

		separator := strings.Repeat("-", 58)
		header := fmt.Sprintf("%-25s %-18s %-18s", " ", "Tareas", " ")
		cabecera := fmt.Sprintf("%s\n%s\n", separator, header)
		fmt.Printf(cabecera)

		if len(tareas) != 0 {
			fmt.Println(strings.Repeat("-", 58))
			fmt.Printf("%-4s %-30s %-12s %-10s\n", "ID", "Título", "Estado", "Prioridad")
			fmt.Println(strings.Repeat("-", 58))
			for i, tarea := range tareas {
				estado := "Pendiente"
				if tarea.Completado {
					estado = "Completado"
				}

				var prioridad string
				var prioridadColor *color.Color

				switch tarea.Prioridad {
				case 1:
					prioridad = "Baja"
					prioridadColor = color.New(color.FgGreen)
				case 2:
					prioridad = "Media"
					prioridadColor = color.New(color.FgYellow)
				default:
					prioridad = "Alta"
					prioridadColor = color.New(color.FgRed)
				}

				estadoColor := color.New(color.FgYellow)
				if estado == "Completado" {
					estadoColor = color.New(color.FgGreen)
				}

				fmt.Printf("%-4d %-30s ", i+1, tarea.Titulo)
				estadoColor.Printf("%-12s ", estado)
				prioridadColor.Printf("%-10s\n", prioridad)
			}
		} else {
			fmt.Println(strings.Repeat("-", 58))
			fmt.Println(" [!] No tienes pendientes")
			fmt.Println("\t")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
