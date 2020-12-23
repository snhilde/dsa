// Package hlist provides linked list functionalities.
package hlist

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	// This is the standard error message when trying to use an invalid list.
	errBadList = fmt.Errorf("list must be created with New() first")
)

// List is the main type for this package. It holds the internal information about the list.
type List struct {
	head   *hnode
	length int
}

// internal type for an individual node in the list
type hnode struct {
	v    interface{}
	next *hnode
}

// New creates a new linked list.
func New() *List {
	return new(List)
}

// String returns a comma-separated list of the string representations of all of the nodes in the linked list.
func (l *List) String() string {
	var b strings.Builder

	if l == nil {
		return "<nil>"
	} else if l.head == nil {
		return "<empty>"
	}

	n := l.head
	for n != nil {
		b.WriteString(fmt.Sprintf("%v", n.v))
		n = n.next
		if n != nil {
			b.WriteString(", ")
		}
	}

	return b.String()
}

// Length gets the number of nodes in the list, or -1 if list hasn't been created yet.
func (l *List) Length() int {
	if l == nil {
		return -1
	}

	return l.length
}

// Insert inserts one or more value into the list at the specified index.
func (l *List) Insert(index int, vs ...interface{}) error {
	n, err := l.getPrior(index)
	if err != nil {
		return err
	}

	if len(vs) == 0 {
		return nil
	}

	// Make temporary list.
	head := newNode(nil)
	nn := head
	for _, v := range vs {
		nn.next = newNode(v)
		nn = nn.next
		l.length++
	}

	// Move past the node we created to make adding smoother.
	head = head.next

	// Link in the temporary list.
	if n == nil {
		// Handle the special case of inserting the first node.
		nn.next = l.head
		l.head = head
	} else {
		nn.next = n.next
		n.next = head
	}

	return nil
}

