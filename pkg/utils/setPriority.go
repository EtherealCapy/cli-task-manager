package utils

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