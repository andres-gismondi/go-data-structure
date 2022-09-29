package stack_test

import (
	"testing"

	"github.com/andres-gismondi/go-data-structure.git/stack"
)

func TestStack_Int(t *testing.T) {
	s := stack.Stack[int]{}

	s.Pop()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Print()
	s.Pop()
	s.Pop()
	s.Print()
}

func TestStack_String(t *testing.T) {
	s := stack.Stack[string]{}

	s.Pop()
	s.Push("a")
	s.Push("b")
	s.Push("c")
	s.Push("d")
	s.Push("e")
	s.Print()
	s.Pop()
	s.Pop()
	s.Print()
}

type StackStruct struct {
	Name string
}

func TestStack_Struct(t *testing.T) {
	s := stack.Stack[StackStruct]{}

	s.Push(StackStruct{Name: "s1"})
	s.Push(StackStruct{Name: "s2"})
	s.Push(StackStruct{Name: "s3"})
	s.Push(StackStruct{Name: "s4"})
	s.Push(StackStruct{Name: "s5"})
	s.Print()
	s.Pop()
	s.Pop()
	s.Print()
}
