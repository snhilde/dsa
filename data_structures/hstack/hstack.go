// Package hstack provides a simple and lean last-in/first-out stack.
package hstack

import (
	"github.com/snhilde/dsa/data_structures/hlist"
	"errors"
)


// Stack is the main type for this package. It holds the internal information about the stack.
type Stack struct {
	hlist.List
}

// Create a new stack.
func New() *Stack {
	return new(Stack)
}


// Add a new node to the top of the stack.
func (s *Stack) Add(v interface{}) error {
	if s == nil {
		return errors.New("Must create stack with New() first")
	}

	return s.Insert(v, 0)
}

// Pop the top item from the stack.
func (s *Stack) Pop() interface{} {
	if s == nil {
		return nil
	}

	return s.Remove(0)
}

// Get the current number of items in the stack.
func (s *Stack) Count() int {
	if s == nil {
		return -1
	}

	return s.Length()
}

// Reset the stack to a new state.
func (s *Stack) Clear() error {
	if s == nil {
		return errors.New("Stack does not exist")
	}

	return s.List.Clear()
}

// Add a stack on top of the current stack, preserving order. This will clear the new stack.
func (s *Stack) Stack(ns *Stack) error {
	// For efficiency, we're going to link the new stack in on top of the current stack.
	if s == nil {
		return errors.New("Current stack does not exist")
	} else if ns == nil || ns.Count() == 0 {
		// Nothing to add.
		return nil
	}

	if err := ns.Merge(&s.List); err != nil {
		return err
	}

	s.List = ns.List
	return ns.Clear()
}

// Display stack contents, from the top to the bottom.
func (s *Stack) String() string {
	if s == nil {
		return "<nil>"
	}

	return s.List.String()
}
