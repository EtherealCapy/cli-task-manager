package tareas

import (
	"encoding/json"
	"fmt"
	"os"
)

var tareasFile = os.ExpandEnv("$HOME/.tasks/tareas.json")

type Tarea struct {
	Titulo     string
	Prioridad  int
	Completado bool
	Fecha      string
}

var tareas []Tarea

func loadTareas() error {
	file, err := os.Open(tareasFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&tareas)
	if err != nil {
		return fmt.Errorf("error al decodificar JSON: %v", err)
	}

	return nil
}

func saveTareas() error {
	file, err := os.Create(tareasFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tareas)
}

func AddTarea(t Tarea) error {
	if err := loadTareas(); err != nil {
		return err
	}

	tareas = append(tareas, t)
	return saveTareas()
}

func RemTarea(i int) error {
	if err := loadTareas(); err != nil {
		return err
	}

	if i < 0 || i >= len(tareas) {
		return os.ErrInvalid
	}
	tareas = append(tareas[:i], tareas[i+1:]...)
	return saveTareas()
}

func ListTareas() []Tarea {
	if err := loadTareas(); err != nil {
		return nil
	}

	return tareas
}

func CompleteTarea(i int) error {
	if err := loadTareas(); err != nil {
		return err
	}
	if i >= 0 && i <= len(tareas) {
		tareas[i-1].Completado = true
		return saveTareas()
	}

	return os.ErrInvalid
}

func UpdateTareas(t []Tarea) error {
	tareas = t
	return saveTareas()
}
