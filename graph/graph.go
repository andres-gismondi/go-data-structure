package graph

type T any
type Node[T any] struct {
	Val   T
	Nodes []*Node[T]
}

func NewNode(val T) *Node[T] {
	return &Node[T]{
		Val:   val,
		Nodes: make([]*Node[T], 0),
	}
}

func (n *Node[T]) AddNode(node *Node[T]) {
	n.Nodes = append(n.Nodes, node)
}

type Visits []*Node[T]

func (vts *Visits) Add(node *Node[T]) {
	*vts = append(*vts, node)
}
func (vts *Visits) WasVisited(id T) bool {
	for _, v := range *vts {
		if v == id {
			return true
		}
	}
	return false
}
