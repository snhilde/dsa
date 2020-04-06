// Package htree provides an interface to create and use balanced AVL binary trees.
package htree

import (
	"errors"
)


// Tree is the main type for this package. It holds information about the entire AVL tree.
type Tree struct {
	trunk  *node
	length  int
}

type node struct {
	v  interface{} // value of node
	b  int         // Balance of node: -1 if left branch is longer, 0 if both branches are even, and 1 if right side is longer
	l *node        // left branch
	r *node        // right branch
}


// New creates a new binary tree, optionally being populated with the provided values.
func New(vs ...interface{}) *Tree {
	t := new(Tree)
	t.Add(vs...)

	return t
}


func tErr() error {
	return errors.New("Tree must be created with New() first")
}
