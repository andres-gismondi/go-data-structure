package balanced_test

import (
	"testing"

	"go-data-structure/tree/balanced"
)

func TestBalancedTree(t *testing.T) {
	tr := balanced.New[int]()

	tr.AddNode(5)
	tr.AddNode(2)
	tr.AddNode(6)

	tr.AddNode(8)
	tr.AddNode(10)

	tr.AddNode(1)
	tr.AddNode(2)
	tr.AddNode(4)

	tr.Print()
}
