// Package hstack provides a simple and lean last-in/first-out stack.
package hstack

import (
	"fmt"
	"github.com/snhilde/dsa/data_structures/hlist"
)

var (
	// This is the standard error message when trying to use an invalid stack.
	errBadStack = fmt.Errorf("must create stack with New() first")
)

// Stack is the main type for this package. It holds the internal information about the stack.
type Stack struct {
	list *hlist.List
}

// New creates a new stack.
func New() *Stack {
	s := new(Stack)
	s.list = hlist.New()
	return s
}

// Add adds one or more new items to the top of the stack. If there is more than one item, the first item will be at the
// top.
func (s *Stack) Add(vs ...interface{}) error {
	if s == nil {
		return errBadStack
	}

	// If caller is trying to add own stack, duplicate it first and then add it.
	for i, v := range vs {
		if t, ok := v.(*Stack); ok {
			if t == s {
				ns, err := s.Copy()
				if err != nil {
					return err
				}
				vs[i] = ns
			}
		}
	}

	return s.list.Insert(0, vs...)
}

// Pop removes the top item from the stack and returns its value.
func (s *Stack) Pop() interface{} {
	if s == nil {
		return nil
	}

	return s.list.Remove(0)
}

// Count gets the current number of items in the stack.
func (s *Stack) Count() int {
	if s == nil {
		return -1
	}

	return s.list.Length()
}

// Copy makes an exact copy of the stack.
func (s *Stack) Copy() (*Stack, error) {
	if s == nil {
		return nil, errBadStack
	}

	nl, err := s.list.Copy()
	if err != nil {
		return nil, err
	}

	ns := New()
	ns.list = nl

	return ns, nil
}

// Merge adds a stack below the current stack. This will take ownership of and clear the provided stack.
func (s *Stack) Merge(ns *Stack) error {
	if s == nil {
		return errBadStack
	} else if ns == nil || ns.Count() == 0 {
		// Nothing to add.
		return nil
	}

	if err := s.list.Merge(ns.list); err != nil {
		return err
	}

	return ns.Clear()
}

// Clear resets the stack to its initial state.
func (s *Stack) Clear() error {
	if s == nil {
		return fmt.Errorf("stack does not exist")
	}

	return s.list.Clear()
}

// String displays the stack's contents, from the top to the bottom.
func (s *Stack) String() string {
	if s == nil {
		return "<nil>"
	}

	return s.list.String()
}
