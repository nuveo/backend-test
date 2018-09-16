package main

import (
	"testing"
)

var w ItemQueue

func initQueue() *ItemQueue {
	if w.items == nil {
		w = ItemQueue{}
		w.New()
	}
	return &w
}

func TestEnqueue(t *testing.T) {
	w := initQueue()
	w.Enqueue(Workflow{"1", "inserted", `{"teste":"teste1"}`, `{"teste1"}`})
	w.Enqueue(Workflow{"2", "inserted", `{"teste":"teste2"}`, `{"teste2"}`})
	w.Enqueue(Workflow{"3", "inserted", `{"teste":"teste3"}`, `{"teste3"}`})

	if size := w.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestDequeue(t *testing.T) {
	w.New()

	w.Enqueue(Workflow{"1", "inserted", `{"teste":"teste1"}`, `{"teste1"}`})
	w.Enqueue(Workflow{"2", "inserted", `{"teste":"teste2"}`, `{"teste2"}`})
	w.Enqueue(Workflow{"3", "inserted", `{"teste":"teste3"}`, `{"teste3"}`})

	w.Dequeue()

	if size := w.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}

	w.Dequeue()
	w.Dequeue()
	if size := len(w.items); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}

	if !w.IsEmpty() {
		t.Errorf("IsEmpty should return true")
	}
}
