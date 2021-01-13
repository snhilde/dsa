package htree

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"
)

// --- TREE TESTS ---

// Test creating a new tree object.
func TestNew(t *testing.T) {
	tr := New()

	if tr.root != nil {
		t.Error("root node is not nil")
	}

	if tr.count != 0 {
		t.Error("Tree claims to have nodes")
	}
}

// Test using bad tree objects with all the methods.
func TestBad(t *testing.T) {
	var tr *Tree

	if err := tr.Add(5, 5); err == nil {
		t.Error("Bad object test: Unexpectedly passed Add")
	}

	item := NewItem(5, 5)
	if err := tr.AddItems(item); err == nil {
		t.Error("Bad object test: Unexpectedly passed AddItems")
	}

	// Make sure it doesn't crash anything.
	tr.Remove(5)

	// Make sure it doesn't crash anything.
	tr.Clear()

	if item := tr.Item(5); item != (Item{}) {
		t.Error("Bad object test: Unexpectedly passed Item")
	}

	if v := tr.Value(5); v != nil {
		t.Error("Bad object test: Unexpectedly passed Value")
	}

	if ok := tr.Match(5); ok {
		t.Error("Bad object test: Unexpectedly passed Match")
	}

	if item := tr.Min(); item != (Item{}) {
		t.Error("Bad object test: Unexpectedly passed Min")
	}

	if item := tr.Max(); item != (Item{}) {
		t.Error("Bad object test: Unexpectedly passed Max")
	}

	if ch := tr.Yield(nil); ch != nil {
		t.Error("Bad object test: Unexpectedly passed Yield")
	}

	if l := tr.List(); l != nil {
		t.Error("Bad object test: Unexpectedly passed List")
	}

	if l := tr.DFS(); l != nil {
		t.Error("Bad object test: Unexpectedly passed DFS")
	}

	if l := tr.BFS(); l != nil {
		t.Error("Bad object test: Unexpectedly passed BFS")
	}

	if values := tr.LeftCenterRight(); values != nil {
		t.Error("Bad object test: Unexpectedly passed LeftCenterRight")
	}

	if values := tr.CenterLeftRight(); values != nil {
		t.Error("Bad object test: Unexpectedly passed CenterLeftRight")
	}

	if values := tr.LeftRightCenter(); values != nil {
		t.Error("Bad object test: Unexpectedly passed LeftRightCenter")
	}

	if values := tr.RightCenterLeft(); values != nil {
		t.Error("Bad object test: Unexpectedly passed RightCenterLeft")
	}

	if values := tr.CenterRightLeft(); values != nil {
		t.Error("Bad object test: Unexpectedly passed CenterRightLeft")
	}

	if values := tr.RightLeftCenter(); values != nil {
		t.Error("Bad object test: Unexpectedly passed RightLeftCenter")
	}

	if tr.String() != "<nil>" {
		t.Error("Bad object test: Unexpectedly passed String")
	}

	if tr.Count() != 0 {
		t.Error("Bad object test: Unexpectedly passed Count")
	}

	if tr.Height() != 0 {
		t.Error("Bad object test: Unexpectedly passed Height")
	}
}

func TestAdd(t *testing.T) {
	tr := New()

	// Do a few simple, hand-built tests to make sure things look right.
	if err := tr.Add(5, 5); err != nil {
		t.Error(err)
	}
	testString(t, tr, "5")
	testCount(t, tr, 1)
	testHeightBalance(t, tr)

	if err := tr.Add(10, 10); err != nil {
		t.Error(err)
	}
	testString(t, tr, "5, 10")
	testCount(t, tr, 2)
	testHeightBalance(t, tr)

	if err := tr.Add(1, 1); err != nil {
		t.Error(err)
	}
	testString(t, tr, "1, 5, 10")
	testCount(t, tr, 3)
	testHeightBalance(t, tr)

	// Now do a larger test to make sure items are inserted in the correct order.
	tr.Clear()
	_, items := buildMiscTree(100000)
	for _, item := range items {
		value := item.GetValue()
		index := item.GetIndex()
		if err := tr.Add(value, index); err != nil {
			t.Error(err)
		}
	}
	testSort(t, tr, items)
	testCount(t, tr, 100000)
	testHeightBalance(t, tr)
}

