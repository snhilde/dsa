// Package hstack provides a simple and lean stack.
package hstack

import (
	"errors"
	"strings"
	"fmt"
)


// --- PACKAGE TYPES ---
// Stack is the main type for this package. It holds the internal information about the stack.
type Stack struct {
	head   *hnode
	length  int
}

// internal type for an individual node in the stack
type hnode struct {
	next  *hnode
	value  interface{}
}


// --- ENTRY FUNCTIONS ---
// Create a new stack.
func New() *Stack {
	return new(Stack)
}


// --- STACK METHODS ---
// Add a new node to the top of the stack.
func (stack *Stack) Add(value interface{}) error {
	if stack == nil {
		return errors.New("Must create stack with New() first")
	}

	node := newNode(value)
	node.next = stack.head
	stack.head = node
	stack.length++

	return nil
}

// Pop the top item from the stack.
func (stack *Stack) Pop() interface{} {
	if stack == nil || stack.head == nil {
		return nil
	}

	pop := stack.head
	stack.head = pop.next
	stack.length--

	return pop.value
}

// Get the current number of items in the stack.
func (stack *Stack) Count() int {
	if stack == nil {
		return -1
	}

	return stack.length
}

// Reset the stack to a new state.
func (stack *Stack) Clear() error {
	if stack == nil {
		return errors.New("Stack does not exist")
	}

	stack.head = nil
	stack.length = 0

	return nil
}

// Add a stack on top of the current stack, preserving order. This will clear the new stack.
func (stack *Stack) Merge(new_stack *Stack) error {
	// For efficiency, we're going to link the new stack in on top of the current stack.
	if stack == nil {
		return errors.New("Current stack does not exist")
	} else if new_stack == nil || new_stack.Count() == 0 {
		// Nothing to add.
		return nil
	}

	// Find the bottom of the new stack and put it on top of the current stack.
	if stack.head == nil {
		stack.head = new_stack.head
		stack.length = new_stack.length
	} else {
		node := new_stack.head
		for node.next != nil {
			node = node.next
		}
		node.next = stack.head
		stack.head = new_stack.head
		stack.length += new_stack.length
	}

	return new_stack.Clear()
}

// Display stack contents, with left being the top.
func (stack *Stack) String() string {
	var b strings.Builder

	if stack == nil {
		return "<nil>"
	} else if stack.Count() == 0 {
		return "<empty>"
	}

	node := stack.head
	b.WriteString(fmt.Sprintf("%v", node.value))
	node = node.next
	for node != nil {
		b.WriteString(fmt.Sprintf(", %v", node.value))
		node = node.next
	}

	return b.String()
}


// --- HELPER FUNCTIONS ---
// internal convenience function for creating a new node
func newNode(value interface{}) *hnode {
	node := new(hnode)
	node.value = value

	return node
}
