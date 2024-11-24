#!/bin/bash

#Ruta del directorio de tareas
TASKS_DIR="$HOME/.tasks"
TASKS_FILE="$TASKS_DIR/tasks.json"

# Crear el directorio donde se almacenarán las tareas si no existe
if [ ! -d "$TASKS_DIR" ]; then
    mkdir -p "$TASKS_DIR" && echo " [*] .tasks directory created successfully"
fi

# Crear el archivo de tareas si no existe
if [ ! -f "$TASKS_FILE" ]; then
    echo "[]" > "$TASKS_FILE" && echo " [*] Task JSON file created successfully"
fi

# Compilar el binario en un directorio temporal
TMP_DIR=$(mktemp -d)
sudo go build -o "$TMP_DIR/task" ./cmd || { echo "Failed to build task binary"; exit 1; }

# Mover el binario al directorio $GOPATH/bin o $HOME/bin
if [ -d "$GOPATH/bin" ]; then
    sudo mv "$TMP_DIR/task" "$GOPATH/bin/"
    echo "  [*] Command created in GOPATH"
else
    sudo mv "$TMP_DIR/task" "$HOME/bin/"
fi

# Eliminar el directorio temporal
rm -rf "$TMP_DIR"

echo "  [*] Task CLI installed. Execute 'task' from your terminal to start."
