// Package hqueue provides a simple and lean first-in/first-out queue.
package hqueue

import (
	"errors"
	"github.com/snhilde/dsa/data_structures/hlist"
)

var (
	// This is the standard error message when trying to use an invalid queue.
	badQueue = errors.New("Must create queue with New() first")
)

// Queue is the main type for this package. It holds the internal information about the queue.
type Queue struct {
	list *hlist.List
}

// New creates a new queue.
func New() *Queue {
	q := new(Queue)
	q.list = hlist.New()
	return q
}

// Add adds one or more new items to the back of the queue. The items will be added in the order provided.
func (q *Queue) Add(vs ...interface{}) error {
	if q == nil {
		return badQueue
	}

	// If caller is trying to add own queue, duplicate it first and then add it.
	for i, v := range vs {
		if t, ok := v.(*Queue); ok {
			if t == q {
				nq, err := q.Copy()
				if err != nil {
					return err
				}
				vs[i] = nq
			}
		}
	}

	return q.list.Append(vs...)
}

// Pop removes the first item in the queue and returns its value.
func (q *Queue) Pop() interface{} {
	if q == nil {
		return nil
	}

	return q.list.Remove(0)
}

// Count gets the current number of items in the queue.
func (q *Queue) Count() int {
	if q == nil {
		return -1
	}

	return q.list.Length()
}

// Copy makes an exact copy of the queue.
func (q *Queue) Copy() (*Queue, error) {
	if q == nil {
		return nil, badQueue
	}

	nl, err := q.list.Copy()
	if err != nil {
		return nil, err
	}

	nq := New()
	nq.list = nl

	return nq, nil
}

// Merge adds a queue behind the current queue, preserving order. This will take ownership of and clear the provided
// queue.
func (q *Queue) Merge(nq *Queue) error {
	if q == nil {
		return badQueue
	} else if nq == nil || nq.Count() == 0 {
		// Nothing to add.
		return nil
	}

	if err := q.list.Merge(nq.list); err != nil {
		return err
	}

	return nq.Clear()
}

// Clear resets the queue to its inital state.
func (q *Queue) Clear() error {
	if q == nil {
		return errors.New("Queue does not exist")
	}

	return q.list.Clear()
}

// String displays the queue's contents, from the top to the bottom.
func (q *Queue) String() string {
	if q == nil {
		return "<nil>"
	}

	return q.list.String()
}
