package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	STATUS_TODO        = "todo"
	STATUS_IN_PROGRESS = "in-progress"
	STATUS_DONE        = "done"
)

type Database struct {
	CountID int
	Tasks   []Task
}

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {
	var args = os.Args
	var operation = args[1]
	switch operation {
	case "list":
		list(args)
	case "add":
		add(args)
	case "update":
		update(args)
	case "delete":
		delete(args)
	case "mark-in-progress":
		updateStatus(args[2], STATUS_IN_PROGRESS)
	case "mark-done":
		updateStatus(args[2], STATUS_DONE)
	default:
		fmt.Println("Method not found!")
	}
}

func list(args []string) {
	fmt.Println("List of all tasks")
	database := getDatabase()

	if len(args) > 2 {
		switch args[2] {
		case "done":
			database.Tasks = filterList(database.Tasks, STATUS_DONE)
		case "todo":
			database.Tasks = filterList(database.Tasks, STATUS_TODO)
		case "in-progress":
			database.Tasks = filterList(database.Tasks, STATUS_IN_PROGRESS)
		}
	}

	for _, v := range database.Tasks {
		fmt.Printf("ID: %v, ", v.ID)
		fmt.Printf("Description: %v, ", v.Description)
		fmt.Printf("Status: %v, ", v.Status)
		fmt.Printf("CreatedAt: %v, ", v.CreatedAt)
		fmt.Printf("UpdatedAt: %v \n", v.UpdatedAt)
	}
}

func filterList(list []Task, status string) []Task {
	var filterList []Task

	for _, v := range list {
		if v.Status == status {
			filterList = append(filterList, v)
		}
	}

	return filterList
}

func add(args []string) {
	database := getDatabase()
	database.CountID++

	var newTask = Task{
		ID:          database.CountID,
		Description: args[2],
		Status:      STATUS_TODO,
		CreatedAt:   time.Now(),
	}

	database.Tasks = append(database.Tasks, newTask)

	saveDatabase(database)
	fmt.Printf("Task added successfully (ID: %d) \n", database.CountID)
}

func update(args []string) {
	database := getDatabase()
	id, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(database.Tasks); i++ {
		if database.Tasks[i].ID == id {
			database.Tasks[i].Description = args[3]
			database.Tasks[i].UpdatedAt = time.Now()
		}
	}

	saveDatabase(database)
	fmt.Printf("Task updated successfully (ID: %s) \n", args[2])
}

func delete(args []string) {
	database := getDatabase()
	id, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(database.Tasks); i++ {
		if database.Tasks[i].ID == id {
			database.Tasks = append(database.Tasks[:i], database.Tasks[i+1:]...)
		}
	}

	saveDatabase(database)
	fmt.Printf("Task deleted successfully (ID: %s) \n", args[2])
}

func updateStatus(id string, status string) {
	database := getDatabase()
	idConvert, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(database.Tasks); i++ {
		if database.Tasks[i].ID == idConvert {
			database.Tasks[i].Status = status
		}
	}

	saveDatabase(database)
	fmt.Printf("Task mark to %s successfully (ID: %s) \n", status, id)
}

func getDatabase() Database {
	var database Database
	file, err := os.ReadFile("database.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &database)

	return database
}

func saveDatabase(database Database) {
	taksJson, _ := json.Marshal(database)
	if err := os.WriteFile("database.json", taksJson, 0644); err != nil {
		panic(err)
	}
}
