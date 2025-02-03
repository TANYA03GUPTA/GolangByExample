package main

import (
	"context"
	"fmt"
	"time"
)

type Task struct {
	ID      int
	Message string
}
type Worker interface {
	Process(ctx context.Context, task Task) string
}
type SimpleWorker struct{}

func (sw SimpleWorker) Process(ctx context.Context, task Task) string {
	
	select {
	case <-time.After(2 * time.Second): // Simulate a task that takes 2 seconds
		return fmt.Sprintf("Simple Worker processed task %d: %s", task.ID, task.Message)
	case <-ctx.Done():
		return fmt.Sprintf("Simple Worker task %d cancelled", task.ID)
	}
}

type ComplexWorker struct{}

func (cw ComplexWorker) Process(ctx context.Context, task Task) string {
	
	select {
	case <-time.After(5 * time.Second): // Simulate a task that takes 5 seconds
		return fmt.Sprintf("Complex Worker processed task %d: %s", task.ID, task.Message)
	case <-ctx.Done():
		return fmt.Sprintf("Complex Worker task %d cancelled", task.ID)
	}
}

func handleTasks(ctx context.Context, tasks []Task, worker Worker, ch chan string) {
	for _, task := range tasks {
		// Process each task concurrently using the worker's Process method
		go func(task Task) {
			result := worker.Process(ctx, task)
			ch <- result
		}(task)
	}
}

func main() {
	// Initialize tasks
	tasks := []Task{
		{ID: 1, Message: "Clean the house"},
		{ID: 2, Message: "Fix the car"},
		{ID: 3, Message: "Prepare dinner"},
	}

	// collect results from workers
	resultChannel := make(chan string, len(tasks))

	//  cancellation after a time
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // Timeout after 3 seconds
	defer cancel()

	worker := SimpleWorker{} // Change this to ComplexWorker to test with the other worker
	handleTasks(ctx, tasks, worker, resultChannel)
 
	for i := 0; i < len(tasks); i++ {
		fmt.Println(<-resultChannel)
	}

	fmt.Println("All tasks processed or cancelled.")
}
