package queue_test

import (
	"testing"

	"github.com/andres-gismondi/go-data-structure.git/queue"
)

func TestQueue_Int(t *testing.T) {
	q := queue.Queue[int]{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Print()

	q.DequeueHead()
	q.Print()

	q.DequeueTail()
	q.Print()
}

func TestQueue_String(t *testing.T) {
	q := queue.Queue[string]{}

	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")
	q.Enqueue("d")
	q.Enqueue("e")
	q.Print()

	q.DequeueHead()
	q.Print()

	q.DequeueTail()
	q.Print()
}

type QueueStruct struct {
	Name string
}

func TestQueue_Struct(t *testing.T) {
	q := queue.Queue[QueueStruct]{}

	q.Enqueue(QueueStruct{Name: "Q1"})
	q.Enqueue(QueueStruct{Name: "Q2"})
	q.Enqueue(QueueStruct{Name: "Q3"})
	q.Enqueue(QueueStruct{Name: "Q4"})
	q.Enqueue(QueueStruct{Name: "Q5"})
	q.Print()

	q.DequeueHead()
	q.Print()

	q.DequeueTail()
	q.Print()
}
