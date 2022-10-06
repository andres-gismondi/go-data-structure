package DFS_test

import (
	"testing"

	"go-data-structure/graph"
	"go-data-structure/graph/DFS"
	"go-data-structure/node"
)

type User struct {
	Name string
	Age  int
}

func TestDFS_success(t *testing.T) {
	andres := User{Name: "andres", Age: 34}
	julian := User{Name: "julian", Age: 34}
	camila := User{Name: "camila", Age: 31}
	florencia := User{Name: "florencia", Age: 65}
	julieta := User{Name: "julieta", Age: 42}

	g := graph.New[User]()
	g.AddAndConnect(julian, func(n, n1 *node.Node[User]) {})
	g.AddAndConnect(camila, func(n, n1 *node.Node[User]) {})
	g.AddAndConnect(florencia, func(n, n1 *node.Node[User]) {})

	g.AddAndConnect(julieta, func(node, julieta *node.Node[User]) {
		if node.Val.Name == "andres" {
			node.AddNode(julieta)
			julieta.AddNode(node)
		}
		if node.Val.Name == "florencia" {
			julieta.AddNode(node)
		}
	})

	g.AddAndConnect(andres, func(node, andres *node.Node[User]) {
		if node.Val.Name == "florencia" || node.Val.Name == "camila" {
			andres.AddNode(node)
		}
	})

	nodeAndres := g.GetNodeBy(func(n *node.Node[User]) bool {
		if n.Val.Name == "andres" {
			return true
		}
		return false
	})

	DFS.Scan(nodeAndres)
}

/*
func TestDFS_circle(t *testing.T) {
	a := graph.New("a")
	b := graph.New("b")
	c := graph.New("c")
	d := graph.New("d")

	c.AddAndConnect(c)

	a.AddAndConnect(c)
	a.AddAndConnect(b)

	b.AddAndConnect(c)
	b.AddAndConnect(d)

	DFS.Scan(a)
}*/
