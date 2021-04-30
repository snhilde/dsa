// Package hlist provides linked list functionalities.
package hlist

import (
	"fmt"
	"reflect"
	"strings"
)

// This is the standard error message when trying to use an invalid list.
var errBadList = fmt.Errorf("list must be created with New() first")

// List is the main type for this package. It holds the internal information about the linked list.
type List struct {
	head   *hnode
	length int
}

// hnode is an internal type for an individual node in the list.
type hnode struct {
	item interface{}
	next *hnode
}

// New creates a new linked list.
func New() *List {
	return new(List)
}

// String returns a comma-separated list of the string representations of all of the items in the
// linked list.
func (l *List) String() string {
	if l == nil {
		return "<nil>"
	} else if l.head == nil {
		return "<empty>"
	}

	builder := new(strings.Builder)
	for node := l.head; node != nil; node = node.next {
		if builder.Len() > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(fmt.Sprintf("%v", node.item))
	}

	return builder.String()
}

// Length gets the number of nodes in the list, or -1 if list hasn't been created yet.
func (l *List) Length() int {
	if l == nil {
		return -1
	}

	return l.length
}

// Insert inserts one or more items into the list at the specified index.
func (l *List) Insert(index int, items ...interface{}) error {
	// Make sure that none of the items is this list itself.
	for _, v := range items {
		if nl, ok := v.(*List); ok {
			if l.Same(nl) {
				return fmt.Errorf("can't add list to itself")
			}
		}
	}

	node, err := l.getPrior(index)
	if err != nil {
		return err
	}

	if len(items) == 0 {
		return nil
	}

	// Build out the chain of items, then link it in at the specified position.
	begin, end, num := buildChain(items)
	if node == nil {
		// This means the list is empty.
		end.next = l.head
		l.head = begin
	} else {
		end.next = node.next
		node.next = begin
	}
	l.length += num

	return nil
}

// Append adds one or more items to the end of the list.
func (l *List) Append(items ...interface{}) error {
	return l.Insert(l.Length(), items...)
}

// Index gets the index of the first matching item, or -1 if not found.
func (l *List) Index(item interface{}) int {
	if l == nil {
		return -1
	}

	i := 0
	for node := l.head; node != nil; node = node.next {
		if reflect.DeepEqual(node.item, item) {
			return i
		}
		i++
	}

	// If we're here, then we didn't find anything.
	return -1
}

// Item gets the item at the index.
func (l *List) Item(index int) interface{} {
	node, err := l.getPrior(index)
	if err != nil {
		return nil
	}

	// If there isn't a node before the specified index, then that means either the index is 0 or
	// this list is empty.
	if node == nil {
		if l.head == nil {
			return nil
		}
		return l.head.item
	}

	// Because we asked for the previous node, we need to move to the next one, which has the
	// correct item.
	return node.next.item
}

// Exists checks whether or not the item exists in the list.
func (l *List) Exists(item interface{}) bool {
	i := l.Index(item)
	if i < 0 {
		// Index didn't find anything, or the list is invalid.
		return false
	}

	return true
}

// Remove removes an item from the list and returns its value.
func (l *List) Remove(index int) interface{} {
	node, err := l.getPrior(index)
	if err != nil {
		return nil
	}

	// If there isn't a node before the specified index, then that means either the index is 0 or
	// this list is empty.
	var pop *hnode
	if node == nil {
		if l.head == nil {
			return nil
		}
		pop = l.head
		l.head = l.head.next
	} else {
		pop = node.next
		node.next = pop.next
	}

	l.length--
	return pop.item
}

// RemoveMatch finds the first item with a matching value and removes it from the list.
func (l *List) RemoveMatch(value interface{}) {
	// If the item exists in the list, then remove it at the index found.
	if i := l.Index(value); i >= 0 {
		l.Remove(i)
	}
}

