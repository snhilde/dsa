package hlist

import (
	"testing"
)


// TESTS
func TestBadPtr(t *testing.T) {
	// Test that using a non-initialized list is handled correctly.
	var l *List

    // Test String().
	if s := l.String(); s != "<nil>" {
		t.Error("unexpectedly passed String() test with bad pointer")
		t.Log("\tExpected: <nil>")
		t.Log("\tReceived:", s)
	}

    // Test Length().
	if n := l.Length(); n != -1 {
		t.Error("unexpectedly passed Length() test with bad pointer")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", n)
	}

    // Test Insert().
	if err := l.Insert(0, "item"); err == nil {
		t.Error("unexpectedly passed Insert() test with bad pointer")
	}

    // Test Append().
	if err := l.Append("item"); err == nil {
		t.Error("unexpectedly passed Append() test with bad pointer")
	}

    // Test Remove().
	if v := l.Remove(0); v != nil {
		t.Error("unexpectedly passed Remove() test with bad pointer")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}

    // Test RemoveMatch().
	if err := l.RemoveMatch("item"); err == nil {
		t.Error("unexpectedly passed RemoveMatch() test with bad pointer")
	}

    // Test Index().
	if i := l.Index("item"); i != -1 {
		t.Error("unexpectedly passed Index() test with bad pointer")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", i)
	}

    // Test Exists().
	if b := l.Exists("item"); b != false {
		t.Error("unexpectedly passed Exists() test with bad pointer")
		t.Log("\tExpected: false")
		t.Log("\tReceived:", b)
	}

    // Test Copy().
	if _, err := l.Copy(); err == nil {
		t.Error("unexpectedly passed Copy() test with bad pointer")
	}

    // Test Merge().
	if err := l.Merge(New()); err == nil {
		t.Error("unexpectedly passed Merge() test with bad pointer")
	}

    // Test Clear().
	if err := l.Clear(); err == nil {
		t.Error("unexpectedly passed Clear() test with bad pointer")
	}

    // Test Sort().
	if err := l.Sort(func(l, r interface{}) bool {return false}); err == nil {
		t.Error("unexpectedly passed Sort() test with bad pointer")
	}

    // Test SortInt().
	if err := l.SortInt(); err == nil {
		t.Error("unexpectedly passed SortInt() test with bad pointer")
	}

    // Test SortStr().
	if err := l.SortStr(); err == nil {
		t.Error("unexpectedly passed SortStr() test with bad pointer")
	}
}

