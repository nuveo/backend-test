package queue

import (
	"fmt"
	"sync"
)

type IdQueue struct {
	items []string
	lock  sync.RWMutex
}

func (q *IdQueue) New() *IdQueue {
	q.items = []string{}
	return q
}

// put a UUID int the end of the queue
func (q *IdQueue) Enqueue(uuid string) {
	q.lock.Lock()
	q.items = append(q.items, uuid) // Simply append to enqueue.
	fmt.Printf("Enqueued: UUID: %v, Queue size %v\n", uuid, q.Size())
	q.lock.Unlock()
}

// remove the UUID of the head of the queue
func (q *IdQueue) Dequeue() {
	q.lock.Lock()
	uuid := q.items[0] // The first element is the one to be dequeued.
	fmt.Printf("Dequeued: UUID: %v, Queue size %v\n", uuid, q.Size())
	q.items = q.items[1:]
	q.lock.Unlock()
}

// verify if the queue is empty
func (q *IdQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// return number of elements in the queue
func (q *IdQueue) Size() int {
	return len(q.items)
}

// return first element of the head of the queue
func (q *IdQueue) First() string {
	q.lock.RLock()
	uuid := q.items[0]
	q.lock.RUnlock()

	return uuid
}