// Copy makes an exact copy of the list.
func (l *List) Copy() (*List, error) {
	if l == nil {
		return nil, errBadList
	}

	// We're going to build an identical chain of nodes here, and then we'll create a new List and
	// link in the chain afterwards.
	cp := New()
	cp.head = newNode(nil)
	for node1, node2 := l.head, cp.head; node1 != nil; node1, node2 = node1.next, node2.next {
		node2.next = newNode(node1.item)
	}

	// Skip past the anchor node created at the beginning of the new list.
	cp.head = cp.head.next

	// Match up the lengths.
	cp.length = l.length

	return cp, nil
}

// Same checks if the two lists point to the same underlying data and are therefore the same list.
func (l *List) Same(list2 *List) bool {
	if l == nil || list2 == nil {
		return false
	}

	if l == list2 {
		return true
	}

	return false
}

// Twin checks if the two lists are separate lists but hold the same contents.
func (l *List) Twin(list2 *List) bool {
	if l == nil || list2 == nil {
		return false
	}

	if l.Same(list2) {
		return false
	}

	// The lists must have the same length.
	if l.length != list2.length {
		return false
	}

	for node1, node2 := l.head, list2.head; node1 != nil; node1, node2 = node1.next, node2.next {
		if node2 == nil {
			// Shouldn't ever be possible because the lists are the same length, but need to check
			// for safety.
			return false
		}

		if !reflect.DeepEqual(node1.item, node2.item) {
			return false
		}
	}

	// If we're here, then all the nodes matched up.
	return true
}

// Merge appends the list to the current list, preserving order. This will take ownership of and
// clear the provided list.
func (l *List) Merge(list2 *List) error {
	if l == nil {
		return errBadList
	}

	if list2 == nil {
		// Nothing to do.
		return nil
	}

	// Find the end of the list.
	node, err := l.getPrior(l.length)
	if err != nil {
		return err
	}

	if node == nil {
		// The first list is empty.
		l.head = list2.head
		l.length = list2.length
	} else {
		node.next = list2.head
		l.length += list2.length
	}

	// Give the first list ownership of all nodes.
	list2.Clear()

	return nil
}

// Clear resets the list to its initial state.
func (l *List) Clear() error {
	if l == nil {
		return errBadList
	}

	// Reset all members.
	*l = *(New())

	return nil
}

// Yield provides an unbuffered channel that will continually pass successive items until the list
// is exhausted. The channel quit is used to communicate when iteration should be stopped. Send an
// empty struct (struct{}{}) on the channel to break the communication. This will happen
// automatically if the list is exhausted. If this is not needed, pass nil as the argument. Use
// Yield if you are concerned about memory usage or don't know how far through the list you will
// iterate; otherwise, use YieldAll.
func (l *List) Yield(quit <-chan struct{}) <-chan interface{} {
	if l == nil || l.head == nil {
		return nil
	}

	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for node := l.head; node != nil; node = node.next {
			// Either block on sending this node's item back on the channel, or break out of the
			// loop if the caller is done receiving items.
			select {
			case ch <- node.item:
			case <-quit:
				return
			}
		}
	}()

	return ch
}

// YieldAll provides a buffered channel that will pass successive items until the list is exhausted.
// Use this if you don't care greatly about memory usage and for convenience.
func (l *List) YieldAll() <-chan interface{} {
	if l == nil || l.head == nil {
		return nil
	}

	ch := make(chan interface{}, l.Length())
	for node := l.head; node != nil; node = node.next {
		ch <- node.item
	}
	close(ch)

	return ch
}

