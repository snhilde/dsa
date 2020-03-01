// Package hlist provides linked list functionalities.
package hlist

import (
	"errors"
	"strings"
	"fmt"
)


// --- PACKAGE TYPES ---
// List is the main type for this package. It holds the internal information about the list.
type List struct {
	head   *hnode
	length  int
}

// internal type for an individual node in the list
type hnode struct {
	next  *hnode
	value  interface{}
}


// --- ENTRY FUNCTIONS ---
// Create a new linked list.
func New() *List {
	return new(List)
}


// --- LIST METHODS ---
func (list *List) String() string {
	var b strings.Builder

	if list == nil {
		return fmt.Sprintf("<nil>")
	} else if list.head == nil {
		return fmt.Sprintf("<empty>")
	}

	node := list.head
	b.WriteString(fmt.Sprintf("%v", node.value))
	for node.next != nil {
		node = node.next
		b.WriteString(fmt.Sprintf(", %v", node.value))
	}

	return b.String()
}

// Get the number of nodes in the list, or -1 if list hasn't been created yet.
func (list *List) Length() int {
	if list == nil {
		return -1
	}

	return list.length
}

// Insert a value into the list at the specified index.
func (list *List) Insert(value interface{}, index int) error {
	if index < 0 {
		return errors.New("Invalid index")
	} else if list.Length() < index {
		return errors.New("Out of bounds")
	}

	new_node := newNode(value)

	if index == 0 {
		new_node.next = list.head
		list.head = new_node
	} else {
		node := list.head
		for i := 1; i < index; i++ {
			node = node.next
		}

		new_node.next = node.next
		node.next = new_node
	}

	list.length++
	return nil
}

// Add one or more values to the end of the list.
func (list *List) Append(values ...interface{}) error {
	if list == nil {
		return errors.New("List must be created with New() first")
	} else if len(values) == 0 {
		return nil
	}

	tmp_list := New()
	node := tmp_list.head
	for _, v := range values {
		if tmp_list.head == nil {
			tmp_list.head = newNode(v)
			tmp_list.length++
			node = tmp_list.head
		} else {
			node.next = newNode(v)
			node = node.next
			tmp_list.length++
		}
	}

	list.Merge(tmp_list)

	return nil
}

// Remove an item for the list and return its value.
func (list *List) Pop(index int) interface{} {
	if list == nil || index < 0 || index >= list.Length() {
		return nil
	}

	var pop *hnode
	// Handle the special case of popping the first node.
	if index == 0 {
		pop = list.head
		list.head = pop.next
	} else {
		node := list.head
		for i := 0; i < index-1; i++ {
			node = node.next
		}
		pop = node.next
		node.next = pop.next
	}

	list.length--
	return pop.value
}

// Find an item by matching value and remove it from the list.
func (list *List) PopMatch(value interface{}) bool {
	if list == nil {
		return false
	}

	// Handle the special case of matching on the first node.
	if list.head.value == value {
		pop := list.head
		list.head = pop.next
		list.length--
		return true
	}

	node := list.head
	for node.next != nil {
		if node.next.value == value {
			pop := node.next
			node.next = pop.next
			list.length--
			return true
		}
		node = node.next
	}

	// If we're here, then we didn't find anything.
	return false
}

// Get the index of the first matching node, if any.
func (list *List) Index(value interface{}) int {
	if list == nil {
		return -1
	}

	length := list.Length()
	node := list.head
	for i := 0; i < length; i++ {
		if node.value == value {
			return i
		}
		node = node.next
	}

	// If we're here, then we didn't find anything.
	return -1
}

// Check whether or not the value exists in the list.
func (list *List) Exists(value interface{}) bool {
	if list == nil {
		return false
	}

	node := list.head
	for node != nil {
		if node.value == value {
			return true
		}
		node = node.next
	}

	// If we're here, then we didn't find anything.
	return false
}

// Append new list to current list. The current list will take ownership of all nodes.
func (list *List) Merge(addition *List) error {
	if list == nil {
		return errors.New("List must be created with New() first")
	} else if addition == nil {
		return nil
	}

	// Find the end of the list.
	if list.head == nil {
		list.head = addition.head
		list.length = addition.length
	} else {
		node := list.head
		for node.next != nil {
			node = node.next
		}
		node.next = addition.head
		list.length += addition.length
	}

	addition.Clear()

	return nil
}

// Clear the list to its inital state.
func (list *List) Clear() {
	if list == nil {
		return
	}

	list.head = nil
	list.length = 0
}

// Sort the list using a modified merge algorithm.
// equality_cb should return a negative value if left should be sorted first or a positive value if right should be sorted first.
func (list *List) Sort(equality_cb func(left, right interface{}) int) error {
	// We are going to use the merge sort algorithm here. However, because length operations are not constant-time, we
	// are not going to divide the list into progressively smaller blocks. Instead, we are going to assume a block size
	// of 2 and iteratively merge-sort blocks of greater and greater size until the list is fully sorted.
	if list == nil {
		return errors.New("List must be created with New() first")
	}

	list_length := list.Length()
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
		block := list.head
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
					tmp_list.Append(right_stack.value)
					right_stack = right_stack.next
					right_len--
				} else if right_len == 0 {
					// Only left stack still has nodes.
					tmp_list.Append(left_stack.value)
					left_stack = left_stack.next
					left_len--
				} else if equality_cb(left_stack.value, right_stack.value) < 0 {
					tmp_list.Append(left_stack.value)
					left_stack = left_stack.next
					left_len--
				} else {
					tmp_list.Append(right_stack.value)
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
		list.head = tmp_list.head
	}

	return nil
}

// Sort the list using a modified merge algorithm.
// Note: all values in the list must be of type int.
func (list *List) SortInt() error {
	return list.Sort(eqInt)
}

// Sort the list using a modified merge algorithm.
// Note: all values in the list must be of type string.
func (list *List) SortStr() error {
	return list.Sort(eqStr)
}


// --- HELPER FUNCTIONS ---
// internal convenience function for creating a new node
func newNode(value interface{}) *hnode {
	node := new(hnode)
	node.value = value

	return node
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
