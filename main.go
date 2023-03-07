package main

import (
	"fmt"
	"time"
)

type Task struct {
	Name     string
	Function func()
	Interval time.Duration
}

func RunScheduler(tasks []Task) {
	// Create a ticker that triggers every second
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// Initialize a map to keep track of the last execution time of each task
	lastExecution := make(map[string]time.Time)

	for {
		select {
		case <-ticker.C:
			// Check if any tasks are due for execution
			for _, task := range tasks {
				if time.Now().Sub(lastExecution[task.Name]) >= task.Interval {
					// Execute the task
					task.Function()

					// Update the last execution time for the task
					lastExecution[task.Name] = time.Now()
				}
			}
		}
	}
}
func Task1() {
	fmt.Println("Task 1 executed")
}

func Task2() {
	fmt.Println("Task 2 executed")
}
func main() {
	tasks := []Task{
		{Name: "Task 1", Function: Task1, Interval: 5 * time.Second},
		{Name: "Task 2", Function: Task2, Interval: 10 * time.Second},
	}

	go RunScheduler(tasks)

	// Keep the main program running indefinitely
	select {}
}
