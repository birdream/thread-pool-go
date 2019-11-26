package main

// Dispatcher dispatch the job to worker
type Dispatcher struct {
	WorkerPool chan chan Job
	maxWorkers int
}

// NewDispatcher ..
func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)

	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

// Run start a dispatcher
func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// received a job outside
			go func(j Job) {
				// try to get a job channel from worker pool
				// will block until it succeed to get a job channel from pool
				jobChan := <-d.WorkerPool

				// send job to job channel
				jobChan <- j
			}(job)
		}
	}
}
