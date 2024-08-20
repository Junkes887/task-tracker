# Task Tracker

![coverage](https://img.shields.io/badge/coverage-93%25-brightgreen)


Task tracker is a project used to track and manage your tasks.

Example
The list of commands and their usage is given below:

````bash
# Adding a new task
go run main.go add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
go run main.go update 1 "Buy groceries and cook dinner"
go run main.go delete 1

# Marking a task as in progress or done
go run main.go mark-in-progress 1
go run main.go mark-done 1

# Listing all tasks
go run main.go list

# Listing tasks by status
go run main.go list done
go run main.go list todo
go run main.go list in-progress`
````
### Task Properties

* ID: A unique identifier for the task
* Description: A short description of the task
* Status: The status of the task (todo, in-progress, done)
* CreatedAt: The date and time when the task was created
* UpdatedAt: The date and time when the task was last updated

### Idea
Project idea basis [roadmap.sh](https://roadmap.sh/projects/task-tracker)

### Test

````bash
# Run test
go test -cover

# Run test and view in html
go test -coverprofile cover.out 
go tool cover -html="cover.out"
````