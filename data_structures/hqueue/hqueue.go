// Package hqueue provides a simple and lean first-in/first-out queue.
package hqueue

import (
	"errors"
	"github.com/snhilde/dsa/data_structures/hlist"
)

// Queue is the main type for this package. It holds the internal information about the queue.
type Queue struct {
	hlist.List
}

// Create a new queue.
func New() *Queue {
	return new(Queue)
}

// Add a new item to the back of the queue.
func (q *Queue) Add(v interface{}) error {
	if q == nil {
		return errors.New("Must create queue with New() first")
	}

	return q.List.Append(v)
}

// Pop the first item in the queue.
func (q *Queue) Pop() interface{} {
	if q == nil {
		return nil
	}

	return q.List.Remove(0)
}

// Get the current number of items in the queue.
func (q *Queue) Count() int {
	if q == nil {
		return -1
	}

	return q.List.Length()
}

// Clear the queue to its inital state.
func (q *Queue) Clear() error {
	if q == nil {
		return errors.New("Queue does not exist")
	}

	return q.List.Clear()
}

// Add a queue behind the current queue, preserving order. This will take ownership of and clear the provided queue.
func (q *Queue) Merge(nq *Queue) error {
	if q == nil {
		return errors.New("Current queue does not exist")
	} else if nq == nil || nq.Count() == 0 {
		// Nothing to add.
		return nil
	}

	if err := q.List.Merge(&nq.List); err != nil {
		return err
	}

	return nq.Clear()
}

// Display queue contents, from the top to the bottom.
func (q *Queue) String() string {
	if q == nil {
		return "<nil>"
	}

	return q.List.String()
}
