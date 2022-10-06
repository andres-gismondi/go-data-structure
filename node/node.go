package node

type Node[T any] struct {
	Val   T
	Nodes []*Node[T]
}

func New[T any](val T, cap int) *Node[T] {
	return &Node[T]{
		Val:   val,
		Nodes: make([]*Node[T], cap, cap),
	}
}

func (n *Node[T]) AddNode(node *Node[T]) {
	n.Nodes = append(n.Nodes, node)
}
