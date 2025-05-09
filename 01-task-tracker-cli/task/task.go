package task

import (
	"fmt"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
}


var tasks []Task

func AddTask(title, description, status string) {
	task := Task{
		ID:          len(tasks) + 1,
		Title:       title,
		Description: description,
		Status:      status,
	}

	tasks = append(tasks, task)
}

func ListTask() {
	if len(tasks) == 0 {
		fmt.Println("Kayitli görev bulunmamaktadir.")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s\n", task.ID, task.Title, task.Description, task.Status)
	}
}
