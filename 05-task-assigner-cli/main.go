package main

import (
	"fmt"
	"task-assigner-cli/task"
	"task-assigner-cli/worker"
)

func main() {
	tasks := []task.Task{
		{ID: 1, Description: "Mail Gonder"},
		{ID: 2, Description: "Rapor Hazirla"},
		{ID: 3, Description: "Loglari Temizle"},
		{ID: 4, Description: "Testleri Kos"},
	}

	workers := []worker.Worker{
		{Name: "Ahmet"},
		{Name: "Mehmet"},
		{Name: "Ayşe"},
	}

	resultChan := make(chan string)

	for i, t := range tasks {
		w := workers[i%len(workers)]
		w.DoTask(t, resultChan)
	}

	for range tasks {
		fmt.Println(<-resultChan)
	}
}
