// Package htree provides an interface to create and use balanced AVL binary trees.
package htree

import (
	"fmt"
	"github.com/snhilde/dsa/data_structures/hqueue"
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
	root  *tnode
	count int
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

		// If the tree is empty, start the root with this node.
		if t.root == nil {
			t.root = newNode(item)
			t.count++
			continue
		}

		// Find the spot where we need to insert this node.
		node, stack := t.root.findNode(item.GetIndex())
		if node != nil {
			// We found a matching index. We only need to update the node's value.
			node.item = item
			continue
		}

		// We're at the end. Pop the last node and add our new item.
		parent := stack.Pop().(*tnode)
		if item.index < parent.index() {
			// Add a new item on the left side.
			parent.left = newNode(item)
		} else {
			// Add a new item on the right side.
			parent.right = newNode(item)
		}
		t.count++

		// Add the parent back to the stack and rebalance the tree (if needed).
		stack.Add(parent)
		t.rebalance(stack, item.GetIndex(), true)
	}

	return nil
}

// Remove removes and returns the item at the provided index from the tree.
func (t *Tree) Remove(index int) Item {
	if t == nil {
		return Item{}
	}

	// TODO

	return Item{}
}

// Clear clears all items in the tree and resets it to a new state.
func (t *Tree) Clear() {
	if t != nil {
		*t = (New())
	}
}

