package main

import (
	"fmt"
	"time"
)

// Task represents one piece of work in a project.
type Task struct {
	Title     string
	Done      bool
	CreatedAt time.Time
}

// Project owns a slice of Task structs. This is an example of nesting one
// struct type inside another struct type.
type Project struct {
	Name  string
	Tasks []Task
}

// AddTask is a method: it belongs to Project. The pointer receiver (*Project)
// lets this method modify the original project's Tasks slice.
func (p *Project) AddTask(title string) {
	p.Tasks = append(p.Tasks, Task{
		Title:     title,
		CreatedAt: time.Now(),
	})
}

// CompleteTask updates one task. It returns true when the requested task
// exists, and false when its index is invalid.
func (p *Project) CompleteTask(index int) bool {
	if index < 0 || index >= len(p.Tasks) {
		return false
	}

	p.Tasks[index].Done = true
	return true
}

// PrintSummary is a read-only method, so it uses a value receiver (Project)
// rather than a pointer receiver.
func (p Project) PrintSummary() {
	fmt.Println("Project:", p.Name)

	for _, task := range p.Tasks {
		status := "not done"
		if task.Done {
			status = "done"
		}

		fmt.Printf("- [%s] %s (created %s)\n", status, task.Title, task.CreatedAt.Format(time.Kitchen))
	}
}

func main() {
	// Make one Project struct. Its Tasks slice starts empty.
	project := Project{Name: "Learn Go structs"}

	// Methods add Task structs to the project.
	project.AddTask("Define a struct type")
	project.AddTask("Use a pointer receiver")
	project.AddTask("Loop through a slice of structs")

	// Mark the second task (index 1) as complete.
	project.CompleteTask(1)

	// Print the final data stored in the nested structs.
	project.PrintSummary()
}
