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

	if tr.trunk != nil {
		t.Error("Trunk node is not nil")
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

	if item := tr.Remove(5); item != (Item{}) {
		t.Error("Bad object test: Unexpectedly passed Remove")
	}

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

	if ch := tr.Yield(nil); ch != nil {
		t.Error("Bad object test: Unexpectedly passed Yield")
	}

	if l := tr.List(); l != nil {
		t.Error("Bad object test: Unexpectedly passed List")
	}

	if tr.String() != "<nil>" {
		t.Error("Bad object test: Unexpectedly passed String")
	}

	if tr.Count() != -1 {
		t.Error("Bad object test: Unexpectedly passed Count")
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
	testBalance(t, tr.trunk)

	if err := tr.Add(10, 10); err != nil {
		t.Error(err)
	}
	testString(t, tr, "5, 10")
	testCount(t, tr, 2)
	testBalance(t, tr.trunk)

	if err := tr.Add(1, 1); err != nil {
		t.Error(err)
	}
	testString(t, tr, "1, 5, 10")
	testCount(t, tr, 3)
	testBalance(t, tr.trunk)

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
	testBalance(t, tr.trunk)
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
	testBalance(t, tr.trunk)

	item2 := NewItem(10, 10)
	item3 := NewItem(1, 1)
	if err := tr.AddItems(item2, item3); err != nil {
		t.Error(err)
	}
	testString(t, tr, "1, 5, 10")
	testCount(t, tr, 3)
	testBalance(t, tr.trunk)

	// Now do a larger test to make sure items are inserted in the correct order.
	tr.Clear()
	_, items := buildMiscTree(100000)
	if err := tr.AddItems(items...); err != nil {
		t.Error(err)
	}
	testSort(t, tr, items)
	testCount(t, tr, 100000)
	testBalance(t, tr.trunk)

	// Test that setting a new value for an item doesn't affect the tree's value until the item is added to the tree
	// again. We're going to get a value from the tree, change its value, and then grab it again to make sure nothing's
	// changed. After that, we're going to add the item again at the same index and then grab it again to make sure the
	// value has been updated.
	tr.Clear()
	tr, items = buildMiscTree(500)
	testCount(t, tr, 500)
	testBalance(t, tr.trunk)

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

func TestBalance(t *testing.T) {
	tr := New()

	// By adding numbers from low to high, we're only going to be performing single left rotations during the
	// rebalances. This will test specifically that single left rotations properly rebalance the branch. We're going to
	// test the balance after every addition.
	for i := 1; i < 1000; i++ {
		tr.Add(i, i)
		testCount(t, tr, i)
		testBalance(t, tr.trunk)
	}

	// By adding numbers from high to low, we're only going to be performing single right rotations during the
	// rebalances. This will test specifically that single right rotations properly rebalance the branch. We're going to
	// test the balance after every addition.
	tr.Clear()
	for i := 1000; i > 0; i-- {
		tr.Add(i, i)
		testCount(t, tr, 1001-i)
		testBalance(t, tr.trunk)
	}

	// Now let's run through trees of increasing size to make sure that all sizes within the range are properly balanced
	// and have the correct height at each node.
	for i := 1; i < 1000; i++ {
		tr, _ := buildNumTree(i, true)
		testBalance(t, tr.trunk)
	}
}

func TestRemove(t *testing.T) {
	// Build a random tree, and test that removing 10 random items one-by-one works as expected.
	count := 500
	tr, items := buildMiscTree(count)
	testSort(t, tr, items)
	testCount(t, tr, count)
	testBalance(t, tr.trunk)

	r := newRand()
	for i := 0; i < 10; i++ {
		index := int(r.Int31n(int32(count - 1)))
		if item := tr.Remove(items[index].GetIndex()); item == (Item{}) {
			t.Error("Failed to remove item", i)
			return
		}

		count--
		tmp := make([]Item, count)
		copy(tmp, items[:index])
		copy(tmp[index:], items[index+1:])
		items = tmp

		testSort(t, tr, items)
		testCount(t, tr, count)
	}

	// Test adding some items back in to make sure the tree can still balance itself.
	_, newItems := buildMiscTree(count)
	if err := tr.AddItems(newItems...); err != nil {
		t.Error(err)
		return
	}
	items = append(items, newItems...)
	testSort(t, tr, items)
	testCount(t, tr, len(items))
	testBalance(t, tr.trunk)
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
	testBalance(t, tr.trunk)

	tr.Clear()
	testString(t, tr, "<empty>")
	testCount(t, tr, 0)
	testBalance(t, tr.trunk)

	// Add 500 items of various types.
	tr, _ = buildMiscTree(500)
	testCount(t, tr, 500)

	tr.Clear()
	testString(t, tr, "<empty>")
	testCount(t, tr, 0)
	testBalance(t, tr.trunk)
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
	testBalance(t, tr.trunk)

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
	testBalance(t, tr.trunk)

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

func TestYield(t *testing.T) {
	tr := New()

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
	testBalance(t, tr.trunk)

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
	testBalance(t, tr.trunk)

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
	// Make sure that the items are returned in sorted order.
	tr, items := buildMiscTree(1000)
	testSort(t, tr, items)
	testCount(t, tr, 1000)
	testBalance(t, tr.trunk)

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
	testBalance(t, tr.trunk)

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
	testBalance(t, tr.trunk)
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

// testBalance checks whether or not all the nodes in the tree have the correct balance.
func testBalance(t *testing.T, node *tnode) int {
	if node == nil {
		return 0
	}

	leftCount := testBalance(t, node.left)
	rightCount := testBalance(t, node.right)

	balance := rightCount - leftCount
	if balance != node.balance() {
		t.Error("Node at index", node.item.index, "has wrong balance")
		t.Log("Should be:", balance)
		t.Log("Node says:", node.balance())
	}

	height := leftCount
	if rightCount > leftCount {
		height = rightCount
	}
	height++
	if height != node.height {
		t.Error("Node at index", node.item.index, "has wrong height")
		t.Log("Should be:", height)
		t.Log("Node says:", node.height)
	}

	return height
}

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
