package queue

import "fmt"

type Queue[T any] []T

func (q *Queue[T]) Enqueue(val T) {
	*q = append(*q, val)
}

func (q *Queue[T]) DequeueHead() T {
	var res T
	if len(*q) == 0 {
		return res
	}

	head := (*q)[0]
	*q = (*q)[1:]

	return head
}

func (q *Queue[T]) DequeueTail() T {
	var res T
	if len(*q) == 0 {
		return res
	}

	last := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]

	return last
}

func (q *Queue[_]) Empty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

func (q *Queue[_]) Print() {
	for _, item := range *q {
		fmt.Printf("%v ;", item)
	}
	fmt.Println()
}