func TestAddItems(t *testing.T) {
	tr := New()

	// Do a few simple, hand-built tests to make sure things look right.
	item1 := NewItem(5, 5)
	if err := tr.AddItems(item1); err != nil {
		t.Error(err)
	}
	testString(t, tr, "5")
	testCount(t, tr, 1)
	testHeightBalance(t, tr)

	item2 := NewItem(10, 10)
	item3 := NewItem(1, 1)
	if err := tr.AddItems(item2, item3); err != nil {
		t.Error(err)
	}
	testString(t, tr, "1, 5, 10")
	testCount(t, tr, 3)
	testHeightBalance(t, tr)

	// Now do a larger test to make sure items are inserted in the correct order.
	tr.Clear()
	_, items := buildMiscTree(100000)
	if err := tr.AddItems(items...); err != nil {
		t.Error(err)
	}
	testSort(t, tr, items)
	testCount(t, tr, 100000)
	testHeightBalance(t, tr)

	// Test that setting a new value for an item doesn't affect the tree's value until the item is added to the tree
	// again. We're going to get a value from the tree, change its value, and then grab it again to make sure nothing's
	// changed. After that, we're going to add the item again at the same index and then grab it again to make sure the
	// value has been updated.
	tr.Clear()
	tr, items = buildMiscTree(500)
	testCount(t, tr, 500)
	testHeightBalance(t, tr)

	for _, v := range items {
		index := v.GetIndex()
		item := tr.Item(index)
		if item == (Item{}) {
			t.Error("Bad item")
			return
		}
		item.SetValue("new value")
		item = tr.Item(index)
		value := item.GetValue()
		if val, ok := value.(string); ok && val == "new value" {
			t.Error("Item in tree has been unexpectedly updated")
			return
		}

		item.SetValue("new value")
		tr.AddItems(item)
		item = tr.Item(index)
		if value := item.GetValue(); value.(string) != "new value" {
			t.Error("Item's value was not updated")
			return
		}
	}
}

func TestBalance(t *testing.T) {
	tr := New()

	// By adding numbers from low to high, we're only going to be performing single left rotations during the
	// rebalances. This will test specifically that single left rotations properly rebalance the branch. We're going to
	// test the balance after every addition.
	for i := 1; i < 1000; i++ {
		tr.Add(i, i)
		testCount(t, tr, i)
		testHeightBalance(t, tr)
	}

	// By adding numbers from high to low, we're only going to be performing single right rotations during the
	// rebalances. This will test specifically that single right rotations properly rebalance the branch. We're going to
	// test the balance after every addition.
	tr.Clear()
	for i := 1000; i > 0; i-- {
		tr.Add(i, i)
		testCount(t, tr, 1001-i)
		testHeightBalance(t, tr)
	}

	// Now let's run through trees of increasing size to make sure that all sizes within the range are properly balanced
	// and have the correct height at each node.
	for i := 1; i < 1000; i++ {
		tr, _ := buildNumTree(i, true)
		testHeightBalance(t, tr)
	}
}

func TestRemove(t *testing.T) {
	// Make sure that a missing index has no effect on the tree.
	count := 100
	tr, _ := buildNumTree(count, false)
	for i := count; i < count*2; i++ {
		tr.Remove(count)
		testCount(t, tr, count)
		testHeightBalance(t, tr)
	}

	// We'll use this tree to test the ability to remove a node without children.
	tr.Clear()
	count = 15

	// First in one direction.
	for i := 1; i <= count; i++ {
		tr.Add(i, i)
	}
	leaves := []int{1, 3, 5, 7, 9, 11, 13, 15}
	for i, leaf := range leaves {
		tr.Remove(leaf)

		// Make sure the node was actually removed.
		if tr.Match(leaf) {
			t.Error("Did not remove", leaf)
		}

		// Make sure that the tree is still in good shape.
		testCount(t, tr, count-1-i)
		testHeightBalance(t, tr)
	}

	// And then in the other direction.
	tr.Clear()
	for i := 1; i <= count; i++ {
		tr.Add(i, i)
	}
	for i := range leaves {
		index := leaves[len(leaves)-1-i]
		tr.Remove(index)

		// Make sure the node was actually removed.
		if tr.Match(index) {
			t.Error("Did not remove", index)
		}

		// Make sure that the tree is still in good shape.
		testCount(t, tr, count-1-i)
		testHeightBalance(t, tr)
	}

	// Let's also test the ability to remove a root node with no children.
	tr.Clear()
	tr.Add(1, 1)
	testCount(t, tr, 1)
	testHeightBalance(t, tr)

	tr.Remove(1)
	testCount(t, tr, 0)
	testHeightBalance(t, tr)

	// We'll use this tree to test the ability to remove a node with one child.
	tr.Clear()
	indexes := []int{3, 2, 4, 1, 5}
	for _, index := range indexes {
		tr.Add(index, index)
	}
	testCount(t, tr, len(indexes))
	testHeightBalance(t, tr)

	tr.Remove(2)
	testCount(t, tr, len(indexes)-1)
	testHeightBalance(t, tr)
	if tr.Match(2) {
		t.Error("Did not remove 2")
	}

	tr.Remove(4)
	testCount(t, tr, len(indexes)-2)
	testHeightBalance(t, tr)
	if tr.Match(4) {
		t.Error("Did not remove 4")
	}

	// Let's also test the ability to remove a root node with one child.
	tr.Clear()
	tr.Add(1, 1)
	tr.Add(2, 2)
	testCount(t, tr, 2)
	testHeightBalance(t, tr)

	tr.Remove(1)
	testCount(t, tr, 1)
	testHeightBalance(t, tr)
	if tr.Match(1) {
		t.Error("Did not remove 1")
	}

	// We'll use this tree to test the ability to remove a node with two children.
	tr.Clear()
	count = 15

	// First in one direction.
	for i := 1; i <= count; i++ {
		tr.Add(i, i)
	}
	parents := []int{2, 6, 10, 14, 4, 12}
	for i, parent := range parents {
		tr.Remove(parent)

		// Make sure the node was actually removed.
		if tr.Match(parent) {
			t.Error("Did not remove", parent)
		}

		// Make sure that the tree is still in good shape.
		testCount(t, tr, count-1-i)
		testHeightBalance(t, tr)
	}

	// And then in the other direction.
	tr.Clear()
	for i := 1; i <= count; i++ {
		tr.Add(i, i)
	}
	for i := range parents {
		index := parents[len(parents)-1-i]
		tr.Remove(index)

		// Make sure the node was actually removed.
		if tr.Match(index) {
			t.Error("Did not remove", index)
		}

		// Make sure that the tree is still in good shape.
		testCount(t, tr, count-1-i)
		testHeightBalance(t, tr)
	}

	// Let's also test the ability to remove a root node with two children.
	tr.Clear()
	tr.Add(2, 2)
	tr.Add(1, 1)
	tr.Add(3, 3)
	testCount(t, tr, 3)
	testString(t, tr, "1, 2, 3")
	testHeightBalance(t, tr)

	tr.Remove(2)
	testCount(t, tr, 2)
	testString(t, tr, "1, 3")
	testHeightBalance(t, tr)

	// And a root node with two children that will need rebalancing after removal.
	tr.Clear()
	tr.Add(3, 3)
	tr.Add(2, 2)
	tr.Add(4, 4)
	tr.Add(1, 1)
	testCount(t, tr, 4)
	testString(t, tr, "1, 2, 3, 4")
	testHeightBalance(t, tr)

	tr.Remove(4)
	testCount(t, tr, 3)
	testString(t, tr, "1, 2, 3")
	testHeightBalance(t, tr)

	// Now that everything looks good, let's set up a large table with non-sequential indexes and remove all nodes from
	// it. We're going to put all the indexes into a hash table so we can have random look up when deciding which node
	// to remove next.
	count = 10000
	indexMap := make(map[int]bool, count)
	tr, items := buildMiscTree(count)
	for _, item := range items {
		indexMap[item.GetIndex()] = true
	}

	for index := range indexMap {
		tr.Remove(index)

		// Make sure the node was actually removed.
		if tr.Value(index) != nil {
			t.Error("Did not remove", index)
		}
		count--

		// Make sure that the tree is still in good shape.
		testCount(t, tr, count)
		testHeightBalance(t, tr)
	}
}

