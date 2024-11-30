package utils

import "github.com/fatih/color"

/*
Función que devuelve el color correspondiente a la prioridad de la tarea

priority: prioridad de la tarea
*/
func GetColorForPriority(priority string) *color.Color {
	switch priority {
	case "Low":
		return color.New(color.FgGreen)
	case "Medium":
		return color.New(color.FgYellow)
	case "High":
		return color.New(color.FgRed)
	default:
		return color.New(color.FgWhite)
	}
}

/*
Función que devuelve el color correspondiente al estado de la tarea

status: estado de la tarea
*/
func GetColorForStatus(status string) *color.Color {
	switch status {
	case "Pending":
		return color.New(color.FgYellow)
	case "Completed":
		return color.New(color.FgGreen)
	default:
		return color.New(color.FgWhite)
	}
}
