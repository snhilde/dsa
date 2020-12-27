// Package htree provides an interface to create and use balanced AVL binary trees.
package htree

import (
	"fmt"
	"github.com/snhilde/dsa/data_structures/hstack"
	"reflect"
	"strings"
)

var (
	// This is the standard error message when trying to use an invalid tree.
	errBadTree = fmt.Errorf("tree must be created with New() first")
	// This is the standard error message when trying to use an invalid item.
	errBadItem = fmt.Errorf("item must be created with NewItem() first")
)

// Tree is the main type for this package. It holds information about the entire AVL tree.
type Tree struct {
	trunk *tnode
	count int
}

// Internal structure of tree nodes
type tnode struct {
	item  Item
	bal   int    // Balance of node: -1 if left branch is longer, 0 if both branches are even, or 1 if right side is longer
	left  *tnode // Left branch
	right *tnode // Right branch
}

// New creates a new binary tree.
func New() Tree {
	t := Tree{}

	return t
}

// Add inserts a value into the tree at the provided index. If an item already exists at the index, then it is replaced
// with the new item.
func (t *Tree) Add(value interface{}, index int) error {
	if t == nil {
		return errBadTree
	}

	item := NewItem(value, index)
	return t.AddItems(item)
}

// AddItems inserts one or more items into the tree. If an item already exists at a given index, then it is replaced
// with the new item.
func (t *Tree) AddItems(items ...Item) error {
	if t == nil {
		return errBadTree
	}

	for _, item := range items {
		// Make sure we have a valid item.
		if item == (Item{}) {
			return errBadItem
		}

		// If the tree is empty, start the trunk with this node.
		if t.trunk == nil {
			t.trunk = new(tnode)
			t.trunk.item = item
			t.count++
			continue
		}

		// Find the spot where we need to insert this node.
		node, stack := t.trunk.findNode(item.index)
		if node != nil {
			// We found a matching index. We only need to update the node's value.
			node.item = item
			continue
		}

		// We're at the end. Pop the last branch and add our new item.
		node = stack.Pop().(*tnode)
		if item.index < node.item.index {
			// Add a new item on the left side.
			node.left = new(tnode)
			node.left.item = item
		} else {
			// Add a new item on the right side.
			node.right = new(tnode)
			node.right.item = item
		}
		t.count++

		// Add the node back to the stack and rebalance the tree (if needed).
		stack.Add(node)
		t.rebalance(stack, item.index, true)
	}

	return nil
}

// Remove removes the item at the provided index from the tree.
func (t *Tree) Remove(index int) error {
	if t == nil {
		return errBadTree
	}

	// TODO: implement

	return nil
}

// Item returns the item at the index, or nothing if no item exists at that index.
func (t *Tree) Item(index int) Item {
	if t == nil {
		return Item{}
	}

	node, _ := t.trunk.findNode(index)
	if node == nil {
		return Item{}
	}

	return node.item
}

// Value returns the value of the item at the given index, or nil if no item exists at that index.
func (t *Tree) Value(index int) interface{} {
	item := t.Item(index)
	if item == (Item{}) {
		return nil
	}

	return item.value
}

// Match returns true if the value exists in the tree or false if it does not.
func (t *Tree) Match(value interface{}) bool {
	if t == nil {
		return false
	}

	quit := make(chan interface{})
	itemChan := t.Yield(quit)
	if itemChan == nil {
		return false
	}

	for item := range itemChan {
		if reflect.DeepEqual(value, item.value) {
			// Close the communication and return true.
			quit <- struct{}{}
			return true
		}
	}

	// If we're here, then we didn't find the item.
	return false
}

// Yield provides an unbuffered channel that will continually pass successive items as the tree is traversed in sorted
// order. The channel quit is used to communicate when iteration should be stopped. Send any value on the cnannel to
// break the communication. If this is not needed, pass nil.
func (t *Tree) Yield(quit <-chan interface{}) <-chan Item {
	if t == nil {
		return nil
	}

	itemChan := make(chan Item)
	go func() {
		defer close(itemChan)

		node := t.trunk
		stack := hstack.New()
		for {
			if node == nil {
				// We've reached the end of this left branch. Grab the last node.
				if stack.Count() == 0 {
					// We've traversed all the nodes.
					break
				}

				// Send out the value.
				node = stack.Pop().(*tnode)
				select {
				case itemChan <- node.item:
					// Left branch is done. Work down the right branch now.
					node = node.right
				case <-quit:
					// The caller has notified us that they are done.
					break
				}
			} else {
				// Add the node to the stack and keep going down the left branch.
				stack.Add(node)
				node = node.left
			}
		}
	}()

	return itemChan
}

// List returns copies of all the items in the tree in sorted order.
func (t *Tree) List() []Item {
	if t == nil {
		return nil
	}

	// By using values passed in a channel, we can be sure that the internal values are safe and not modifiable.
	itemChan := t.Yield(nil)
	if itemChan == nil {
		return nil
	}

	list := make([]Item, t.Count())
	i := 0
	for item := range itemChan {
		list[i] = item
		i++
	}

	return list
}