func TestClear(t *testing.T) {
	tr := New()

	// Do a few simple, hand-built tests to make sure things look right.
	item1 := NewItem(5, 5)
	if err := tr.AddItems(item1); err != nil {
		t.Error(err)
	}
	testString(t, tr, "5")
	testCount(t, tr, 1)
	testHeightBalance(t, tr)

	tr.Clear()
	testString(t, tr, "<empty>")
	testCount(t, tr, 0)
	testHeightBalance(t, tr)

	// Add 500 items of various types.
	tr, _ = buildMiscTree(500)
	testCount(t, tr, 500)

	tr.Clear()
	testString(t, tr, "<empty>")
	testCount(t, tr, 0)
	testHeightBalance(t, tr)
}

func TestItem(t *testing.T) {
	tr := New()

	// Do a few simple, hand-built tests to make sure things look right.
	tr.Add(5, 5)
	tr.Add(10, 10)
	tr.Add(1, 1)
	if item := tr.Item(5); item.GetValue() != 5 {
		t.Error("Expected 5, Received", item.GetValue())
	}
	if item := tr.Item(10); item.GetValue() != 10 {
		t.Error("Expected 10, Received", item.GetValue())
	}
	if item := tr.Item(1); item.GetValue() != 1 {
		t.Error("Expected 1, Received", item.GetValue())
	}

	// Test that some other values are not present.
	if item := tr.Item(2); item != (Item{}) {
		t.Error("Expected nothing, Received", item.GetValue())
	}
	if item := tr.Item(20); item != (Item{}) {
		t.Error("Expected nothing, Received", item.GetValue())
	}
	if item := tr.Item(100); item != (Item{}) {
		t.Error("Expected nothing, Received", item.GetValue())
	}

	// Now do a larger test to make sure the correct item is returned.
	tr, items := buildMiscTree(100000)
	for _, item := range items {
		value := item.GetValue()
		index := item.GetIndex()
		treeItem := tr.Item(index)
		if !reflect.DeepEqual(treeItem.GetValue(), value) {
			t.Error("Wrong value")
			t.Log("Expected:", value)
			t.Log("Received:", treeItem.GetValue())
		} else if treeItem.GetIndex() != index {
			t.Error("Wrong index")
			t.Log("Expected:", index)
			t.Log("Received:", treeItem.GetIndex())
		}
	}
}

