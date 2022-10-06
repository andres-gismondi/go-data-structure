package balanced

import (
	"fmt"

	"go-data-structure/node"
	t "go-data-structure/tree"
)

type Item interface {
	int | string
}

type tree[T Item] struct {
	Root *node.Node[T]
}

func New[T Item]() t.Tree[T] {
	return &tree[T]{}
}

func (t *tree[T]) AddNode(val T) {
	n := node.New[T](val, 2)
	if t.Root == nil {
		t.Root = n
		return
	}

	t.add(t.Root, n)
}

func (t *tree[T]) add(node, appendedNode *node.Node[T]) {
	// value is uppermore than actual node then we go to right
	if appendedNode.Val > node.Val {
		if node.Nodes[1] == nil {
			node.Nodes[1] = appendedNode
			return
		}
		t.add(node.Nodes[1], appendedNode)
	}

	// value is lower than actual node then we go to left
	if appendedNode.Val < node.Val {
		if node.Nodes[0] == nil {
			node.Nodes[0] = appendedNode
			return
		}
		t.add(node.Nodes[0], appendedNode)
	}
}

func (t *tree[_]) Print() {
	t.print(t.Root)
}

func (t *tree[T]) print(node *node.Node[T]) {
	if node == nil {
		return
	}
	fmt.Printf("%v ;", node.Val)
	for _, n := range node.Nodes {

		t.print(n)
	}
}
