package circles

import (
	"go-data-structure/graph"
	"go-data-structure/node"
)

func Directed[K any](n *node.Node[K]) bool {
	visited := graph.Visits[K]{}
	return circle[K](n, visited)
}

// visits parameter must be a pointer in order to have
// all elements inside when the function pop from the stack in recursion
func circle[K any](node *node.Node[K], visited graph.Visits[K]) bool {
	// if this node was visited means that a circle exist
	if visited.WasVisited(node) {
		return true
	}

	// add node to visits nodes
	visited.Add(node)

	// no more nodes to iterate
	if node.Nodes == nil {
		return false
	}

	// iterate over each adjacent node recursively
	for _, n := range node.Nodes {
		if res := circle(n, visited); res {
			return true
		}
	}

	return false
}