func TestValue(t *testing.T) {
	tr := New()

	// Do a few simple, hand-built tests to make sure things look right.
	tr.Add(5, 5)
	tr.Add(10, 10)
	tr.Add(1, 1)
	if v := tr.Value(5); v != 5 {
		t.Error("Expected 5, Received", v)
	}
	if v := tr.Value(10); v != 10 {
		t.Error("Expected 10, Received", v)
	}
	if v := tr.Value(1); v != 1 {
		t.Error("Expected 1, Received", v)
	}

	// Test that some other values are not present.
	if v := tr.Value(2); v != nil {
		t.Error("Expected nothing, Received", v)
	}
	if v := tr.Value(20); v != nil {
		t.Error("Expected nothing, Received", v)
	}
	if v := tr.Value(100); v != nil {
		t.Error("Expected nothing, Received", v)
	}

	// Now do a larger test to make sure indexes and values are properly tied and look-up is correct.
	tr, items := buildMiscTree(100000)
	for _, item := range items {
		value := item.GetValue()
		index := item.GetIndex()
		treeItem := tr.Item(index)
		if !reflect.DeepEqual(treeItem.GetValue(), value) {
			t.Error("Wrong value")
			t.Log("Expected:", value)
			t.Log("Received:", treeItem.GetValue())
		}
	}
}

func TestMatch(t *testing.T) {
	tr := New()
	r := newRand()

	// Make two lists of 500 items each. One list will be added to the tree, and the other won't. We'll then check that
	// the added ones do match and the not-added ones don't match.
	presentItems := make([]Item, 500)
	absentItems := make([]Item, 500)
	for i := 0; i < 1000; i++ {
		value := r.Int()
		item := NewItem(value, value)
		index := i / 2
		if i%2 == 0 {
			presentItems[index] = item
		} else {
			absentItems[index] = item
		}
	}

	if err := tr.AddItems(presentItems...); err != nil {
		t.Error(err)
		return
	}
	testCount(t, tr, 500)
	testHeightBalance(t, tr)

	// Make sure that all of these items match.
	for i, item := range presentItems {
		if !tr.Match(item.GetValue()) {
			t.Error("Missing item at index", i)
		}
	}

	// Make sure that none of these items match. (There's a very low probablity that we'll have a value collision with
	// r.Int(), and we can just run the tests again if we do).
	for i, item := range absentItems {
		if tr.Match(item.GetValue()) {
			t.Error("Unexpected item at index", i)
		}
	}

	// Test items with different types.
	tr.Clear()
	_, items := buildMiscTree(1000)
	for i, item := range items {
		index := i / 2
		if i%2 == 0 {
			presentItems[index] = item
		} else {
			absentItems[index] = item
		}
	}

	if err := tr.AddItems(presentItems...); err != nil {
		t.Error(err)
		return
	}
	testSort(t, tr, presentItems)
	testCount(t, tr, 500)
	testHeightBalance(t, tr)

	// Make sure that all of these items match.
	for i, item := range presentItems {
		if !tr.Match(item.GetValue()) {
			t.Error("Missing item at index", i)
		}
	}

	// Make sure that none of these items match. (There's a very low probablity that we'll have a value collision with
	// r.Int(), and we can just run the tests again if we do).
	for i, item := range absentItems {
		if tr.Match(item.GetValue()) {
			t.Error("Unexpected item at index", i)
		}
	}
}

func TestMin(t *testing.T) {
	count := 500
	tr, items := buildMiscTree(count)
	testSort(t, tr, items)
	testCount(t, tr, count)

	item := items[0]
	if min := tr.Min(); !reflect.DeepEqual(min, item) {
		t.Error("Minimum items differ")
		t.Log("Expected:", item)
		t.Log("Received:", min)
	}

	// Add some items and keep checking
	_, additions := buildMiscTree(count)
	for i, addition := range additions {
		tr.AddItems(addition)
		items = append(items, addition)
		testSort(t, tr, items)
		testCount(t, tr, count+i+1)
		testHeightBalance(t, tr)

		item := items[0]
		if min := tr.Min(); !reflect.DeepEqual(min, item) {
			t.Error("Minimum items differ")
			t.Log("Expected:", item)
			t.Log("Received:", min)
		}
	}

	// Remove all of the min nodes while continuosly checking.
	for i := 0; i < len(items)-1; i++ {
		item := items[i]
		tr.Remove(item.GetIndex())

		item = items[i+1]
		if min := tr.Min(); !reflect.DeepEqual(min, item) {
			t.Error("Minimum items differ", item)
			t.Log("Expected:", item)
			t.Log("Received:", min)
		}
	}
}

