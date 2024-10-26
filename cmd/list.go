package cmd

import (
	"fmt"
	"tareas-cli/tareas"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listar todas las tareas existentes",
	Long:  "Este comando muestra todas las tareas",
	Run: func(cmd *cobra.Command, args []string) {
		tareas := tareas.ListTareas()
		if tareas == nil {
			fmt.Println("Error al cargar las tareas")
			return
		}

		fmt.Println("\n  -------------------")
		fmt.Println("	TAREAS")
		fmt.Println("  -------------------")

		if len(tareas) != 0 {
			for i, tarea := range tareas {
				estado := "Pendiente"
				if tarea.Completado {
					estado = "Completado"
				}

				var prioridad string

				if tarea.Prioridad == 1 {
					prioridad = "Baja"
				} else if tarea.Prioridad == 2 {
					prioridad = "Media"
				} else {
					prioridad = "Alta"
				}

				fmt.Printf("\n[%d] %s", i+1, tarea.Titulo)
				if estado == "Pendiente" {
					color.Set(color.FgYellow)
					fmt.Printf("\t[%s]", estado)
					color.Unset()
				} else {
					color.Set(color.FgGreen)
					fmt.Printf("\t[%s]", estado)
					color.Unset()
				}

				if prioridad == "Baja" {
					color.Set(color.FgGreen)
					fmt.Printf("\t[%s]", prioridad)
					color.Unset()
				} else if prioridad == "Media" {
					color.Set(color.FgYellow)
					fmt.Printf("\t[%s]", prioridad)
					color.Unset()
				} else {
					color.Set(color.FgRed)
					fmt.Printf("\t[%s]", prioridad)
					color.Unset()
				}

				fmt.Printf("\t[%s]", tarea.Fecha)
			}
			fmt.Println("\n")
		} else {
			fmt.Println(" [!] No tienes pendientes")
			fmt.Println("\t")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
