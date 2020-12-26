// Package hqueue provides a simple and lean first-in/first-out queue.
package hqueue

import (
	"fmt"
	"github.com/snhilde/dsa/data_structures/hlist"
)

var (
	// This is the standard error message when trying to use an invalid queue.
	errBadQueue = fmt.Errorf("must create queue with New() first")
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
func (q *Queue) Add(items ...interface{}) error {
	if q == nil {
		return errBadQueue
	}

	// To prevent infinite recursion, make sure that none of the items is this queue itself.
	for _, v := range items {
		if t, ok := v.(*Queue); ok {
			if q.Same(t) {
				return fmt.Errorf("can't add queue to itself")
			}
		}
	}

	return q.list.Append(items...)
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
		return nil, errBadQueue
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
		return errBadQueue
	} else if nq == nil || nq.Count() == 0 {
		// Nothing to add.
		return nil
	}

	if err := q.list.Merge(nq.list); err != nil {
		return err
	}

	return nq.Clear()
}

// Clear resets the queue to its initial state.
func (q *Queue) Clear() error {
	if q == nil {
		return fmt.Errorf("queue does not exist")
	}

	return q.list.Clear()
}

// Same checks whether or not the two queues point to the same underlying data.
func (q *Queue) Same(queue2 *Queue) bool {
	if q == nil || queue2 == nil {
		return false
	}

	return q.list.Same(queue2.list)
}

// String displays the queue's contents, from the top to the bottom.
func (q *Queue) String() string {
	if q == nil {
		return "<nil>"
	}

	return q.list.String()
}