func TestMax(t *testing.T) {
	count := 500
	tr, items := buildMiscTree(count)
	testSort(t, tr, items)
	testCount(t, tr, count)

	item := items[len(items)-1]
	if max := tr.Max(); !reflect.DeepEqual(max, item) {
		t.Error("Maximum items differ")
		t.Log("Expected:", item)
		t.Log("Received:", max)
	}

	// Add some items and keep checking
	_, additions := buildMiscTree(count)
	for i, addition := range additions {
		tr.AddItems(addition)
		items = append(items, addition)
		testSort(t, tr, items)
		testCount(t, tr, count+i+1)
		testHeightBalance(t, tr)

		item := items[len(items)-1]
		if max := tr.Max(); !reflect.DeepEqual(max, item) {
			t.Error("Maximum items differ")
			t.Log("Expected:", item)
			t.Log("Received:", max)
		}
	}

	// Remove all of the max nodes while continuosly checking.
	for i := len(items) - 1; i > 0; i-- {
		item := items[i]
		tr.Remove(item.GetIndex())

		item = items[i-1]
		if max := tr.Max(); !reflect.DeepEqual(max, item) {
			t.Error("Maximum items differ", item)
			t.Log("Expected:", item)
			t.Log("Received:", max)
		}
	}
}

func TestYield(t *testing.T) {
	tr := New()
	tr.Add(1, 1)

	// Test that the capacity of the quit channel doesn't matter.
	if ch := tr.Yield(nil); ch == nil {
		t.Error("Received nil channel with nil quit")
	}
	quit := make(chan interface{}, 0)
	if ch := tr.Yield(quit); ch == nil {
		t.Error("Received nil channel with quit with a buffer of 0")
	}
	quit = make(chan interface{}, 1)
	if ch := tr.Yield(quit); ch == nil {
		t.Error("Received nil channel with quit with a buffer of 1")
	}
	quit = make(chan interface{}, 10)
	if ch := tr.Yield(quit); ch == nil {
		t.Error("Received nil channel with quit with a buffer of 10")
	}

	// Test that indexes are sorted and the correct items are returned.
	tr, items := buildMiscTree(500)
	testSort(t, tr, items)
	testCount(t, tr, 500)
	testHeightBalance(t, tr)

	yieldChan := tr.Yield(nil)
	if yieldChan == nil {
		t.Error("Yield returned nil channel")
		return
	}

	i := 0
	for item := range yieldChan {
		if !reflect.DeepEqual(items[i].GetValue(), item.GetValue()) {
			t.Error("Items differ at index", i)
			t.Log("Expected:", items[i].GetValue())
			t.Log("Received:", item.GetValue())
		}
		i++
	}

	// Test that sending on quit channel also closes the yield channel.
	quit = make(chan interface{}, 0)
	yieldChan = tr.Yield(quit)
	if yieldChan == nil {
		t.Error("Yield returned nil channel")
		return
	}
	testCount(t, tr, 500)
	testHeightBalance(t, tr)

	// Grab the first two items.
	for i := 0; i < 2; i++ {
		<-yieldChan
	}

	// Send on the channel, grab the last item, and make sure that we can't grab another item.
	quit <- struct{}{}
	select {
	case <-yieldChan:
	default:
		t.Error("Did not receive last item")
	}
	select {
	case <-yieldChan:
		t.Error("Unexpectedly received item 4")
	default:
	}
}

func TestList(t *testing.T) {
	// Make sure an empty tree returns nil.
	tr := New()
	if items := tr.List(); items != nil {
		t.Error("Should have received nothing")
	}

	// Make sure that the items are returned in sorted order.
	tr, items := buildMiscTree(1000)
	testSort(t, tr, items)
	testCount(t, tr, 1000)
	testHeightBalance(t, tr)

	listItems := tr.List()
	if len(listItems) != 1000 {
		t.Error("List did not return all items")
		return
	}

	for i, item := range listItems {
		if !reflect.DeepEqual(items[i].GetValue(), item.GetValue()) {
			t.Error("Items differ at index", i)
		}
	}
}

func TestDFS(t *testing.T) {
	// Make sure an empty tree returns nil.
	tr := New()
	if items := tr.DFS(); items != nil {
		t.Error("Should have received nothing")
	}

	// Make sure that the items are returned in sorted order.
	tr, items := buildMiscTree(1000)
	testSort(t, tr, items)
	testCount(t, tr, 1000)
	testHeightBalance(t, tr)

	traversed := tr.DFS()
	if len(traversed) != 1000 {
		t.Error("DFS did not return all items")
		return
	}

	for i, item := range traversed {
		if !reflect.DeepEqual(item, items[i].GetValue()) {
			t.Error("Items differ at index", i)
		}
	}
}

