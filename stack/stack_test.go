package stack_test

import (
	"testing"

	"go-data-structure/stack"
)

func TestStack_Int(t *testing.T) {
	s := stack.New[int]()

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
	s := stack.New[string]()

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

type ObjectTest struct {
	Name string
}

func TestStack_Struct(t *testing.T) {
	s := stack.New[ObjectTest]()

	s.Push(ObjectTest{Name: "s1"})
	s.Push(ObjectTest{Name: "s2"})
	s.Push(ObjectTest{Name: "s3"})
	s.Push(ObjectTest{Name: "s4"})
	s.Push(ObjectTest{Name: "s5"})
	s.Print()
	s.Pop()
	s.Pop()
	s.Print()
}

func TestStruct(t *testing.T) {
	s := stack.New[int]()

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