// Sort sorts the list using a modified merge algorithm. The comparison function less should return
// true only if left should be sorted before right.
func (l *List) Sort(less func(left, right interface{}) bool) error {
	// We are going to use the merge sort algorithm here. However, because length operations are not
	// constant-time, we are not going to divide the list into progressively smaller blocks.
	// Instead, we are going to assume a block size of 2 and iteratively merge-sort blocks of
	// greater and greater size until the list is fully sorted.
	if l == nil {
		return errBadList
	} else if less == nil {
		return fmt.Errorf("missing comparison callback")
	}

	listLen := l.Length()
	if listLen < 2 {
		// Already sorted.
		return nil
	}

	// Two stacks make up a block. Before each loop, both stacks in a block are sorted. Merging them
	// together in sorted order will yield a sorted block. The smallest stack size of already-sorted
	// nodes is 1. We'll begin to merge stacks from there and work up. When the stack size is at
	// least as big as the entire list, then everything must be sorted.
	for stackLen := 1; stackLen < listLen; stackLen *= 2 {
		blockLen := stackLen * 2
		tmpList := New()
		block := l.head
		numBlocks := (listLen + blockLen - 1) / blockLen
		for i := 0; i < numBlocks; i++ {
			// Get the start of the left stack.
			leftStack := block
			leftLen := stackLen
			// Get the start of the right stack.
			rightStack := block
			rightLen := stackLen
			for j := 0; j < stackLen && rightStack != nil; j++ {
				rightStack = rightStack.next
			}

			// If this is the last block and it's not a full block, then we'll have to handle some
			// special conditions.
			if i+1 == numBlocks && listLen%blockLen != 0 {
				nodesLeft := listLen - (i * blockLen)
				// If we don't even have a full stack, then this block is already in sorted order.
				if nodesLeft <= stackLen {
					// Add the sorted stack/block to the list and move up to the next stack size.
					tailList := New()
					tailList.head = block
					tailList.length = nodesLeft
					tmpList.Merge(tailList)
					continue
				}
				// Shrink our right stack to the correct length.
				rightLen = nodesLeft - stackLen
			}

			// Merge the stacks in sorted order.
			tmpLen := leftLen + rightLen
			for j := 0; j < tmpLen; j++ {
				if leftLen == 0 {
					// Only right stack still has nodes.
					tmpList.Append(rightStack.item)
					rightStack = rightStack.next
					rightLen--
				} else if rightLen == 0 {
					// Only left stack still has nodes.
					tmpList.Append(leftStack.item)
					leftStack = leftStack.next
					leftLen--
				} else if less(leftStack.item, rightStack.item) {
					tmpList.Append(leftStack.item)
					leftStack = leftStack.next
					leftLen--
				} else {
					tmpList.Append(rightStack.item)
					rightStack = rightStack.next
					rightLen--
				}
			}

			// Move to the next block.
			for j := 0; j < tmpLen; j++ {
				block = block.next
			}
		}

		// Hold on to what we have so far.
		l.head = tmpList.head
	}

	return nil
}

// SortInt sorts the list using a modified merge algorithm. Note: all items in the list must be of
// type int.
func (l *List) SortInt() error {
	return l.Sort(func(left, right interface{}) bool {
		return left.(int) < right.(int)
	})
}

// SortStr sorts the list using a modified merge algorithm. Note: all items in the list must be of
// type string.
func (l *List) SortStr() error {
	return l.Sort(func(l, r interface{}) bool {
		return l.(string) < r.(string)
	})
}

// newNode is an internal convenience function for creating a new node.
func newNode(item interface{}) *hnode {
	node := new(hnode)
	node.item = item

	return node
}

// buildChain strings together a new chain of linked nodes with the items provided. It returns the
// first and last nodes in the chain as well as the number of nodes.
func buildChain(items []interface{}) (*hnode, *hnode, int) {
	if len(items) == 0 {
		return nil, nil, 0
	}

	// We're going to start by creating an anchor node, and then we'll build out the chain of nodes
	// from that. When we're done adding nodes, we'll move forward past the anchor node to get to
	// the real start of the chain.
	anchor := newNode(nil)
	tail := anchor
	num := 0
	for _, item := range items {
		tail.next = newNode(item)
		tail = tail.next
		num++
	}

	return anchor.next, tail, num
}

// helper to get the node immediately before the specified index.
func (l *List) getPrior(index int) (*hnode, error) {
	if l == nil {
		return nil, errBadList
	}
	if index < 0 {
		return nil, fmt.Errorf("invalid index")
	}
	if index > l.length {
		return nil, fmt.Errorf("out of bounds")
	}

	// There is no node before the first node.
	if index == 0 {
		return nil, nil
	}

	node := l.head
	for i := 0; i < index-1; i++ {
		if node == nil {
			return nil, fmt.Errorf("error finding node at index")
		}
		node = node.next
	}

	return node, nil
}
