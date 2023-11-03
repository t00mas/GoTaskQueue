# GoTaskQueue

GoTaskQueue is a simple, concurrent task queue implemented in Go. It uses a ticker to periodically check the queue and run tasks if available.

## Features

- Enqueue tasks
- Dequeue tasks
- Check if the queue is empty
- Get the size of the queue
- Peek at the front of the queue
- Start and stop the queue

## Usage

First, import the package:

```go
import "github.com/t00mas/GoTaskQueue"
```

Next, initialize and start the queue:

```go
q := &GoTaskQueue.Queue{} q.Start()
```

Define a task:

```go
task := func() { // Your task code here }
```

Enqueue the task:

```go
q.Enqueue(task)
```

You can also stop the queue when you're done with it:

```go
q.Stop()
```

## Testing

To run the tests, use the `go test` command:

```bash
go test
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
