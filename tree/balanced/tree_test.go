package balanced_test

import (
	"testing"

	"github.com/andres-gismondi/go-data-structure.git/tree/balanced"
)

func TestBalancedTree(t *testing.T) {
	tres := balanced.Tree{}
	tres.AddNode(5)
	tres.AddNode(2)
	tres.AddNode(6)
	tres.Print()
}
