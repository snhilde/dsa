// Package hqueue provides a simple and lean first-in/first-out queue.
package hqueue

import (
	"github.com/snhilde/dsa/data_structures/hlist"
	"errors"
)


// Queue is the main type for this package. It holds the internal information about the queue.
type Queue struct {
	list *hlist.List
}

// Create a new queue.
func New() *Queue {
	q := new(Queue)
	q.list = hlist.New()
	return q
}

// Add a new item or items to the back of the queue. If there is more than one item, the first item will be at the front
// of the queue.
func (q *Queue) Add(vs ...interface{}) error {
	if q == nil {
		return qErr()
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

// Pop the first item in the queue.
func (q *Queue) Pop() interface{} {
	if q == nil {
		return nil
	}

	return q.list.Remove(0)
}

// Get the current number of items in the queue.
func (q *Queue) Count() int {
	if q == nil {
		return -1
	}

	return q.list.Length()
}

// Clear the queue to its inital state.
func (q *Queue) Clear() error {
	if q == nil {
		return errors.New("Queue does not exist")
	}

	return q.list.Clear()
}

// Make an exact copy of the queue.
func (q *Queue) Copy() (*Queue, error) {
	if q == nil {
		return nil, qErr()
	}

	nl, err := q.list.Copy()
	if err != nil {
		return nil, err
	}

	nq := New()
	nq.list = nl

	return nq, nil
}

// Add a queue behind the current queue, preserving order. This will take ownership of and clear the provided queue.
func (q *Queue) Merge(nq *Queue) error {
	if q == nil {
		return qErr()
	} else if nq == nil || nq.Count() == 0 {
		// Nothing to add.
		return nil
	}

	if err := q.list.Merge(nq.list); err != nil {
		return err
	}

	return nq.Clear()
}

// Display queue contents, from the top to the bottom.
func (q *Queue) String() string {
	if q == nil {
		return "<nil>"
	}

	return q.list.String()
}


func qErr() error {
	return errors.New("Must create queue with New() first")
}
