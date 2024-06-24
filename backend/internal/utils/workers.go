package utils

type Task func(args ...interface{}) (result interface{}, err error)

type Result struct {
	workerID int
	Result   interface{}
	Err      error
}

type Worker struct {
	id         int
	taskQueue  <-chan TaskWrapper
	resultChan chan<- Result
}

type TaskWrapper struct {
	task Task
	args []interface{}
}

func (w *Worker) Start() {
	go func() {
		for taskWrapper := range w.taskQueue {
			result, err := taskWrapper.task(taskWrapper.args...)
			w.resultChan <- Result{workerID: w.id, Result: result, Err: err}
		}
	}()
}

type WorkerPool struct {
	taskQueue   chan TaskWrapper
	resultChan  chan Result
	workerCount int
}

func NewWorkerPool(workerCount int, queueSize int) *WorkerPool {
	return &WorkerPool{
		taskQueue:   make(chan TaskWrapper, queueSize),
		resultChan:  make(chan Result, queueSize),
		workerCount: workerCount,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		worker := Worker{id: i, taskQueue: wp.taskQueue, resultChan: wp.resultChan}
		worker.Start()
	}
}

func (wp *WorkerPool) Submit(task Task, args ...interface{}) {
	wp.taskQueue <- TaskWrapper{task: task, args: args}
}

func (wp *WorkerPool) GetResult() Result {
	return <-wp.resultChan
}

func (wp *WorkerPool) Close() {
	close(wp.resultChan)
	close(wp.taskQueue)
}
