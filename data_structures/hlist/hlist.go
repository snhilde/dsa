// Package hlist provides linked list functionalities.
package hlist

import (
	"errors"
	"strings"
	"fmt"
)


// List is the main type for this package. It holds the internal information about the list.
type List struct {
	head   *hnode
	length  int
}

// internal type for an individual node in the list
type hnode struct {
	v     interface{}
	next *hnode
}


// Create a new linked list.
func New() *List {
	return new(List)
}


func (l *List) String() string {
	var b strings.Builder

	if l == nil {
		return "<nil>"
	} else if l.head == nil {
		return "<empty>"
	}

	n := l.head
	b.WriteString(fmt.Sprintf("%v", n.v))
	for n.next != nil {
		n = n.next
		b.WriteString(fmt.Sprintf(", %v", n.v))
	}

	return b.String()
}

// Get the number of nodes in the list, or -1 if list hasn't been created yet.
func (l *List) Length() int {
	if l == nil {
		return -1
	}

	return l.length
}

// Insert a value into the list at the specified index.
func (l *List) Insert(v interface{}, index int) error {
	n, err := l.getPrior(index)
	if err != nil {
		return err
	}

	nn := newNode(v)
	if n == nil {
		// Handle the special case of inserting the first node.
		nn.next = l.head
		l.head = nn
	} else {
		nn.next = n.next
		n.next = nn
	}

	l.length++
	return nil
}

// Add one or more values to the end of the list.
func (l *List) Append(values ...interface{}) error {
	if l == nil {
		return lErr()
	}

	tmp := New()
	n := tmp.head
	for _, v := range values {
		if tmp.head == nil {
			tmp.head = newNode(v)
			tmp.length++
			n = tmp.head
		} else {
			n.next = newNode(v)
			n = n.next
			tmp.length++
		}
	}

	l.Merge(tmp)

	return nil
}

// Remove an item from the list and return its value.
func (l *List) Remove(index int) interface{} {
	if l == nil || l.head == nil {
		return nil
	}

	n, err := l.getPrior(index)
	if err != nil {
		return nil
	}

	var pop *hnode
	if n == nil {
		// Handle the special case of popping the first node.
		pop = l.head
		l.head = l.head.next
	} else {
		pop = n.next
		n.next = pop.next
	}

	l.length--
	return pop.v
}

// Find the first item with a matching value, and remove it from the list.
func (l *List) RemoveMatch(v interface{}) error {
	if l == nil {
		return lErr()
	}

	if l.head == nil {
		// Nothing to match.
		return nil
	}

	if l.head.v == v {
		// Handle the special case of matching on the first node.
		pop := l.head
		l.head = pop.next
		l.length--
	} else {
		n := l.head
		for n.next != nil {
			if n.next.v == v {
				pop := n.next
				n.next = pop.next
				l.length--
				break
			}
			n = n.next
		}
	}

	return nil
}

// Get the index of the first matching node, or -1 if not found.
func (l *List) Index(v interface{}) int {
	if l == nil {
		return -1
	}

	n := l.head
	i := 0
	for n != nil {
		if n.v == v {
			return i
		}
		n = n.next
		i++
	}

	// If we're here, then we didn't find anything.
	return -1
}

// Check whether or not the value exists in the list.
func (l *List) Exists(v interface{}) bool {
	if l == nil {
		return false
	}

	n := l.head
	for n != nil {
		if n.v == v {
			return true
		}
		n = n.next
	}

	// If we're here, then we didn't find anything.
	return false
}

// Append new list to current list. The current list will take ownership of all nodes.
func (l *List) Merge(addition *List) error {
	if l == nil {
		return lErr()
	}

	if addition == nil {
		// Nothing to do.
		return nil
	}

	// Find the end of the list.
	n, err := l.getPrior(l.length)
	if err != nil {
		return err
	}

	if n == nil {
		// List is empty.
		l.head = addition.head
		l.length = addition.length
	} else {
		n.next = addition.head
		l.length += addition.length
	}

	addition.Clear()
	return nil
}

// Clear the list to its inital state.
func (l *List) Clear() error {
	if l == nil {
		return lErr()
	}

	l.head = nil
	l.length = 0

	return nil
}

