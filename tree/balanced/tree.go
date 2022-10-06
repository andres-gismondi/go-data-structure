package balanced

import "fmt"

// Balanced Tree is binary tree
// We only need a Node
// I create a slice type to manage all the nodes we append
type T any
type Node[T any] struct {
	Value T
	Nodes []*Node[T]
}

func (n *Node[T]) AddNode(node *Node[T]) {
	n.Nodes = append(n.Nodes, node)
}

type Tree struct {
	Root *Node[int]
}

func (t *Tree) AddNode(val int) {
	node := &Node[int]{Value: val}
	if t.Root == nil {
		t.Root = node
		return
	}

	t.add(t.Root, node)
}

func (t *Tree) add(node, appendedNode *Node[int]) {
	// base case when we are in the last node
	if node.Nodes == nil {
		node.AddNode(appendedNode)
		return
	}

	// value is uppermore than actual node we go to right
	if appendedNode.Value > node.Value {
		t.add(node.Nodes[1], appendedNode)
	}

	// value is lower than actual node we go to left
	if appendedNode.Value < node.Value {
		t.add(node.Nodes[0], appendedNode)
	}
}

func (t *Tree) Print() {
	t.print(t.Root)
}

func (t *Tree) print(node *Node[int]) {
	if node.Nodes == nil {
		fmt.Printf("%v ;", node.Value)
		return
	}

	for _, n := range node.Nodes {
		t.print(n)
	}
}
