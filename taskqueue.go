package gotaskqueue

type TaskQueue interface {
	Enqueue(task Task)
	Dequeue() Task
	IsEmpty() bool
	Size() int
	Peek() Task
}
