package graph_test

import (
	"testing"

	"go-data-structure/graph"
	"go-data-structure/node"
)

type GT struct {
	OurTag   string
	TheirTag string
}

func TestGraph_removeNode(t *testing.T) {
	g := graph.Graph[GT]{Nodes: make([]*node.Node[GT], 0)}

	f := func(n1, n2 *node.Node[GT]) {
		if n1.Val.OurTag == n2.Val.OurTag {
			n1.AddNode(n2)
			n2.AddNode(n1)
		}
		if n1.Val.TheirTag == n2.Val.TheirTag {
			n1.AddNode(n2)
			n2.AddNode(n1)
		}
	}

	a := GT{OurTag: "x", TheirTag: "0"}
	b := GT{OurTag: "x", TheirTag: "1"}
	c := GT{OurTag: "y", TheirTag: "1"}
	d := GT{OurTag: "z", TheirTag: "2"}

	g.AddAndConnect(a, f)
	g.AddAndConnect(b, f)
	g.AddAndConnect(c, f)
	g.AddAndConnect(d, f)

}
