package DFS_test

import (
	"testing"

	"github.com/andres-gismondi/go-data-structure.git/graph"
	"github.com/andres-gismondi/go-data-structure.git/graph/DFS"
)

type User struct {
	Name string
	Age  int
}

func TestDFS_success(t *testing.T) {
	andres := graph.NewNode(User{Name: "andres", Age: 34})
	julian := graph.NewNode(User{Name: "julian", Age: 34})
	camila := graph.NewNode(User{Name: "camila", Age: 31})
	florencia := graph.NewNode(User{Name: "florencia", Age: 65})
	julieta := graph.NewNode(User{Name: "julieta", Age: 42})

	julieta.AddNode(julian)
	julieta.AddNode(florencia)

	andres.AddNode(julieta)
	andres.AddNode(florencia)
	andres.AddNode(camila)

	DFS.Scan(andres)
}