// String returns a printable representation of the items in the tree in sorted order.
func (t *Tree) String() string {
	if t == nil {
		return "<nil>"
	} else if t.Count() == 0 {
		return "<empty>"
	}

	itemChan := t.Yield(nil)
	if itemChan == nil {
		return "<nil>"
	}

	var b strings.Builder
	for item := range itemChan {
		b.WriteString(fmt.Sprintf("%v, ", item.value))
	}

	// Remove the last comma/space before returning the string.
	s := b.String()
	return strings.TrimSuffix(s, ", ")
}

// Count returns the number of items in the tree, or -1 on error.
func (t *Tree) Count() int {
	if t == nil {
		return -1
	}

	return t.count
}

// Item is the type for each item in the tree. It holds the value of the item and its index for sorting.
type Item struct {
	value interface{}
	index int
}

// NewItem creates a new item with the provided value and index.
func NewItem(value interface{}, index int) Item {
	item := Item{
		value: value,
		index: index,
	}

	return item
}

// GetValue returns the value of this item, or nil if the item is bad.
func (i *Item) GetValue() interface{} {
	if i == nil {
		return nil
	}

	return i.value
}

// GetIndex returns the index of this item, or -1 if the item is bad.
func (i *Item) GetIndex() int {
	if i == nil {
		return -1
	}

	return i.index
}

// SetValue changes the value of the item to the provided value.
func (i *Item) SetValue(value interface{}) error {
	if i == nil {
		return errBadItem
	}

	i.value = value

	return nil
}

// findNode will iterate down a tree until it finds a matching index. If no matching index is found, then it will
// return nil for the node. Additionally, it will build a stack of all the nodes traversed on the way.
func (n *tnode) findNode(index int) (*tnode, *hstack.Stack) {
	stack := hstack.New()

	node := n
	for node != nil {
		if index == node.item.index {
			break
		}

		stack.Add(node)
		if index < node.item.index {
			node = node.left
		} else {
			node = node.right
		}
	}

	return node, stack
}

// rebalance will calculate the balances of the nodes in the path and perform any necessary rotation operations to
// rebalance the tree.
func (t *Tree) rebalance(stack *hstack.Stack, index int, added bool) {
	var node *tnode
	for stack.Count() > 0 {
		node = stack.Pop().(*tnode)
		if index < node.item.index {
			if added {
				node.bal--
			} else {
				node.bal++
			}
		} else {
			if added {
				node.bal++
			} else {
				node.bal--
			}
		}

		if (added && node.bal == 0) || (!added && (node.bal == -1 || node.bal == 1)) {
			// The operation did not change the length of the longest branch. We can stop checking for imbalance now.
			break
		} else if node.bal == -2 || node.bal == 2 {
			// We have an imbalance. Rotate the nodes to fix this. This will change the root node of this branch, so
			// we'll need to link it back in after the rotation operation is done.
			// tree.
			rotated := rotate(node, index)
			if stack.Count() == 0 {
				// We're at the top of the tree.
				t.trunk = rotated
			} else {
				node = stack.Pop().(*tnode)
				if index < node.item.index {
					node.left = rotated
				} else {
					node.right = rotated
				}
			}
			break
		}
		// Nothing found yet. Keep going up.
	}
}

// rotate will perform the necessary rotations to rebalance the tree from this node down.
func rotate(top *tnode, index int) *tnode {
	// When rebalancing, we only really care about two nodes: the node that first had the -2 or 2 imbalance and the node
	// directly below it on the insertion side. We'll call these the top node and bottom node. The top node was sent as
	// the first argument to this function. We'll get the bottom node in a bit.
	// To rebalance the tree, we're going to rearrange some nodes around a single node, an operation commonly referred
	// to as a rotation. We'll need to do either a single rotation or a double rotation. If the insertion path is on the
	// same side of both the top and bottom node, then we need to do only a single rotation. If the sides are different,
	// then we'll need to do a double rotation.
	var bottom *tnode
	var left bool
	var double bool

	if index < top.item.index {
		left = true
		bottom = top.left
		if index > bottom.item.index {
			double = true
		}
	} else {
		left = false
		bottom = top.right
		if index < bottom.item.index {
			double = true
		}
	}

	if double {
		// The insertion path is on different sides of the top and bottom nodes, so we have to do a double rotation.
		// We'll do the unique part first here, and then we'll do the shared part below.
		bottom.bal = 0
		if left {
			top.left = bottom.right
			bottom.right = top.left.left
			top.left.left = bottom
			bottom = top.left
		} else {
			top.right = bottom.left
			bottom.left = top.right.right
			top.right.right = bottom
			bottom = top.right
		}
	}

	// Now, we'll do the shared rotation on the top node that all balance operations will need.
	// TODO: I believe the balances are not always going to be returned to 0 here. Investigate this more.
	top.bal = 0
	bottom.bal = 0
	if left {
		if bottom.right == nil {
			top.bal = 1
		}
		top.left = bottom.right
		bottom.right = top
	} else {
		if bottom.left == nil {
			top.bal = -1
		}
		top.right = bottom.left
		bottom.left = top
	}

	// Pass the new top of this branch of the tree back to the caller for proper linking.
	return bottom
}
