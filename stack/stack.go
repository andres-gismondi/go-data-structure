package stack

import "fmt"

type Stack[T any] []T

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
