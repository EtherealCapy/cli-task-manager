package utils

import "github.com/fatih/color"

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
