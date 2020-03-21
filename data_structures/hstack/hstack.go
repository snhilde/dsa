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
func (stack *Stack) Add(value interface{}) error {
	if stack == nil {
		return errors.New("Must create stack with New() first")
	}

	return stack.Insert(value, 0)
}

// Pop the top item from the stack.
func (stack *Stack) Pop() interface{} {
	if stack == nil {
		return nil
	}

	return stack.List.Pop(0)
}

// Get the current number of items in the stack.
func (stack *Stack) Count() int {
	if stack == nil {
		return -1
	}

	return stack.Length()
}

// Reset the stack to a new state.
func (stack *Stack) Clear() error {
	if stack == nil {
		return errors.New("Stack does not exist")
	}

	return stack.List.Clear()
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

	if err := new_stack.List.Merge(&stack.List); err != nil {
		return err
	}

	stack.List = new_stack.List
	return new_stack.Clear()
}

// Display stack contents, with left being the top.
func (stack *Stack) String() string {
	if stack == nil {
		return "<nil>"
	}

	return stack.List.String()
}
