package datastructure

import (
	"testing"
)

type intKey int

// satisfy avl.Key
func (k intKey) Less(k2 AVLKey) bool { return k < k2.(intKey) }
func (k intKey) Eq(k2 AVLKey) bool   { return k == k2.(intKey) }

func TestInsert(t *testing.T) {
	var tree *AVLNode
	AVLInsert(&tree, intKey(5))
	AVLInsert(&tree, intKey(4))
	AVLInsert(&tree, intKey(9))
	AVLInsert(&tree, intKey(100))

	if tree.Data != intKey(5) {
		t.Error("Misplaced!")
	}
	if tree.Link[0].Data != intKey(4) {
		t.Error("Misplaced!")
	}
	if tree.Link[1].Data != intKey(9) {
		t.Error("Misplaced!")
	}
	if tree.Link[1].Link[1].Data != intKey(100) {
		t.Error("Misplaced!")
	}
}
