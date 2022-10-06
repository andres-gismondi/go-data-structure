package BFS

import (
	"fmt"

	"go-data-structure/graph"
	"go-data-structure/node"
	"go-data-structure/queue"
)

func Scan[K any](n *node.Node[K]) {
	visits := graph.Visits[K]{}

	q := queue.New[*node.Node[K]]()
	q.Enqueue(n)

	for !q.Empty() {
		n = q.DequeueHead()
		fmt.Printf("%v ;", n.Val)
		visits.Add(n)

		for _, n := range n.Nodes {
			if !visits.WasVisited(n) {
				q.Enqueue(n)
			}
		}
	}
}
