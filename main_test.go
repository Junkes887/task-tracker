package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	saveDatabase(Database{
		CountID: 0,
		Tasks:   []Task{},
	})

	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestMethodNotFound(t *testing.T) {
	os.Args = []string{"", "teste"}
	main()
	t.Log("Method not found!")
}

func TestAdd(t *testing.T) {
	os.Args = []string{"", "add", "teste"}
	main()
	t.Logf("Task added successfully (ID: %d)", 1)
}

func TestUpdate(t *testing.T) {
	os.Args = []string{"", "update", "1", "teste update"}
	main()
	t.Logf("Task updated successfully (ID: %d)", 1)
}

func TestList(t *testing.T) {
	os.Args = []string{"", "list"}
	main()
	t.Log("List of all tasks")
}

func TestListTodo(t *testing.T) {
	os.Args = []string{"", "list", "todo"}
	main()
	t.Log("List of all tasks")
}

func TestListDone(t *testing.T) {
	os.Args = []string{"", "list", "done"}
	main()
	t.Log("List of all tasks")
}

func TestListInProgress(t *testing.T) {
	os.Args = []string{"", "list", "in-progress"}
	main()
	t.Log("List of all tasks")
}

func TestMarkInProgress(t *testing.T) {
	os.Args = []string{"", "mark-in-progress", "1"}
	main()
	t.Logf("Task mark to %s successfully (ID: %s) \n", STATUS_IN_PROGRESS, "1")
}

func TestMarkInDone(t *testing.T) {
	os.Args = []string{"", "mark-done", "1"}
	main()
	t.Logf("Task mark to %s successfully (ID: %s) \n", STATUS_DONE, "1")
}

func TestDelete(t *testing.T) {
	os.Args = []string{"", "delete", "1"}
	main()
	t.Logf("Task deleted successfully (ID: %d)", 1)
}
