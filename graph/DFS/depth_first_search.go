package DFS

import (
	"fmt"

	"go-data-structure/graph"
	"go-data-structure/node"
)

func Scan[K any](node *node.Node[K]) {
	visits := graph.Visits[K]{}
	dfs(node, &visits)
}

func dfs[K any](node *node.Node[K], visits *graph.Visits[K]) {
	if visits.WasVisited(node) {
		return
	}

	fmt.Printf("%v ;", node.Val)
	visits.Add(node)
	for _, n := range node.Nodes {
		dfs(n, visits)
	}
}
