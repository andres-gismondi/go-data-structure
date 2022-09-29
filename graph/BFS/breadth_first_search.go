package BFS

import (
	"fmt"

	"github.com/andres-gismondi/go-data-structure.git/graph"
	"github.com/andres-gismondi/go-data-structure.git/queue"
)

func Scan(node *graph.Node[graph.T]) {
	visits := graph.Visits{}
	q := queue.Queue[*graph.Node[graph.T]]{}
	q.Enqueue(node)

	for !q.Empty() {
		node = q.DequeueHead()
		fmt.Printf("%v ;", node.Val)
		visits.Add(node)

		for _, n := range node.Nodes {
			if !visits.WasVisited(n) {
				q.Enqueue(n)
			}
		}
	}
}