func TestBFS(t *testing.T) {
	// Make sure an empty tree returns nil.
	tr := New()
	if values := tr.BFS(); values != nil {
		t.Error("Should have received nothing")
	}

	// Do a quick, easily provable test to make sure that the basics are not broken.
	for i := 1; i <= 11; i++ {
		tr.Add(i, i)
	}
	testCount(t, tr, 11)
	testString(t, tr, "1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11")

	knownGood := []interface{}{4, 2, 8, 1, 3, 6, 10, 5, 7, 9, 11}
	if values := tr.BFS(); !reflect.DeepEqual(values, knownGood) {
		t.Error("BFS failed easy test")
		return
	}
	return

	// We can use this simple formula to manually build the correct values for a tree whose count is one less than a
	// power of 2:
	// 1. The first level of the tree is the power of 2 down minus 1.
	// 2. Build out every level after that by going through the nodes from the previous level and adding and subtracting
	//    the current step, which is decremented every level to the next power of 2 down.
	// 3. Add each level to the master list.
	count := 65535
	root := count / 2
	priorLevel := []int{root}
	set1 := priorLevel
	for step := (root + 1) / 2; step > 0; step /= 2 {
		currentLevel := make([]int, 0)
		for _, v := range priorLevel {
			currentLevel = append(currentLevel, v-step)
			currentLevel = append(currentLevel, v+step)
		}
		set1 = append(set1, currentLevel...)
		priorLevel = currentLevel
	}

	// Now that we have the values, let's build out the tree.
	tr, _ = buildNumTree(count, false)
	testCount(t, tr, count)
	testHeightBalance(t, tr)

	// Check that both sets of values are equal.
	set2 := tr.BFS()
	if len(set1) != len(set2) {
		t.Error("BFS sets have different lengths")
		t.Log("Set 1:", len(set1))
		t.Log("Set 2:", len(set2))
		return
	}
	for i, v := range set1 {
		if v != set2[i].(int) {
			t.Error("BFS failed manual test at index", i)
		}
	}
}

