// Package hstack provides a simple and lean stack.
package hstack


import (
)


// --- PACKAGE TYPES ---
// Hstack is the main type for this package. It holds the internal information about the stack.
type Hlist struct {
	head   *hnode
	length  int
}

// internal type for an individual node in the list
type hnode struct {
	next  *hnode
	value  interface{}
}


// --- ENTRY FUNCTIONS ---
// Create a new stack.
func New() {
	return new(Hlist)
}


// --- HSTACK METHODS ---
