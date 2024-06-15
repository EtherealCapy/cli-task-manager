# Task Manager (CLI)

## Overview
This task management application is a command-line tool designed to help you manage your tasks efficiently from the terminal. This tool allows you to add, delete, list, and complete tasks, as well as clean up completed tasks.

## Features
- **Add Task**: Adds a new task with a specified priority.
- **Remove Task**: Deletes a task by its index.
- **List Tasks**: Displays all tasks with their statuses and priorities.
- **Complete Task**: Marks a task as completed.
- **Clean Tasks**: Deletes all tasks that have been marked as completed.

## Installation
Clone the repository, and navigate to the project directory:
```sh
  git clone https://github.com/EtherealCapy/cli-task-manager
  cd cli-task-manager
```
Ejecuta el instalador
```sh
  ./install.sh
```

## Usage
**Add**
```sh
  task add ["Name"] [1-3 Priority range]
```

**Remove**
```sh
  task rm [Index]
```

**List**
```sh
  task list
```

**Complete**
```sh
  task complete [Index]
```

**Clean**
```sh
  task clean
```
## Qtile Intregration
The `tasks.py` file contains a script that loads the number of each priority tasks into a Qtile widget. Just add this file to `/usr/locale/bin` and use the command
```sh
  tasks.py
```
Use the `sample_widget.py` to guide on how to create the final widget. Do not change the name of each TextBox or the command must not work.
