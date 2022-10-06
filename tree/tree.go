package tree

type Tree[T any] interface {
	AddNode(val T)
	Print()
}