// Sort the list using a modified merge algorithm.
// equality_cb should return a negative value if left should be sorted first or a positive value if right should be sorted first.
func (l *List) Sort(equality_cb func(left, right interface{}) int) error {
	// We are going to use the merge sort algorithm here. However, because length operations are not constant-time, we
	// are not going to divide the list into progressively smaller blocks. Instead, we are going to assume a block size
	// of 2 and iteratively merge-sort blocks of greater and greater size until the list is fully sorted.
	if l == nil {
		return lErr()
	}

	list_length := l.Length()
	if list_length < 2 {
		// Already sorted.
		return nil
	}

	// Two stacks make up a block. Before each loop, both stacks in a block are sorted. Merging them together in sorted
	// order will yield a sorted block. The smallest stack size of already-sorted nodes is 1. We'll begin to merge
	// stacks from there and work up. When the stack size is at least as big as the entire list, then everything must be
	// sorted.
	for stack_len := 1; stack_len < list_length; stack_len *= 2 {
		block_len := stack_len * 2
		tmp_list := New()
		block := l.head
		num_blocks := (list_length + block_len - 1) / block_len
		for i := 0; i < num_blocks; i++ {
			// Get the start of the left stack.
			left_stack := block
			left_len := stack_len
			// Get the start of the right stack.
			right_stack := block
			right_len := stack_len
			for j := 0; j < stack_len && right_stack != nil; j++ {
				right_stack = right_stack.next
			}

			// If this is the last block and it's not a full block, then we'll have to handle some special conditions.
			if i+1 == num_blocks && list_length % block_len != 0 {
				nodes_left := list_length - (i * block_len)
				// If we don't even have a full stack, then this block is already in sorted order.
				if nodes_left <= stack_len {
					// Add the sorted stack/block to the list and move up to the next stack size.
					tail_list := New()
					tail_list.head = block
					tail_list.length = nodes_left
					tmp_list.Merge(tail_list)
					continue
				}
				// Shrink our right stack to the correct length.
				right_len = nodes_left - stack_len
			}

			// Merge the stacks in sorted order.
			tmp_len := left_len + right_len
			for j := 0; j < tmp_len; j++ {
				if left_len == 0 {
					// Only right stack still has nodes.
					tmp_list.Append(right_stack.v)
					right_stack = right_stack.next
					right_len--
				} else if right_len == 0 {
					// Only left stack still has nodes.
					tmp_list.Append(left_stack.v)
					left_stack = left_stack.next
					left_len--
				} else if equality_cb(left_stack.v, right_stack.v) < 0 {
					tmp_list.Append(left_stack.v)
					left_stack = left_stack.next
					left_len--
				} else {
					tmp_list.Append(right_stack.v)
					right_stack = right_stack.next
					right_len--
				}
			}

			// Move to the next block.
			for j := 0; j < tmp_len; j++ {
				block = block.next
			}
		}

		// Hold on to what we have so far.
		l.head = tmp_list.head
	}

	return nil
}

// Sort the list using a modified merge algorithm.
// Note: all values in the list must be of type int.
func (l *List) SortInt() error {
	return l.Sort(eqInt)
}

// Sort the list using a modified merge algorithm.
// Note: all values in the list must be of type string.
func (l *List) SortStr() error {
	return l.Sort(eqStr)
}


// internal convenience function for creating a new node
func newNode(v interface{}) *hnode {
	n := new(hnode)
	n.v = v

	return n
}

// integer equality callback for SortInt() method
func eqInt(left, right interface{}) int {
	return left.(int) - right.(int)
}

// string equality callback for SortStr() method
func eqStr(left, right interface{}) int {
	var min int

	l := left.(string)
	r := right.(string)

	lLen := len(l)
	rLen := len(r)

	if lLen == rLen || lLen < rLen {
		min = lLen
	} else {
		min = rLen
	}

	for i := 0; i < min; i++ {
		if l[i] == r[i] {
			continue
		} else if l[i] < r[i] {
			return -1
		} else {
			return 1
		}
	}

	// If we're here, then one of two things happened: either both values are the same or one value is a substring of
	// another. We'll compare based on length to favor the shorter value.
	return lLen - rLen
}

// helper to return standard error on bad list.
func lErr() error {
	return errors.New("List must be created with New() first")
}

// helper to get the node immediately before the specified index.
func (l *List) getPrior(index int) (*hnode, error) {
	if l == nil {
		return nil, lErr()
	} else if index < 0 {
		return nil, errors.New("Invalid index")
	} else if l.length < index {
		return nil, errors.New("Out of bounds")
	}

	if l.head == nil || index == 0 {
		return nil, nil
	}

	n := l.head
	for i := 0; i < index-1; i++ {
		if n == nil {
			return nil, errors.New("Error finding node at index")
		}
		n = n.next
	}

	return n, nil
}
