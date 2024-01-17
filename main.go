package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	resultMap = map[string]int{}
	mutex     sync.Mutex
)

type Task func() int

func main() {
	tasks := map[string]Task{
		"Task 1": func() int {
			return 1 + 1 + 1
		},
		"Task 2": func() int {
			return 2 + 2 + 2
		},
		"Task 3": func() int {
			return 3 + 3 + 3
		},
		"Task 4": func() int {
			return 4 + 4 + 4
		},
		"Task 5": func() int {
			return 5 + 5 + 5
		},
	}

	var wg sync.WaitGroup

	for taskName, taskFunc := range tasks {
		wg.Add(1)
		go getTaskDone(taskName, taskFunc, &wg)
	}

	wg.Wait()

	for tskName, result := range resultMap {
		fmt.Printf("Task Name : %s\nTask Result : %d\n\n", tskName, result)
	}

}

func getTaskDone(taskName string, taskFunc Task, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(500 * time.Millisecond)

	taskresult := taskFunc()
	mutex.Lock()
	resultMap[taskName] = taskresult
	mutex.Unlock()
}
