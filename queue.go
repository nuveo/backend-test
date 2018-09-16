package main

import (
	"sync"
)

// ItemQueue is a queue of workflow Items.
type ItemQueue struct {
	items []Workflow
	lock  sync.RWMutex
}

// New creates a new workflow ItemQueue.
func (w *ItemQueue) New() *ItemQueue {
	w.items = []Workflow{}

	return w
}

// Enqueue adds a workflow Item to the end of the queue.
func (w *ItemQueue) Enqueue(i Workflow) {
	w.lock.Lock()
	w.items = append(w.items, i)
	w.lock.Unlock()
}

// Dequeue removes a workflow Item from the first position of the queue.
func (w *ItemQueue) Dequeue() Workflow {
	w.lock.Lock()
	workflow := w.items[0]
	w.items = w.items[1:len(w.items)]
	w.lock.Unlock()

	return workflow
}

// Remove removes the selected workflow Item from the the queue.
func (w *ItemQueue) Remove(id int) {
	w.lock.Lock()
	for i, v := range w.items {
		if v.UUID == id {
			w.items = append(w.items[:i], w.items[i+1:]...)
		}
	}
	w.lock.Unlock()

	return
}

// Front returns the first Item in the queue without removing it.
func (w *ItemQueue) Front() *Workflow {
	w.lock.RLock()
	workflow := w.items[0]
	w.lock.RUnlock()

	return &workflow
}

// IsEmpty verifies if the queue is empty.
func (w *ItemQueue) IsEmpty() bool {
	return len(w.items) == 0
}

// Size returns the number of Items in the queue.
func (w *ItemQueue) Size() int {
	return len(w.items)
}
