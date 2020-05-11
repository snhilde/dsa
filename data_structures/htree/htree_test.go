package htree

import (
	"testing"
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
	testString(t, tr, "<empty>")
	testLength(t, tr, 0)

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
}

func TestAddItems(t *testing.T) {
}

func TestRemove(t *testing.T) {
}

func TestValue(t *testing.T) {
}

func TestItem(t *testing.T) {
}

func TestMatch(t *testing.T) {
}

func TestYield(t *testing.T) {
}

func TestList(t *testing.T) {
}


// --- ITEM TESTS ---


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
