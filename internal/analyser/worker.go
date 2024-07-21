package analyser

// Task is the type for a function provided
// to the worker pool
type Task func() []string

// worker is a function that can be asynchronously scheduled to read
// work from an upstream channel to process
func worker(id int, upstream <-chan Task, downstream chan []string) {
	for task := range upstream {
		downstream <- task()
	}
	close(downstream)
}
