package helpers

import (
	"fmt"
	"sync"
)

// ElementQueue defines the type of the queue
type ElementQueue struct {
	elements []string
	lock     sync.RWMutex
}

// Create allows create a new queue.
func (e *ElementQueue) Create() *ElementQueue {
	e.elements = []string{}

	return e
}

// IsEmpty returns if queue is empty
func (e *ElementQueue) IsEmpty() bool {
	return len(e.elements) == 0
}

// IsFull returns if queue is full
func (e *ElementQueue) IsFull() bool {
	return len(e.elements) == e.Size()
}

// Size returns the queue size
func (e *ElementQueue) Size() int {
	return len(e.elements)
}

// Enqueue add an item to the queue
func (e *ElementQueue) Enqueue(item string) {
	e.lock.Lock()

	e.elements = append(e.elements, item)
	fmt.Printf("Enqueued: %v", item)

	e.lock.Unlock()

}

// Dequeue remove an item to the first position queue
func (e *ElementQueue) Dequeue() string {
	e.lock.Lock()
	item := e.elements[0]

	e.elements = e.elements[1:len(e.elements)]
	e.lock.Unlock()

	return item

}

// First returns a first item of the queue
func (e *ElementQueue) First() string {
	e.lock.Lock()
	item := e.elements[0]
	e.lock.Unlock()

	return item

}