func TestLeftCenterRight(t *testing.T) {
	testTraversal(t, (*Tree).LeftCenterRight, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
}

func TestCenterLeftRight(t *testing.T) {
	testTraversal(t, (*Tree).CenterLeftRight, []int{7, 3, 1, 0, 2, 5, 4, 6, 11, 9, 8, 10, 13, 12, 14, 15})
}

func TestLeftRightCenter(t *testing.T) {
	testTraversal(t, (*Tree).LeftRightCenter, []int{0, 2, 1, 4, 6, 5, 3, 8, 10, 9, 12, 15, 14, 13, 11, 7})
}

func TestRightCenterLeft(t *testing.T) {
	testTraversal(t, (*Tree).RightCenterLeft, []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
}

func TestCenterRightLeft(t *testing.T) {
	testTraversal(t, (*Tree).CenterRightLeft, []int{7, 11, 13, 14, 15, 12, 9, 10, 8, 3, 5, 6, 4, 1, 2, 0})
}

func TestRightLeftCenter(t *testing.T) {
	testTraversal(t, (*Tree).RightLeftCenter, []int{15, 14, 12, 13, 10, 8, 9, 11, 6, 4, 5, 2, 0, 1, 3, 7})
}

// --- ITEM TESTS ---

func TestNewItem(t *testing.T) {
	item := NewItem("a", 1)
	if item == (Item{}) {
		t.Error("New item is empty")
	}

	// Test that default values yields an empty object.
	if item := NewItem(nil, 0); item != (Item{}) {
		t.Error("Default item is not empty")
	}
}

func TestBadItem(t *testing.T) {
	var item *Item

	// GetValue
	if value := item.GetValue(); value != nil {
		t.Error("Unexpectedly passed bad item test for GetValue")
	}

	// SetValue
	if err := item.SetValue(5); err == nil {
		t.Error("Unexpectedly passed bad item test for SetValue")
	}
}

func TestGetValue(t *testing.T) {
	// Make 500 items by hand, and test that they have the proper values.
	values := buildValues(500)
	items := make([]Item, 500)
	for i := range items {
		items[i] = NewItem(values[i], i)
	}

	for i, item := range items {
		if value := item.GetValue(); !reflect.DeepEqual(values[i], value) {
			t.Error("Item", i, "returned the wrong value")
			t.Error("Expected:", values[i])
			t.Error("Received:", value)
		}
	}

	// Test that we can get an empty value.
	item := NewItem(nil, 0)
	if value := item.GetValue(); value != nil {
		t.Error("Empty item did not return nil value")
	}

	// Test that we can build a tree and get the same item values.
	tr, items := buildMiscTree(500)
	testSort(t, tr, items)
	testCount(t, tr, 500)
	for i, v := range items {
		index := v.GetIndex()
		item := tr.Item(index)
		if item == (Item{}) {
			t.Error("Invalid item at index", i)
		} else {
			val1 := v.GetValue()
			val2 := item.GetValue()
			if !reflect.DeepEqual(val1, val2) {
				t.Error("Item's value is different than tree's value")
				t.Error("Expected:", val1)
				t.Error("Received:", val2)
			}
		}
	}
}

func TestGetIndex(t *testing.T) {
	// Make 500 items by hand, and test that they have the proper indexes.
	items := make([]Item, 500)
	for i := range items {
		items[i] = NewItem(i, i)
	}

	for i, item := range items {
		if item.GetIndex() != i {
			t.Error("Item", i, "returned the wrong index")
			t.Error("Expected:", i)
			t.Error("Received:", item.GetIndex())
		}
	}

	// Test that we can build a tree and get the same item indexes.
	tr, items := buildMiscTree(500)
	testSort(t, tr, items)
	testCount(t, tr, 500)
	for i, v := range items {
		index := v.GetIndex()
		item := tr.Item(index)
		if item == (Item{}) {
			t.Error("Invalid item at index", i)
		} else {
			if item.GetIndex() != index {
				t.Error("Item's index is different than tree's index")
				t.Error("Expected:", index)
				t.Error("Received:", item.GetIndex())
			}
		}
	}
}

func TestSetValue(t *testing.T) {
	// Test that a new value is correctly reflected in the item.
	values := buildValues(500)
	items := make([]Item, 500)
	for i := range items {
		items[i] = NewItem(values[i], i)
	}

	for i, item := range items {
		if value := item.GetValue(); !reflect.DeepEqual(values[i], value) {
			t.Error("Item", i, "returned the wrong value")
			t.Error("Expected:", values[i])
			t.Error("Received:", value)
		}
	}

	newValues := buildValues(500)
	for i, item := range items {
		if err := item.SetValue(newValues[i]); err != nil {
			t.Error(err)
		}
		// Save the modified item.
		items[i] = item
	}

	for i, item := range items {
		if value := item.GetValue(); !reflect.DeepEqual(newValues[i], value) {
			t.Error("Item", i, "returned the wrong new value")
			t.Error("Expected:", newValues[i])
			t.Error("Received:", value)
		}
	}

	// Make sure that setting a new value for an item doesn't affect the item's value in the tree.
	tr, items := buildMiscTree(1000)
	testSort(t, tr, items)
	testCount(t, tr, 1000)

	for _, v := range items {
		index := v.GetIndex()
		item := tr.Item(index)
		if item == (Item{}) {
			t.Error("Bad item")
			continue
		}
		item.SetValue("new value")
		item = tr.Item(index)
		value := item.GetValue()
		if val, ok := value.(string); ok && val == "new value" {
			t.Error("Item in tree has been unexpectedly updated")
			continue
		}

		item.SetValue("new value")
		tr.AddItems(item)
		item = tr.Item(index)
		if value := item.GetValue(); value.(string) != "new value" {
			t.Error("Item's value was not updated")
			continue
		}
	}
}

func TestSetIndex(t *testing.T) {
	// Test that a new index is correctly reflected in the item.
	items := make([]Item, 500)
	for i := range items {
		items[i] = NewItem(i, i)
	}

	for i, item := range items {
		if item.GetIndex() != i {
			t.Error("Item", i, "returned the wrong index")
			t.Error("Expected:", i)
			t.Error("Received:", item.GetIndex())
		}
	}

	for i, item := range items {
		if err := item.SetIndex(i + 1000); err != nil {
			t.Error(err)
		}
		// Save the modified item.
		items[i] = item
	}

	for i, item := range items {
		if item.GetIndex() != i+1000 {
			t.Error("Item", i, "returned the wrong new index")
			t.Error("Expected:", i+1000)
			t.Error("Received:", item.GetIndex())
		}
	}

	// Make sure that setting a new index for an item doesn't affect the item's index in the tree.
	tr, items := buildMiscTree(1000)
	testSort(t, tr, items)
	testCount(t, tr, 1000)
	testHeightBalance(t, tr)

	newItems := make([]Item, len(items))
	for i, v := range items {
		index := v.GetIndex()
		newIndex := i
		item := tr.Item(index)
		if item == (Item{}) {
			t.Error("Bad item")
			continue
		}
		item.SetIndex(newIndex)
		item = tr.Item(index)
		if item.GetIndex() == newIndex {
			t.Error("Item in tree has been unexpectedly updated")
			continue
		}

		item.SetIndex(newIndex)
		newItems[i] = item
		tr.AddItems(item)

		item = tr.Item(newIndex)
		if item == (Item{}) {
			t.Error("Bad item at new index")
		} else if item.GetIndex() != newIndex {
			t.Error("Item's index was not updated")
		}
	}
	items = append(items, newItems...)
	testSort(t, tr, items)
	testCount(t, tr, 2000)
	testHeightBalance(t, tr)
}

// --- TREE BENCHMARKS ---

func Benchmark100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = buildNumTree(100, false)
	}
}

func Benchmark1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = buildNumTree(1000, false)
	}
}

func Benchmark10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = buildNumTree(10000, false)
	}
}

func Benchmark100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = buildNumTree(100000, false)
	}
}

func Benchmark1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = buildNumTree(1000000, false)
	}
}

func BenchmarkList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tr, _ := buildMiscTree(1000)
		_ = tr.List()
	}
}

func BenchmarkDFS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tr, _ := buildMiscTree(1000)
		_ = tr.DFS()
	}
}

// --- HELPER FUNCTIONS ---

// testString checks whether or not the string representation of the tree is the same as expected.
func testString(t *testing.T, tr Tree, want string) {
	s := tr.String()
	if s != want {
		t.Error("Expected:", want)
		t.Error("Received:", s)
	}
}

// testCount checks whether or not the tree has the expected number of nodes.
func testCount(t *testing.T, tr Tree, want int) {
	count := tr.Count()
	if count != want {
		t.Error("Want", want, "items")
		t.Error("Have", count, "items")
	}
}