// Append adds one or more values to the end of the list.
func (l *List) Append(values ...interface{}) error {
	if l == nil {
		return errBadList
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

// Index gets the index of the first matching node, or -1 if not found.
func (l *List) Index(v interface{}) int {
	if l == nil {
		return -1
	}

	n := l.head
	i := 0
	for n != nil {
		if reflect.DeepEqual(n.v, v) {
			return i
		}
		n = n.next
		i++
	}

	// If we're here, then we didn't find anything.
	return -1
}

// Value gets the value at the index.
func (l *List) Value(index int) interface{} {
	if l == nil || l.head == nil {
		return nil
	}

	n, err := l.getPrior(index)
	if err != nil {
		return nil
	}

	if n == nil {
		return l.head.v
	}

	return n.next.v
}

// Exists checks whether or not the value exists in the list.
func (l *List) Exists(v interface{}) bool {
	i := l.Index(v)
	if i < 0 {
		// Index didn't find anything, or the list is invalid.
		return false
	}

	return true
}

// Remove removes an item from the list and returns its value.
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

// RemoveMatch finds the first item with a matching value and removes it from the list.
func (l *List) RemoveMatch(v interface{}) {
	i := l.Index(v)
	if i < 0 {
		// Index() didn't find anything, or the list is invalid.
		return
	}

	l.Remove(i)
	return
}

// Copy makes an exact copy of the list.
func (l *List) Copy() (*List, error) {
	if l == nil {
		return nil, errBadList
	}

	// We'll add a helper node to the beginning of the new list to make adding the other nodes easier.
	nl := New()
	nl.head = newNode(nil)

	n := l.head
	nn := nl.head
	for n != nil {
		nn.next = newNode(n.v)
		n = n.next
		nn = nn.next
		nl.length++
	}

	// Get rid of the helper node.
	nl.head = nl.head.next

	return nl, nil
}

// Same checks if the two lists point to the same underlying data and are therefore the same list.
func (l *List) Same(nl *List) bool {
	if l == nil || nl == nil {
		return false
	}

	if l == nl {
		return true
	}

	return false
}

// Twin checks if the two lists are separate lists but hold the same contents.
func (l *List) Twin(nl *List) bool {
	if l == nil || nl == nil {
		return false
	}

	if l.Same(nl) {
		return false
	}

	// The lists must have the same length.
	if l.length != nl.length {
		return false
	}

	// The lists must not point to the same nodes.
	if l.head == nl.head {
		return false
	}

	n := l.head
	nn := nl.head
	for n != nil {
		if !reflect.DeepEqual(n.v, nn.v) {
			return false
		}
		n = n.next
		nn = nn.next
	}

	// If we're here, then all the nodes matched up.
	return true
}

// Merge appends the list to the current list, preserving order. This will take ownership of and clear the provided list.
func (l *List) Merge(addition *List) error {
	if l == nil {
		return errBadList
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

// Clear resets the list to its inital state.
func (l *List) Clear() error {
	if l == nil {
		return errBadList
	}

	l.head = nil
	l.length = 0

	return nil
}

// Yield provides an unbuffered channel that will continually pass successive node values until the list is exhausted.
// The channel quit is used to communicate when iteration should be stopped. Send any value on the cnannel (or close it)
// to break the communication. This will happen automatically if the list is exhausted. If this is not needed, pass nil
// as the argument. Use Yield if you are concerned about memory usage or don't know how far through the list you will
// iterate; otherwise, use YieldAll.
func (l *List) Yield(quit <-chan interface{}) <-chan interface{} {
	if l == nil || l.head == nil {
		return nil
	}

	ch := make(chan interface{})
	go func() {
		defer close(ch)
		n := l.head
		for n != nil {
			select {
			case ch <- n.v:
				n = n.next
			case <-quit:
				return
			}
		}
	}()

	return ch
}

// YieldAll provides a buffered channel that will pass successive node values until the list is exhausted.
// Use this if you don't care greatly about memory usage and for convenience.
func (l *List) YieldAll() <-chan interface{} {
	if l == nil || l.head == nil {
		return nil
	}

	ch := make(chan interface{}, l.Length())
	n := l.head
	for n != nil {
		ch <- n.v
		n = n.next
	}
	close(ch)

	return ch
}

// Sort sorts the list using a modified merge algorithm. cmp should return true if left should be sorted first or false
// if right should be sorted first.
func (l *List) Sort(cmp func(left, right interface{}) bool) error {
	// We are going to use the merge sort algorithm here. However, because length operations are not constant-time, we
	// are not going to divide the list into progressively smaller blocks. Instead, we are going to assume a block size
	// of 2 and iteratively merge-sort blocks of greater and greater size until the list is fully sorted.
	if l == nil {
		return errBadList
	} else if cmp == nil {
		return fmt.Errorf("missing equality comparison callback")
	}

	listLen := l.Length()
	if listLen < 2 {
		// Already sorted.
		return nil
	}

	// Two stacks make up a block. Before each loop, both stacks in a block are sorted. Merging them together in sorted
	// order will yield a sorted block. The smallest stack size of already-sorted nodes is 1. We'll begin to merge
	// stacks from there and work up. When the stack size is at least as big as the entire list, then everything must be
	// sorted.
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

			// If this is the last block and it's not a full block, then we'll have to handle some special conditions.
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
					tmpList.Append(rightStack.v)
					rightStack = rightStack.next
					rightLen--
				} else if rightLen == 0 {
					// Only left stack still has nodes.
					tmpList.Append(leftStack.v)
					leftStack = leftStack.next
					leftLen--
				} else if cmp(leftStack.v, rightStack.v) {
					tmpList.Append(leftStack.v)
					leftStack = leftStack.next
					leftLen--
				} else {
					tmpList.Append(rightStack.v)
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

// SortInt sorts the list using a modified merge algorithm. Note: all values in the list must be of type int.
func (l *List) SortInt() error {
	return l.Sort(cmpInt)
}

// SortStr sorts the list using a modified merge algorithm. Note: all values in the list must be of type string.
func (l *List) SortStr() error {
	return l.Sort(cmpStr)
}

// internal convenience function for creating a new node
func newNode(v interface{}) *hnode {
	n := new(hnode)
	n.v = v

	return n
}

// integer equality callback for SortInt() method
func cmpInt(left, right interface{}) bool {
	if left.(int) < right.(int) {
		return true
	}

	return false
}

// string equality callback for SortStr() method
func cmpStr(l, r interface{}) bool {
	lv := []rune(l.(string))
	rv := []rune(r.(string))

	for i, v := range lv {
		if i == len(rv) {
			// If we're here, then r is a prefix of l.
			return false
		}

		if v == rv[i] {
			continue
		} else if v < rv[i] {
			return true
		} else {
			return false
		}
	}

	// If we're here, then either l is a prefix of r or both strings are equal.
	return true
}

// helper to get the node immediately before the specified index.
func (l *List) getPrior(index int) (*hnode, error) {
	if l == nil {
		return nil, errBadList
	} else if index < 0 {
		return nil, fmt.Errorf("invalid index")
	} else if index > l.length {
		return nil, fmt.Errorf("out of bounds")
	}

	if l.head == nil || index == 0 {
		return nil, nil
	}

	n := l.head
	for i := 0; i < index-1; i++ {
		if n == nil {
			return nil, fmt.Errorf("error finding node at index")
		}
		n = n.next
	}

	return n, nil
}
