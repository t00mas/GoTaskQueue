package gotaskqueue

import "time"

type Queue struct {
	tasks  []Task
	ticker *time.Ticker
	quit   chan struct{}
}

func (q *Queue) Enqueue(t Task) {
	q.tasks = append(q.tasks, t)
}

func (q *Queue) Dequeue() Task {
	task := q.tasks[0]
	q.tasks = q.tasks[1:]
	return task
}

func (q *Queue) IsEmpty() bool {
	return len(q.tasks) == 0
}

func (q *Queue) Size() int {
	return len(q.tasks)
}

func (q *Queue) Peek() Task {
	return q.tasks[0]
}

func (q *Queue) Start() {
	q.ticker = time.NewTicker(500 * time.Millisecond)
	q.quit = make(chan struct{})
	go func() {
		for {
			select {
			case <-q.ticker.C:
				if !q.IsEmpty() {
					task := q.Dequeue()
					task()
				}
			case <-q.quit:
				q.ticker.Stop()
				return
			}
		}
	}()
}

func (q *Queue) Stop() {
	close(q.quit)
}