func TestBadArgs(t *testing.T) {
	// Test that passing bad values to methods is handled correctly.
	l := New()
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

    // Test Insert().
	if err := l.Insert(-1, "item"); err == nil {
		t.Error("unexpectedly passed Insert() test for negative index")
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

    // Test Remove().
	if v := l.Remove(-1); v != nil {
		t.Error("unexpectedly passed Remove() test for negative index")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

    // Test Remove().
	if v := l.Remove(0); v != nil {
		t.Error("unexpectedly passed Remove() test for empty list")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

    // Test Remove().
	if v := l.Remove(100); v != nil {
		t.Error("unexpectedly passed Remove() test for out-of-range index")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

    // Test Sort().
	if err := l.Sort(nil); err == nil {
		t.Error("unexpectedly passed Sort() test for missing sort cb")
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)
}

func TestNew(t *testing.T) {
	if l := New(); l == nil {
		t.Error("Failed to create new list")
	}
}

func TestInsert(t *testing.T) {
	l := New()

	// Add an item to the beginning.
	if err := l.Insert(0, 5); err != nil {
		t.Error(err)
	}
	checkString(t, l, "5")
	checkLength(t, l, 1)

	// Add another item to the beginning.
	if err := l.Insert(0, "bananas"); err != nil {
		t.Error(err)
	}
	checkString(t, l, "bananas, 5")
	checkLength(t, l, 2)

	// Add an item in the middle.
	if err := l.Insert(1, []float64{3.14, 1.23}); err != nil {
		t.Error(err)
	}
	checkString(t, l, "bananas, [3.14 1.23], 5")
	checkLength(t, l, 3)

	// Add some items to the end.
	if err := l.Insert(l.Length(), 1, 2, 3); err != nil {
		t.Error(err)
	}
	checkString(t, l, "bananas, [3.14 1.23], 5, 1, 2, 3")
	checkLength(t, l, 6)

	// Add a nil item.
	if err := l.Insert(0, nil); err != nil {
		t.Error(err)
	}
	checkString(t, l, "<nil>, bananas, [3.14 1.23], 5, 1, 2, 3")
	checkLength(t, l, 7)
}

func TestAppend(t *testing.T) {
	l := New()

	// Add one item.
	if err := l.Append(5); err != nil {
		t.Error(err)
	}
	checkString(t, l, "5")
	checkLength(t, l, 1)

	// Add another item.
	if err := l.Append("bananas"); err != nil {
		t.Error(err)
	}
	checkString(t, l, "5, bananas")
	checkLength(t, l, 2)

	// Add multiple items.
	if err := l.Append([]interface{}{3.14, 1.23}...); err != nil {
		t.Error(err)
	}
	checkString(t, l, "5, bananas, 3.14, 1.23")
	checkLength(t, l, 4)

	// Add multiple items of different types.
	if err := l.Append("a", 1, rune('0')); err != nil {
		t.Error(err)
	}
	checkString(t, l, "5, bananas, 3.14, 1.23, a, 1, 48")
	checkLength(t, l, 7)

	// Add a nil item.
	if err := l.Append(nil); err != nil {
		t.Error(err)
	}
	checkString(t, l, "5, bananas, 3.14, 1.23, a, 1, 48, <nil>")
	checkLength(t, l, 8)

	// Add multiple items of different types as the first items.
	l.Clear()
	if err := l.Append("a", 1, rune('0')); err != nil {
		t.Error(err)
	}
	checkString(t, l, "a, 1, 48")
	checkLength(t, l, 3)
}

func TestRemove(t *testing.T) {
	l := New()

	l.Append(1, 2, 3, "4", []byte{0x05, 0x06})
	checkString(t, l, "1, 2, 3, 4, [5 6]")
	checkLength(t, l, 5)

	// Remove the 3rd item.
	if value := l.Remove(2); value != 3 {
		t.Error("Error removing 3rd item")
		t.Log("Expected: 3")
		t.Log("Received:", value)
	}
	checkString(t, l, "1, 2, 4, [5 6]")
	checkLength(t, l, 4)

	// Remove the new 3rd item.
	if value := l.Remove(2); value != "4" {
		t.Error("Error removing new 3rd item")
		t.Log("Expected: 4")
		t.Log("Received:", value)
	}
	checkString(t, l, "1, 2, [5 6]")
	checkLength(t, l, 3)

	// Remove the last item.
	value := l.Remove(l.Length()-1)
	if len(value.([]byte)) != 2 || value.([]byte)[0] != 0x05 || value.([]byte)[1] != 0x06 {
		t.Error("Error removing last item")
		t.Log("Expected: [0x05, 0x06")
		t.Log("Received:", value)
	}
	checkString(t, l, "1, 2")
	checkLength(t, l, 2)

	// Remove the first item.
	if value := l.Remove(0); value != 1 {
		t.Error("Error removing first item")
		t.Log("Expected: 1")
		t.Log("Received:", value)
	}
	checkString(t, l, "2")
	checkLength(t, l, 1)

	// Remove the last remaining item.
	if value := l.Remove(0); value != 2 {
		t.Error("Error removing last remaining item")
		t.Log("Expected: 2")
		t.Log("Received:", value)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

	// Make sure nothing is left.
	if value := l.Remove(0); value != nil {
		t.Error("Unexpectedly found an item")
		t.Log("Expected: nil")
		t.Log("Received:", value)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)
}

func TestRemoveMatch(t *testing.T) {
	l := New()
	l.Append(1, "apples", 3, []int{4, 5}, 3.14)
	checkString(t, l, "1, apples, 3, [4 5], 3.14")
	checkLength(t, l, 5)

	if err := l.RemoveMatch(3); err != nil {
		t.Error(err)
	}
	checkString(t, l, "1, apples, [4 5], 3.14")
	checkLength(t, l, 4)

	if err := l.RemoveMatch(1); err != nil {
		t.Error(err)
	}
	checkString(t, l, "apples, [4 5], 3.14")
	checkLength(t, l, 3)

	if err := l.RemoveMatch(nil); err != nil {
		t.Error(err)
	}
	checkString(t, l, "apples, [4 5], 3.14")
	checkLength(t, l, 3)

	if err := l.RemoveMatch("apples"); err != nil {
		t.Error(err)
	}
	checkString(t, l, "[4 5], 3.14")
	checkLength(t, l, 2)

	if err := l.RemoveMatch(3.14); err != nil {
		t.Error(err)
	}
	checkString(t, l, "[4 5]")
	checkLength(t, l, 1)

	// Try to remove a non-existant item.
	if err := l.RemoveMatch(10); err != nil {
		t.Error(err)
	}
	checkString(t, l, "[4 5]")
	checkLength(t, l, 1)

	if err := l.RemoveMatch([]int{4, 5}); err != nil {
		t.Error(err)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)
}

func TestIndex(t *testing.T) {
	l := New()
	l.Append("apples", 1, 3, 3.14, []byte{0xEE, 0xFF}, "aardvark")
	checkString(t, l, "apples, 1, 3, 3.14, [238 255], aardvark")
	checkLength(t, l, 6)

	// Check index of 1.
	if i := l.Index(1); i != 1 {
		t.Error("Incorrect index for 1")
		t.Log("Expected: 1")
		t.Log("Received:", i)
	}

	// Check index of "apples".
	if i := l.Index("apples"); i != 0 {
		t.Error("Incorrect index for \"apples\"")
		t.Log("Expected: 0")
		t.Log("Received:", i)
	}

	// Check index of pi.
	if i := l.Index(3.14); i != 3 {
		t.Error("Incorrect index for 3.14")
		t.Log("Expected: 3")
		t.Log("Received:", i)
	}

	// Remove and index and check index of pi again.
	l.Remove(1)
	if i := l.Index(3.14); i != 2 {
		t.Error("Incorrect index for 3.14")
		t.Log("Expected: 2")
		t.Log("Received:", i)
	}

	// Try to find a non-existant item.
	if i := l.Index(10); i != -1 {
		t.Error("Unexpectedly passed no-match test")
		t.Log("Expected: -1")
		t.Log("Received:", i)
	}

	// Test matching a slice.
	if i := l.Index([]byte{0xEE, 0xFF}); i != 3 {
		t.Error("Incorrect index for []byte{0xEE, 0xFF}")
		t.Log("Expected: 3")
		t.Log("Received:", i)
	}

	// Test not matching a slice.
	if i := l.Index([]byte{0xAA, 0xBB}); i != -1 {
		t.Error("Unexpectedly passed no-match slice test")
		t.Log("Expected: -1")
		t.Log("Received:", i)
	}
}


// HELPERS
func checkString(t *testing.T, l *List, want string) {
	if l.String() != want {
		t.Error("List contents are incorrect")
		t.Log("Expected:", want)
		t.Log("Received:", l)
	}
}

func checkLength(t *testing.T, l *List, want int) {
	if l.Length() != want {
		t.Error("Incorrect length")
		t.Log("Expected:", want)
		t.Log("Received:", l.Length())
	}
}
