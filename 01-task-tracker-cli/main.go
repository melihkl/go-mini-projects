package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task/task"
)


func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Başlık: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Açıklama: ")
	description, _ := reader.ReadString('\n')

	fmt.Print("Durum: ")
	status, _ := reader.ReadString('\n')

	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)
	status = strings.TrimSpace(status)

	// Add the task
	task.AddTask(title, description, status)

	// List all tasks
	task.ListTask()
}

