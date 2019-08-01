package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// ThreadSafeQueue is something else
type ThreadSafeQueue struct {
	queue   []interface{}
	maxSize int
	mu      *sync.Mutex
	cond    *sync.Cond
}

// NewThreadSafeQueue build a new ThreadSafeQueue
func NewThreadSafeQueue(maxSize int) *ThreadSafeQueue {
	var queue []interface{}
	mu := sync.Mutex{}
	// mu1 := sync.Mutex{}

	pool := &ThreadSafeQueue{
		queue:   queue,
		maxSize: maxSize,
		mu:      &mu,
		cond:    sync.NewCond(&mu),
	}

	return pool
}

// Size return the queue size
func (p *ThreadSafeQueue) Size() int {
	p.mu.Lock()
	defer p.mu.Unlock()

	return len(p.queue)
}

// Put put one ele into queue
func (p *ThreadSafeQueue) Put(item interface{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.maxSize != 0 && p.Size() >= p.maxSize {
		return errors.New("Over size")
	}

	p.queue = append(p.queue, item)

	// notify someone who might waiting
	// info one go rutinn
	// p.cond.Signal()
	// info all go rutinn
	p.cond.Broadcast()

	return nil
}

// BatchPut put many ele into queue
func (p *ThreadSafeQueue) BatchPut(items []interface{}) error {

	for _, item := range items {
		if err := p.Put(item); err != nil {
			return err
		}
	}

	return nil
}

// Pop ssss
func (p *ThreadSafeQueue) Pop(block bool) interface{} {
	if p.Size() == 0 {
		p.mu.Lock()
		defer p.mu.Unlock()

		if block {
			fmt.Print("waiting....\n")
			p.cond.Wait()
		} else {
			return nil
		}
		// return nil
	}

	var item interface{}

	if len(p.queue) > 0 {
		// pop
		item, p.queue = p.queue[len(p.queue)-1], p.queue[:len(p.queue)-1]
	}

	return item
}

// Get sss
func (p *ThreadSafeQueue) Get(idx int) interface{} {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.queue[idx]
}

func producer(queue *ThreadSafeQueue) {
	i := 0
	for {
		i++
		queue.Put(i)
		fmt.Printf("+++put item from q: %d \n", i)
		time.Sleep(time.Second * 3)
	}
}

func consumer(queue *ThreadSafeQueue) {
	i := 0
	for {
		i++
		item := queue.Pop(true)
		fmt.Printf("---get item from q: %d , who am I : %d\n", item, i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	queue := NewThreadSafeQueue(0)

	go producer(queue)
	go consumer(queue)

	time.Sleep(time.Second * 5000)
}
