package gotaskqueue

import (
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	// Initialize the queue
	q := &Queue{}

	// Start the queue
	q.Start()

	// Create a channel to receive the result of the task
	done := make(chan bool)

	// Define a task
	task := func() {
		time.Sleep(10 * time.Millisecond)
		done <- true
	}

	// Enqueue the task
	q.Enqueue(task)

	// Wait for the task to complete
	result := <-done

	// Stop the queue
	q.Stop()

	if !result {
		t.Errorf("Task did not complete successfully")
	}
}