// testSort sorts the provided items and then checks that the tree has the same order of items in it.
func testSort(t *testing.T, tr Tree, items []Item) {
	// Sort the list of items.
	sort.Slice(items, func(i, j int) bool {
		return items[i].index < items[j].index
	})

	// Build a string of values from the items.
	var b strings.Builder
	for _, item := range items {
		fmt.Fprintf(&b, "%v, ", item.GetValue())
	}
	s := strings.TrimSuffix(b.String(), ", ")

	// Check that the string is the same as the tree's string.
	testString(t, tr, s)
}

// testHeightBalance checks the height and balance of the tree.
func testHeightBalance(t *testing.T, tr Tree) {
	height := testBalance(t, tr.root)
	if height != tr.Height() {
		t.Error("Tree is reporting wrong height")
		t.Log("Expected", height)
		t.Log("Received", tr.Height())
	}
}

// testBalance checks whether or not all the nodes in the tree have the correct balance and returns the longest
// sub-branch at this node.
func testBalance(t *testing.T, node *tnode) int {
	if node == nil {
		return 0
	}

	// Make sure the indexes are in proper order.
	if node.left != nil && node.left.index() >= node.index() {
		t.Error(node.left.index(), "is not less than", node.index())
	}
	if node.right != nil && node.right.index() <= node.index() {
		t.Error(node.right.index(), "is not less than", node.index())
	}

	// Make sure the balance is being correctly calculated.
	leftCount := testBalance(t, node.left)
	rightCount := testBalance(t, node.right)

	balance := rightCount - leftCount
	if balance != node.balance() {
		t.Error("Node at index", node.index(), "has wrong balance")
		t.Log("Should be:", balance)
		t.Log("Node says:", node.balance())
	}

	// Make sure the height is being correctly calculated.
	height := leftCount
	if rightCount > leftCount {
		height = rightCount
	}
	height++
	if height != node.height {
		t.Error("Node at index", node.index(), "has wrong height")
		t.Log("Should be:", height)
		t.Log("Node says:", node.height)
	}

	return height
}

// testTraversal runs through a tree traversal using the specified Tree method and checks that the output is the same as
// expected. This uses only int types.
func testTraversal(t *testing.T, method func(*Tree) []interface{}, expect []int) {
	tr, _ := buildNumTree(len(expect), false)
	testCount(t, tr, len(expect))

	values := method(&tr)

	if len(expect) != len(values) {
		t.Error("Lengths do not match")
		t.Log("Expected:", len(expect))
		t.Log("Received:", len(values))
	}

	for i, value := range values {
		if value != expect[i] {
			t.Error("Values do not match at index", i)
			t.Log("Expected:", expect[i])
			t.Log("Received:", value)
		}
	}
}

// newRand returns a new random number generator.
func newRand() *rand.Rand {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	random := rand.New(source)

	return random
}

// buildNumTree creates a new tree and populates it with count items, either randomly or by iterating from low to high.
// It returns the new tree as well as the indexes of all the items.
func buildNumTree(count int, random bool) (Tree, []int) {
	tr := New()
	indexes := make([]int, count)

	if random {
		r := newRand()
		for i := range indexes {
			// v := r.Int()
			v := int(r.Int31n(999))
			tr.Add(v, v)
			indexes[i] = v
		}
	} else {
		indexes = nil
		for i := 0; i < count; i++ {
			tr.Add(i, i)
		}
	}

	return tr, indexes
}

// buildMiscTree creates a new tree and populates it with count items with random values of various types. It returns
// the new tree as well as all the items.
func buildMiscTree(count int) (Tree, []Item) {
	if count < 1 {
		return Tree{}, nil
	}

	// Build out the items first.
	values := buildValues(count)
	if values == nil {
		return Tree{}, nil
	}

	// Make all the items.
	r := newRand()
	items := make([]Item, count)
	for i := 0; i < count; i++ {
		items[i] = NewItem(values[i], r.Int())
	}

	// Add the items to a new tree.
	t := New()
	if err := t.AddItems(items...); err != nil {
		return Tree{}, nil
	}
	return t, items
}

// buildValues build out a slice of random values of length count.
func buildValues(count int) []interface{} {
	if count < 1 {
		return nil
	}

	r := newRand()
	values := make([]interface{}, count)
	for i := 0; i < count; i++ {
		var value interface{}
		switch i % 6 {
		case 0:
			value = r.Int()
		case 1:
			value = r.Float64()
		case 2:
			value = r.Uint32()
		case 3:
			value = rune(r.Int31())
		case 4:
			value = string([]byte{byte(r.Int31n(94) + 32), byte(r.Int31n(94) + 32), byte(r.Int31n(94) + 32), byte(r.Int31n(94) + 32)})
		case 5:
			value = []int{r.Int(), r.Int(), r.Int(), r.Int(), r.Int(), r.Int(), r.Int()}
		}
		values[i] = value
	}

	return values
}
