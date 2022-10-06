package stack

import "fmt"

type Stack[T any] []T

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() T {
	var res T
	if len(*s) == 0 {
		return res
	}

	res = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return res
}

func (s *Stack[_]) Print() {
	for _, item := range *s {
		fmt.Printf("%v ;", item)
	}
	fmt.Println()
}

type Struct[T any] struct {
	Items []T
}

func (s *Struct[T]) Push(val T) {
	s.Items = append(s.Items, val)
}

func (s *Struct[T]) Pop() T {
	var res T
	if len(s.Items) == 0 {
		return res
	}

	res = s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]

	return res
}

func (s *Struct[_]) Empty() bool {
	return len(s.Items) == 0
}

func (s *Struct[T]) Print() {
	for _, item := range s.Items {
		fmt.Printf("%v ;", item)
	}
	fmt.Println()
}
