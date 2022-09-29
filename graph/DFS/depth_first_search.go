package DFS

import (
	"fmt"

	"github.com/andres-gismondi/go-data-structure.git/graph"
)

func Scan(node *graph.Node[graph.T]) {
	visits := graph.Visits{}
	dfs(node, visits)
}

func dfs(node *graph.Node[graph.T], visits graph.Visits) {
	if visits.WasVisited(node) {
		return
	}

	fmt.Printf("%v ;", node.Val)
	visits.Add(node)
	for _, n := range node.Nodes {
		dfs(n, visits)
	}
}
