package worker

import (
	"fmt"
	"task-assigner-cli/task"
	"time"
)

type Worker struct {
	Name string
}

func (w Worker) DoTask(t task.Task, resultChan chan<- string) {
	go func() {
		time.Sleep(time.Duration(1+t.ID%3) * time.Second)
		result := fmt.Sprintf("Görev %d (%s) tamamlandi -> %s", t.ID, t.Description, w.Name)
		resultChan <- result
	}()
}
