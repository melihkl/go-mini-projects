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

func DeleteTask(id int) {
	index := -1

	for i, task := range tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Printf("ID %d nolu görev silindi\n", id)

	if index == -1 {
		fmt.Printf("ID %d nolu görev bulunamadı\n", id)
		return
	}
}

func UpdateTask(id int, title, description, status string) {
	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			found = true
			tasks[i].Title = title
			tasks[i].Description = description
			tasks[i].Status = status
			fmt.Printf("ID %d nolu görev güncellendi\n", id)
			return
		}
	}

	if !found {
		fmt.Printf("ID %d nolu görev bulunmamaktadır\n", id)
	}
}

func Exists(id int) bool {
	for _, t := range tasks {
		if t.ID == id {
			return true
		}
	}
	return false
}
