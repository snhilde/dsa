package htree

import (
	"testing"
	"time"
	"math/rand"
	"sort"
	"strings"
	"fmt"
)

// --- TREE TESTS ---
// Test creating a new tree object.
func TestNew(t *testing.T) {
	tr := New()
	if tr == nil {
		t.Error("Failed to create tree")
	}

	if tr.trunk != nil {
		t.Error("Trunk node is not nil")
	}

	if tr.length != 0 {
		t.Error("Tree claims to have nodes")
	}
}

// Test using bad tree objects with all the methods.
func TestBad(t *testing.T) {
	var tr *Tree

	testString(t, tr, "<nil>")
	testLength(t, tr, -1)

	if err := tr.Add(5, 5); err == nil {
		t.Error("Bad object test: Unexpectedly passed Add")
	}
	testString(t, tr, "<nil>")
	testLength(t, tr, -1)

	item := NewItem(5, 5)
	if err := tr.AddItems(item); err == nil {
		t.Error("Bad object test: Unexpectedly passed AddItems")
	}
	testString(t, tr, "<nil>")
	testLength(t, tr, -1)

	if err := tr.Remove(5); err == nil {
		t.Error("Bad object test: Unexpectedly passed Remove")
	}
	testString(t, tr, "<nil>")
	testLength(t, tr, -1)

	if v := tr.Value(5); v != nil {
		t.Error("Bad object test: Unexpectedly passed Value")
	}
	testString(t, tr, "<nil>")
	testLength(t, tr, -1)

	if item := tr.Item(5); item != nil {
		t.Error("Bad object test: Unexpectedly passed Item")
	}
	testString(t, tr, "<nil>")
	testLength(t, tr, -1)

	if ok := tr.Match(5); ok {
		t.Error("Bad object test: Unexpectedly passed Match")
	}
	testString(t, tr, "<nil>")
	testLength(t, tr, -1)

	if ch := tr.Yield(nil); ch != nil {
		t.Error("Bad object test: Unexpectedly passed Yield")
	}
	testString(t, tr, "<nil>")
	testLength(t, tr, -1)

	if l := tr.List(); l != nil {
		t.Error("Bad object test: Unexpectedly passed List")
	}
	testString(t, tr, "<nil>")
	testLength(t, tr, -1)

	if tr.String() != "<nil>" {
		t.Error("Bad object test: Unexpectedly passed String")
	}
	testString(t, tr, "<nil>")
	testLength(t, tr, -1)

	if tr.Length() != -1 {
		t.Error("Bad object test: Unexpectedly passed Length")
	}
	testString(t, tr, "<nil>")
	testLength(t, tr, -1)
}

// Begin testing the various methods for Tree.
func TestAdd(t *testing.T) {
	tr := New()

	// Do a few simple, hand-built tests to make sure things look right.
	if err := tr.Add(5, 5); err != nil {
		t.Error(err)
	}
	testString(t, tr, "5")
	testLength(t, tr, 1)

	if err := tr.Add(10, 10); err != nil {
		t.Error(err)
	}
	testString(t, tr, "5, 10")
	testLength(t, tr, 2)

	if err := tr.Add(1, 1); err != nil {
		t.Error(err)
	}
	testString(t, tr, "1, 5, 10")
	testLength(t, tr, 3)

	// Now do a larger test to make sure items are inserted in the correct order.
	var b strings.Builder
	var nums []int
	tr, nums = buildTree(100000, true)
	sort.Ints(nums)
	for _, v := range nums {
		b.WriteString(fmt.Sprintf("%v, ", v))
	}
	s := strings.TrimSuffix(b.String(), ", ")
	testString(t, tr, s)
	testLength(t, tr, 100000)
}

func TestAddItems(t *testing.T) {
	tr := New()

	// Do a few simple, hand-built tests to make sure things look right.
	item1 := NewItem(5, 5)
	if err := tr.AddItems(item1); err != nil {
		t.Error(err)
	}
	testString(t, tr, "5")
	testLength(t, tr, 1)

	item2 := NewItem(10, 10)
	item3 := NewItem(1, 1)
	if err := tr.AddItems(item2, item3); err != nil {
		t.Error(err)
	}
	testString(t, tr, "1, 5, 10")
	testLength(t, tr, 3)

	// Now do a larger test to make sure items are inserted in the correct order.
	var b strings.Builder
	tr = New()
	nums := make([]int, 10000)
	items := make([]*Item, 10000)
	r := newRand()

	// Build all the items.
	for i := range nums {
		v := r.Int()
		nums[i] = v
		items[i] = NewItem(v, v)
	}

	// Add them all into the tree.
	if err := tr.AddItems(items...); err != nil {
		t.Error(err)
	}

	// Check that everything was added in sorted order.
	sort.Ints(nums)
	for _, v := range nums {
		b.WriteString(fmt.Sprintf("%v, ", v))
	}
	s := strings.TrimSuffix(b.String(), ", ")
	testString(t, tr, s)
	testLength(t, tr, 10000)
}

func TestRemove(t *testing.T) {
	// TODO
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
	var nums []int
	tr, nums = buildTree(100000, true)
	for _, v := range nums {
		if val := tr.Value(v); val != v {
			t.Error("Expected", v, "| Received", val)
		}
	}
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
	if item := tr.Item(2); item != nil {
		t.Error("Expected nothing, Received", item.GetValue())
	}
	if item := tr.Item(20); item != nil {
		t.Error("Expected nothing, Received", item.GetValue())
	}
	if item := tr.Item(100); item != nil {
		t.Error("Expected nothing, Received", item.GetValue())
	}

	// Now do a larger test to make sure the correct item is returned.
	var nums []int
	tr, nums = buildTree(100000, true)
	for _, v := range nums {
		if item := tr.Item(v); item.GetValue() != v {
			t.Error("Wrong value: Expected", v, "| Received", item.GetValue())
		} else if item.GetIndex() != v {
			t.Error("Wrong index: Expected", v, "| Received", item.GetIndex())
		}
	}
}

func TestMatch(t *testing.T) {
}

func TestYield(t *testing.T) {
}

func TestList(t *testing.T) {
}


// --- ITEM TESTS ---


// --- TREE BENCHMARKS ---
func Benchmark100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = buildTree(100, false)
	}
}

func Benchmark1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = buildTree(1000, false)
	}
}

func Benchmark10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = buildTree(10000, false)
	}
}

func Benchmark100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = buildTree(100000, false)
	}
}

func Benchmark1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = buildTree(1000000, false)
	}
}


// --- HELPER FUNCTIONS ---
func testString(t *testing.T, tr *Tree, want string) {
	s := tr.String()
	if s != want {
		t.Error("Expected:", want)
		t.Error("Received:", s)
	}
}

func testLength(t *testing.T, tr *Tree, want int) {
	length := tr.Length()
	if length != want {
		t.Error("Want", want, "items")
		t.Error("Have", length, "items")
	}
}

func newRand() *rand.Rand {
	seed   := time.Now().UnixNano()
	source := rand.NewSource(seed)
	random := rand.New(source)

	return random
}

// buildTree creates a new tree and populates it with count items, either randomly or by iterating from low to high. It
// returns the new tree as well as the indexes of all the items.
func buildTree(count int, random bool) (*Tree, []int) {
	tr := New()
	indexes := make([]int, count)

	if random {
		r := newRand()
		for i := range indexes {
			v := r.Int()
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
