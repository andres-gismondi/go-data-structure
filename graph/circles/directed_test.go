package circles_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go-data-structure/graph"
	"go-data-structure/graph/circles"
	"go-data-structure/node"
)

func TestDirectedCircles_OneCircle(t *testing.T) {
	g := graph.New[string]()
	g.AddNode("a")
	g.AddNode("b")
	g.AddNode("c")
	g.AddNode("d")

	g.Connect(getNode(g, "a"), func(n, a *node.Node[string]) {
		if n.Val == "b" || n.Val == "c" {
			a.AddNode(n)
		}
	})
	g.Connect(getNode(g, "b"), func(n, b *node.Node[string]) {
		if n.Val == "d" || n.Val == "c" {
			b.AddNode(n)
		}
	})
	g.Connect(getNode(g, "c"), func(n, c *node.Node[string]) {
		if n.Val == "c" {
			c.AddNode(n)
		}
	})
	g.Connect(getNode(g, "d"), func(n, d *node.Node[string]) {})

	a := getNode(g, "a")

	var got bool
	got = circles.Directed(a)
	assert.Equal(t, true, got)
}

func getNode(g *graph.Graph[string], e string) *node.Node[string] {
	return g.GetNodeBy(func(n *node.Node[string]) bool {
		if n.Val == e {
			return true
		}
		return false
	})
}
