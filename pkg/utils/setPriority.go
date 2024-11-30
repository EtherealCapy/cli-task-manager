package utils

/*
Función que establece la prioridad de la tarea

priority: int - Prioridad de la tarea
*/
func SetPriority(priority int) string {
	switch priority {
	case 1:
		return "Low"
	case 2:
		return "Medium"
	default:
		return "High"
	}
}
