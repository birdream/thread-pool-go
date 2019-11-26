package main

import "os"

var (
	MaxWorker = os.Getenv("MAX_WORKERS")
	MaxQueue  = os.Getenv("MAX_QUEUE")
)

// Job every job to run
type Job struct {
	Payload interface{}
}

// JobQueue a buffer channel that we can send requests on
var JobQueue chan Job

// Worker execute the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

// NewWorker get a worker
func NewWorker(wp chan chan Job) Worker {
	return Worker{
		WorkerPool: wp,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// Start start the worker to consume the job
func (w Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.JobChannel:
				job = job
				// we receive a job here
				// do something about the job
			case <-w.quit:
				// some one call stop the worker
				// we will stop the for loop and exist the go runtine
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
