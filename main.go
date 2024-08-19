package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

var countId = 0

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
		list()
	case "add":
		add(args)
	case "update":
		update(args)
	case "delete":
		delete(args)
	default:
		fmt.Println("Method not found!")
	}
}

func list() {
	database := getDatabase()
	for _, v := range database.Tasks {
		fmt.Printf("ID: %v, ", v.ID)
		fmt.Printf("Description: %v, ", v.Description)
		fmt.Printf("Status: %v, ", v.Status)
		fmt.Printf("CreatedAt: %v, ", v.CreatedAt)
		fmt.Printf("UpdatedAt: %v \n", v.UpdatedAt)
	}
}

func add(args []string) {
	database := getDatabase()
	database.CountID++

	var newTask = Task{
		ID:          database.CountID,
		Description: args[2],
		Status:      "todo",
		CreatedAt:   time.Now(),
	}

	database.Tasks = append(database.Tasks, newTask)

	saveDatabase(database)
	fmt.Println("Add task " + strconv.Itoa(database.CountID))
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
	fmt.Println("Update task " + args[2])
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
	fmt.Println("Delete task " + args[2])
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
