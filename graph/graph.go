package graph

import "go-data-structure/node"

type Graph[K any] struct {
	Nodes []*node.Node[K]
}

func New[K any]() *Graph[K] {
	return &Graph[K]{Nodes: make([]*node.Node[K], 0)}
}

func (g *Graph[K]) AddNode(value K) {
	nde := node.New[K](value, 0)
	g.Nodes = append(g.Nodes, nde)
}

func (g *Graph[K]) Connect(node *node.Node[K], f func(nn, nValue *node.Node[K])) {
	for _, n := range g.Nodes {
		f(n, node)
	}
}

func (g *Graph[K]) AddAndConnect(value K, f func(nn, nValue *node.Node[K])) {
	nde := node.New[K](value, 0)

	g.Nodes = append(g.Nodes, nde)
	if len(g.Nodes) == 1 {
		return
	}

	for _, n := range g.Nodes {
		f(n, nde)
	}
}

func (g *Graph[K]) GetNodeBy(f func(n *node.Node[K]) bool) *node.Node[K] {
	for _, nde := range g.Nodes {
		if f(nde) {
			return nde
		}
	}
	return nil
}

func (g *Graph[K]) RemoveNode(n *node.Node[K]) {
	var index int
	var exist bool

	for i, nde := range g.Nodes {
		if nde == n {
			index = i
			exist = true
			break
		}
	}

	if exist {
		// first index
		if index == 0 {
			g.Nodes = g.Nodes[1:]
			return
		}
		// last index
		if len(g.Nodes)-1 == index {
			g.Nodes = g.Nodes[:index-1]
			return
		}

		first := g.Nodes[:index-1]

		// when index is 1 to avoid use this element we have to start from 0
		if index == 1 {
			first = g.Nodes[:1]
		}
		second := g.Nodes[index+1:]

		// new slice to accomplish new structure
		var modified []*node.Node[K]
		modified = append(modified, first...)
		modified = append(modified, second...)

		g.Nodes = modified
	}
}

func (g *Graph[K]) RemoveNodes() {
	g.Nodes = nil
}

type Visits[K any] []*node.Node[K]

func (vts *Visits[K]) Add(node *node.Node[K]) {
	*vts = append(*vts, node)
}
func (vts *Visits[K]) WasVisited(id *node.Node[K]) bool {
	for _, v := range *vts {
		if v == id {
			return true
		}
	}
	return false
}