// Item returns the item at the index, or nothing if no item exists at that index.
func (t *Tree) Item(index int) Item {
	if t == nil {
		return Item{}
	}

	node, _ := t.root.findNode(index)
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
			// Close the communication and return true. If Yield has finished sending everything, then it won't be
			// listening on quit anymore.
			select {
			case quit <- struct{}{}:
			default:
			}
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
	if t == nil || t.Count() == 0 {
		return nil
	}

	itemChan := make(chan Item)
	go func() {
		defer close(itemChan)

		node := t.root
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
	if t == nil || t.Count() == 0 {
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

// DFS traverses the tree in a depth-first search pattern and returns all values in the order encountered.
func (t *Tree) DFS() []interface{} {
	// Even though we could use other methods here like List or Yield, we're going to do a direct implementation for
	// better performance.
	if t == nil || t.Count() == 0 {
		return nil
	}

	numNodes := t.Count()
	values := make([]interface{}, numNodes)

	stack := hstack.New()
	node := t.root

	i := 0
	for i < numNodes {
		if node == nil {
			node = stack.Pop().(*tnode)
			values[i] = node.item.GetValue()
			i++
			node = node.right
		} else {
			stack.Add(node)
			node = node.left
		}
	}

	return values
}

// BFS traverses the tree in a breadth-first search pattern and returns all values in the order encountered.
func (t *Tree) BFS() []interface{} {
	if t == nil || t.Count() == 0 {
		return nil
	}

	values := make([]interface{}, t.Count())

	queue := hqueue.New()
	queue.Add(t.root)

	for i := 0; i < t.Count(); i++ {
		node := queue.Pop().(*tnode)
		values[i] = node.value()

		if node.left != nil {
			queue.Add(node.left)
		}
		if node.right != nil {
			queue.Add(node.right)
		}
	}

	return values
}

// LeftCenterRight traverses the tree in a left-center-right search pattern and returns all values in the order
// encountered. When a node is reached, this descends to the left first, then reads the value at the current node, then
// descends to the right. This is equivalent to a depth-first search.
func (t *Tree) LeftCenterRight() []interface{} {
	if t == nil || t.Count() == 0 {
		return nil
	}

	values := make([]interface{}, t.Count())

	stack := hstack.New()
	node := t.root

	i := 0
	for i < t.Count() {
		if node == nil {
			node = stack.Pop().(*tnode)
			values[i] = node.value()
			i++
			node = node.right
		} else {
			stack.Add(node)
			node = node.left
		}
	}

	return values
}

// CenterLeftRight traverses the tree in a center-left-right search pattern and returns all values in the order
// encountered. When a node is reached, this reads the value at the current node first, then descends to the left, then
// descends to the right.
func (t *Tree) CenterLeftRight() []interface{} {
	if t == nil || t.Count() == 0 {
		return nil
	}

	values := make([]interface{}, t.Count())

	stack := hstack.New()
	node := t.root

	i := 0
	for i < t.Count() {
		if node == nil {
			node = stack.Pop().(*tnode)
			node = node.right
		} else {
			values[i] = node.value()
			i++
			stack.Add(node)
			node = node.left
		}
	}

	return values
}

// LeftRightCenter traverses the tree in a left-right-center search pattern and returns all values in the order
// encountered. When a node is reached, this descends to the left first, then descends to the right, then reads the
// value at the current node.
func (t *Tree) LeftRightCenter() []interface{} {
	type wrapper struct {
		node      *tnode
		rightDone bool
	}

	if t == nil || t.Count() == 0 {
		return nil
	}

	values := make([]interface{}, t.Count())

	stack := hstack.New()
	node := t.root

	i := 0
	for i < t.Count() {
		if node == nil {
			d := stack.Pop().(wrapper)
			node = d.node
			if d.rightDone {
				values[i] = node.value()
				i++
				node = nil
			} else {
				d.rightDone = true
				stack.Add(d)
				node = node.right
			}
		} else {
			stack.Add(wrapper{node, false})
			node = node.left
		}
	}

	return values
}

// RightCenterLeft traverses the tree in a right-center-left search pattern and returns all values in the order
// encountered. When a node is reached, this descends to the right first, then reads the value at the current node, then
// descends to the left. This is equivalent to a reverse depth-first search.
func (t *Tree) RightCenterLeft() []interface{} {
	if t == nil || t.Count() == 0 {
		return nil
	}

	values := make([]interface{}, t.Count())

	stack := hstack.New()
	node := t.root

	i := 0
	for i < t.Count() {
		if node == nil {
			node = stack.Pop().(*tnode)
			values[i] = node.value()
			i++
			node = node.left
		} else {
			stack.Add(node)
			node = node.right
		}
	}

	return values
}

// CenterRightLeft traverses the tree in a center-right-left search pattern and returns all values in the order
// encountered. When a node is reached, this reads the value at the current node first, then descends to the right, then
// descends to the left.
func (t *Tree) CenterRightLeft() []interface{} {
	if t == nil || t.Count() == 0 {
		return nil
	}

	values := make([]interface{}, t.Count())

	stack := hstack.New()
	node := t.root

	i := 0
	for i < t.Count() {
		if node == nil {
			node = stack.Pop().(*tnode)
			node = node.left
		} else {
			values[i] = node.value()
			i++
			stack.Add(node)
			node = node.right
		}
	}

	return values
}

// RightLeftCenter traverses the tree in a right-left-center search pattern and returns all values in the order
// encountered. When a node is reached, this descends to the right first, then descends to the left, then reads the
// value at the current node.
func (t *Tree) RightLeftCenter() []interface{} {
	type wrapper struct {
		node     *tnode
		leftDone bool
	}

	if t == nil || t.Count() == 0 {
		return nil
	}

	values := make([]interface{}, t.Count())

	stack := hstack.New()
	node := t.root

	i := 0
	for i < t.Count() {
		if node == nil {
			d := stack.Pop().(wrapper)
			node = d.node
			if d.leftDone {
				values[i] = node.value()
				i++
				node = nil
			} else {
				d.leftDone = true
				stack.Add(d)
				node = node.left
			}
		} else {
			stack.Add(wrapper{node, false})
			node = node.right
		}
	}

	return values
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

// Count returns the number of items in the tree.
func (t *Tree) Count() int {
	if t == nil {
		return 0
	}

	return t.count
}

// Height returns the length of the longest branch from root to leaf.
func (t *Tree) Height() int {
	if t == nil || t.root == nil {
		return 0
	}

	return t.root.height
}

// Internal structure of tree nodes
type tnode struct {
	item   Item
	height int    // Longest height of subtree from this node down
	left   *tnode // Left branch
	right  *tnode // Right branch
}

// newNode returns a new tnode with item item.
func newNode(item Item) *tnode {
	n := new(tnode)
	n.item = item
	n.height = 1

	return n
}

// value returns the value of the item at this node.
func (n *tnode) value() interface{} {
	if n == nil {
		return nil
	}

	return n.item.GetValue()
}

// index returns the index of the item at this node.
func (n *tnode) index() int {
	if n == nil {
		return 0
	}

	return n.item.GetIndex()
}

// balance returns the balance of the tree in the set {-2,-1,0,1,2} as per the rules of AVL trees.
func (n *tnode) balance() int {
	if n == nil {
		return 0
	}

	return n.rightCount() - n.leftCount()
}

// leftCount returns the number of nodes on the left side of this node.
func (n *tnode) leftCount() int {
	if n != nil && n.left != nil {
		return n.left.height
	}
	return 0
}

// rightCount returns the number of nodes on the right side of this node.
func (n *tnode) rightCount() int {
	if n != nil && n.right != nil {
		return n.right.height
	}
	return 0
}

// Item is the type for each item in the tree. It holds the value of the item and its index for sorting.
type Item struct {
	value interface{}
	index int
}

// NewItem creates a new item with the provided value, stored at index.
func NewItem(value interface{}, index int) Item {
	item := Item{
		value: value,
		index: index,
	}

	return item
}

// GetValue returns the value of this item, or nil if the item is invalid.
func (i *Item) GetValue() interface{} {
	if i == nil {
		return nil
	}

	return i.value
}

// GetIndex returns the index of this item, or 0 if the item is invalid.
func (i *Item) GetIndex() int {
	if i == nil {
		return 0
	}

	return i.index
}

// SetValue sets the item's value.
func (i *Item) SetValue(value interface{}) error {
	if i == nil {
		return errBadItem
	}

	i.value = value

	return nil
}

// SetIndex sets the item's index.
func (i *Item) SetIndex(index int) error {
	if i == nil {
		return errBadItem
	}

	i.index = index

	return nil
}

// findNode will iterate down a tree until it finds a matching index. If no matching index is found, then this returns
// nil for the node. Additionally, it builds a stack of all the nodes traversed along the way.
func (n *tnode) findNode(index int) (*tnode, *hstack.Stack) {
	stack := hstack.New()

	node := n
	for node != nil {
		if index == node.index() {
			break
		}

		stack.Add(node)
		if index < node.index() {
			node = node.left
		} else {
			node = node.right
		}
	}

	return node, stack
}

// rebalance calculates the balances of the nodes in the path and performs any necessary rotation operations to
// rebalance the tree.
func (t *Tree) rebalance(stack *hstack.Stack, index int, added bool) {
	for stack.Count() > 0 {
		node := stack.Pop().(*tnode)

		bal := node.balance()
		if (added && bal == 0) || (!added && (bal == -1 || bal == 1)) {
			// The operation did not change the length of the longest sub-branch. We can stop checking for imbalance now.
			break
		}

		if (added && (bal == -1 || bal == 1)) || (!added && bal == 0) {
			// The longest sub-branch below this node is now one node longer/shorter. We'll change the height of this
			// sub-branch by one and keep moving up the insertion/deletion path.
			if added {
				node.height++
			} else {
				node.height--
			}
			continue
		}

		// We have an imbalance. Rotate the nodes to fix this. This will change the root node of this sub-branch, so
		// we'll need to link it back in after the rotation operation is done.
		rotated := rotate(node, index)
		if stack.Count() == 0 {
			// We're at the top of the tree.
			t.root = rotated
		} else {
			node = stack.Pop().(*tnode)
			if index < node.index() {
				node.left = rotated
			} else {
				node.right = rotated
			}
		}

		// If we had to rebalance after adding a node, then the tree is now correct and we can stop.
		if added {
			break
		}
	}
}

// rotate performs the necessary rotations to rebalance the tree from this node down.
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

	if index < top.index() {
		left = true
		bottom = top.left
		if index > bottom.index() {
			double = true
		}
	} else {
		left = false
		bottom = top.right
		if index < bottom.index() {
			double = true
		}
	}

	if double {
		// The insertion path is on different sides of the top and bottom nodes, so we have to do a double rotation.
		// We'll do the unique part first here, and then we'll do the common part below.
		oldBottom := bottom
		if left {
			newBottom := oldBottom.right
			top.left = newBottom
			oldBottom.right = newBottom.left
			newBottom.left = oldBottom
			bottom = newBottom
		} else {
			newBottom := oldBottom.left
			top.right = newBottom
			oldBottom.left = newBottom.right
			newBottom.right = oldBottom
			bottom = newBottom
		}
		updateHeight(oldBottom)
	}

	// Now, we'll do the common rotation on the top node that all balance operations will need.
	if left {
		top.left = bottom.right
		bottom.right = top
	} else {
		top.right = bottom.left
		bottom.left = top
	}
	updateHeight(top)
	updateHeight(bottom)

	// Pass the new top of this branch of the tree back to the caller for proper linking.
	return bottom
}

// updateHeight recalculates and sets the node's height.
func updateHeight(n *tnode) {
	if n != nil {
		leftCount := n.leftCount()
		rightCount := n.rightCount()

		if leftCount > rightCount {
			n.height = leftCount + 1
		} else {
			n.height = rightCount + 1
		}
	}
}
