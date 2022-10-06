package BFS_test

import (
	"testing"

	"go-data-structure/graph"
	"go-data-structure/graph/BFS"
	"go-data-structure/node"
)

type User struct {
	Name string
	Age  int
}

func TestBFS_success(t *testing.T) {
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
		if node.Val.Name == "julian" {
			julieta.AddNode(node)
		}
		if node.Val.Name == "florencia" {
			julieta.AddNode(node)
		}
	})

	g.AddAndConnect(andres, func(node, andres *node.Node[User]) {
		if node.Val.Name == "florencia" || node.Val.Name == "camila" || node.Val.Name == "julieta" {
			andres.AddNode(node)
		}
	})

	nodeAndres := g.GetNodeBy(func(n *node.Node[User]) bool {
		if n.Val.Name == "andres" {
			return true
		}
		return false
	})

	BFS.Scan(nodeAndres)
}
