package hlist_test

import (
	"reflect"
	"testing"

	"github.com/snhilde/dsa/data_structures/hlist"
)

func TestBadPtr(t *testing.T) {
	// Test that using a non-initialized list is handled correctly.
	var l *hlist.List

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

	// Test Index().
	if i := l.Index("item"); i != -1 {
		t.Error("unexpectedly passed Index() test with bad pointer")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", i)
	}

	// Test Item().
	if v := l.Item(0); v != nil {
		t.Error("unexpectedly passed Item() test with bad pointer")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}

	// Test Items().
	if v := l.Items(); v != nil {
		t.Error("unexpectedly passed Items() test with bad pointer")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}

	// Test Exists().
	if b := l.Exists("item"); b != false {
		t.Error("unexpectedly passed Exists() test with bad pointer")
		t.Log("\tExpected: false")
		t.Log("\tReceived:", b)
	}

	// Test Remove().
	if v := l.Remove(0); v != nil {
		t.Error("unexpectedly passed Remove() test with bad pointer")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}

	// Test Copy().
	if _, err := l.Copy(); err == nil {
		t.Error("unexpectedly passed Copy() test with bad pointer")
	}

	// Test Same().
	if l.Same(hlist.New()) {
		t.Error("unexpectedly passed Same() test with bad pointer")
	}

	// Test Twin().
	if l.Twin(hlist.New()) {
		t.Error("unexpectedly passed Twin() test with bad pointer")
	}

	// Test Merge().
	if err := l.Merge(hlist.New()); err == nil {
		t.Error("unexpectedly passed Merge() test with bad pointer")
	}

	// Test Clear().
	if err := l.Clear(); err == nil {
		t.Error("unexpectedly passed Clear() test with bad pointer")
	}

	// Test Yield().
	if ch := l.Yield(nil); ch != nil {
		t.Error("unexpectedly passed Yield() test with bad pointer")
	}

	// Test Sort().
	if err := l.Sort(func(l, r interface{}) bool { return false }); err == nil {
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
	l := hlist.New()
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

	// Test Insert().
	if err := l.Insert(-1, "item"); err == nil {
		t.Error("unexpectedly passed Insert() test for negative index")
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

	// Test Insert().
	if err := l.Insert(100, "item"); err == nil {
		t.Error("unexpectedly passed Insert() test for out-of-range index")
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

	// Test Item().
	if v := l.Item(-1); v != nil {
		t.Error("unexpectedly passed Item() test for negative index")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

	// Test Item().
	if v := l.Item(0); v != nil {
		t.Error("unexpectedly passed Item() test for empty list")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

	// Test Item().
	if v := l.Item(100); v != nil {
		t.Error("unexpectedly passed Item() test for out-of-range index")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
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

	// Test Same().
	if l.Same(nil) {
		t.Error("unexpectedly passed Same() test for missing arg")
	}

	// Test Twin().
	if l.Twin(nil) {
		t.Error("unexpectedly passed Twin() test for missing arg")
	}

	// Test Sort().
	if err := l.Sort(nil); err == nil {
		t.Error("unexpectedly passed Sort() test for missing sort cb")
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

	// Test trying to add list to itself.
	l.Append(1, 2, 3)
	checkString(t, l, "1, 2, 3")
	checkLength(t, l, 3)

	if err := l.Append(4, l, 6); err == nil {
		t.Error("unexpectedly passed Append() test for adding list to itself")
	}
	checkString(t, l, "1, 2, 3")
	checkLength(t, l, 3)

	if err := l.Insert(2, 4, l, 6); err == nil {
		t.Error("unexpectedly passed Insert() test for adding list to itself")
	}
	checkString(t, l, "1, 2, 3")
	checkLength(t, l, 3)
}

func TestNew(t *testing.T) {
	if l := hlist.New(); l == nil {
		t.Error("Failed to create new list")
	}
}

func TestInsert(t *testing.T) {
	l := hlist.New()

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
	l := hlist.New()

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

func TestIndex(t *testing.T) {
	l := hlist.New()
	l.Append("apples", 1, 3, 3.14, []byte{0xEE, 0xFF}, "aardvark")
	checkString(t, l, "apples, 1, 3, 3.14, [238 255], aardvark")
	checkLength(t, l, 6)

	// Check index of 1.
	if i := l.Index(1); i != 1 {
		t.Error("Incorrect index for 1")
		t.Log("\tExpected: 1")
		t.Log("\tReceived:", i)
	}

	// Check index of "apples".
	if i := l.Index("apples"); i != 0 {
		t.Error("Incorrect index for \"apples\"")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", i)
	}

	// Check index of pi.
	if i := l.Index(3.14); i != 3 {
		t.Error("Incorrect index for 3.14")
		t.Log("\tExpected: 3")
		t.Log("\tReceived:", i)
	}

	// Remove an item and check index of pi again.
	l.Remove(1)
	if i := l.Index(3.14); i != 2 {
		t.Error("Incorrect index for 3.14")
		t.Log("\tExpected: 2")
		t.Log("\tReceived:", i)
	}

	// Try to find a non-existent item.
	if i := l.Index(10); i != -1 {
		t.Error("Unexpectedly passed no-match test")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", i)
	}

	// Test matching a slice.
	if i := l.Index([]byte{0xEE, 0xFF}); i != 3 {
		t.Error("Incorrect index for []byte{0xEE, 0xFF}")
		t.Log("\tExpected: 3")
		t.Log("\tReceived:", i)
	}

	// Test not matching a slice.
	if i := l.Index([]byte{0xAA, 0xBB}); i != -1 {
		t.Error("Unexpectedly passed no-match slice test")
		t.Log("\tExpected: -1")
		t.Log("\tReceived:", i)
	}
}

func TestItem(t *testing.T) {
	l := hlist.New()
	l.Append("apples", 1, 3, 3.14, []byte{0xEE, 0xFF}, "aardvark")
	checkString(t, l, "apples, 1, 3, 3.14, [238 255], aardvark")
	checkLength(t, l, 6)

	// Check Item at first position.
	if v := l.Item(0); v != "apples" {
		t.Error("Incorrect Item for index 0")
		t.Log("\tExpected: apples")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "apples, 1, 3, 3.14, [238 255], aardvark")
	checkLength(t, l, 6)

	// Check Item at 3rd position.
	if v := l.Item(2); v != 3 {
		t.Error("Incorrect Item for index 2")
		t.Log("\tExpected: 3")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "apples, 1, 3, 3.14, [238 255], aardvark")
	checkLength(t, l, 6)

	// Check Item at 4th position.
	if v := l.Item(3); v != 3.14 {
		t.Error("Incorrect Item for index 3")
		t.Log("\tExpected: 3.14")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "apples, 1, 3, 3.14, [238 255], aardvark")
	checkLength(t, l, 6)

	// Check Item at last position.
	if v := l.Item(5); v != "aardvark" {
		t.Error("Incorrect Item for index 5")
		t.Log("\tExpected: aardvark")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "apples, 1, 3, 3.14, [238 255], aardvark")
	checkLength(t, l, 6)

	// Remove an item and check value at index 3 again.
	l.Remove(1)
	if v := l.Item(3); !reflect.DeepEqual(v, []byte{0xEE, 0xFF}) {
		t.Error("Incorrect value for index 3")
		t.Log("\tExpected: [238 255]")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "apples, 3, 3.14, [238 255], aardvark")
	checkLength(t, l, 5)

	// Try to find a non-existent item.
	if v := l.Item(10); v != nil {
		t.Error("Unexpectedly passed no-match test")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}
}

func TestItems(t *testing.T) {
	l := hlist.New()
	l.Append(1, "2", 3.14)
	checkString(t, l, "1, 2, 3.14")
	checkLength(t, l, 3)

	// Test that Items() returns each item but does not alter the list in any way.
	items := l.Items()
	if items == nil {
		t.Error("Failed to receive items")
	} else if len(items) != 3 {
		t.Error("Did not receive all items in list")
		t.Log("\tExpected: 3")
		t.Log("\tReceived:", len(items))
	}
	checkString(t, l, "1, 2, 3.14")
	checkLength(t, l, 3)

	// Make sure that all of the items in the slice are correct.
	for i, v := range items {
		switch i {
		case 0:
			if v != 1 {
				t.Error("Error receiving 1st item")
				t.Log("\tExpected: 1")
				t.Log("\tReceived:", v)
			}
		case 1:
			if v != "2" {
				t.Error("Error receiving 2nd item")
				t.Log("\tExpected: 2")
				t.Log("\tReceived:", v)
			}
		case 2:
			if v != 3.14 {
				t.Error("Error receiving 3rd item")
				t.Log("\tExpected: 3.14")
				t.Log("\tReceived:", v)
			}
		}
	}
	checkString(t, l, "1, 2, 3.14")
	checkLength(t, l, 3)

	// Test that modifying list does not affect anything in channel.
	items = l.Items()
	if items == nil {
		t.Error("Failed to receive items")
	} else if len(items) != 3 {
		t.Error("Did not receive all items in list")
		t.Log("\tExpected: 3")
		t.Log("\tReceived:", len(items))
	}
	l.Remove(2)
	l.Remove(0)
	checkString(t, l, "2")
	checkLength(t, l, 1)
	for i, v := range items {
		switch i {
		case 0:
			if v != 1 {
				t.Error("Error receiving 1st item")
				t.Log("\tExpected: 1")
				t.Log("\tReceived:", v)
			}
		case 1:
			if v != "2" {
				t.Error("Error receiving 2nd item")
				t.Log("\tExpected: 2")
				t.Log("\tReceived:", v)
			}
		case 2:
			if v != 3.14 {
				t.Error("Error receiving 3rd item")
				t.Log("\tExpected: 3.14")
				t.Log("\tReceived:", v)
			}
		}
	}

	// Test that new items are in the new list.
	l.Append("dd", "ee", "ff")
	checkString(t, l, "2, dd, ee, ff")
	checkLength(t, l, 4)
	items = l.Items()
	if items == nil {
		t.Error("Failed to receive items")
	} else if len(items) != 4 {
		t.Error("Did not receive all items in list")
		t.Log("\tExpected: 4")
		t.Log("\tReceived:", len(items))
	}
	for i, v := range items {
		switch i {
		case 0:
			if v != "2" {
				t.Error("Error receiving 1st item")
				t.Log("\tExpected: 2")
				t.Log("\tReceived:", v)
			}
		case 1:
			if v != "dd" {
				t.Error("Error receiving 2nd item")
				t.Log("\tExpected: dd")
				t.Log("\tReceived:", v)
			}
		case 2:
			if v != "ee" {
				t.Error("Error receiving 3rd item")
				t.Log("\tExpected: ee")
				t.Log("\tReceived:", v)
			}
		case 3:
			if v != "ff" {
				t.Error("Error receiving 4rd item")
				t.Log("\tExpected: ff")
				t.Log("\tReceived:", v)
			}
		}
	}
}

func TestExists(t *testing.T) {
	l := hlist.New()
	l.Append(1, 3, "apples", []byte{0xEE, 0xFF}, 3.14, 4)
	checkString(t, l, "1, 3, apples, [238 255], 3.14, 4")
	checkLength(t, l, 6)

	// Check that 3 exists.
	if ret := l.Exists(3); !ret {
		t.Error("Expected match for 3, but didn't find it")
	}

	// Check that "apples" exists.
	if ret := l.Exists("apples"); !ret {
		t.Error("Expected match for \"apples\", but didn't find it")
	}

	// Check that 3.14 exists.
	if ret := l.Exists(3.14); !ret {
		t.Error("Expected match for 3,14, but didn't find it")
	}

	// Try to find a non-existent item.
	if ret := l.Exists(10); ret {
		t.Error("Unexpectedly passed no-match test")
	}

	// Test matching a slice.
	if ret := l.Exists([]byte{0xEE, 0xFF}); !ret {
		t.Error("Expected match for []byte{0xEE, 0xFF}, but didn't find it")
	}

	// Test not matching a slice.
	if ret := l.Exists([]byte{0xBB, 0xCC}); ret {
		t.Error("Unexpectedly passed no-match slice test")
	}
}

func TestRemove(t *testing.T) {
	l := hlist.New()

	l.Append(1, 2, 3, "4", []byte{0x05, 0x06})
	checkString(t, l, "1, 2, 3, 4, [5 6]")
	checkLength(t, l, 5)

	// Remove the 3rd item.
	if value := l.Remove(2); value != 3 {
		t.Error("Error removing 3rd item")
		t.Log("\tExpected: 3")
		t.Log("\tReceived:", value)
	}
	checkString(t, l, "1, 2, 4, [5 6]")
	checkLength(t, l, 4)

	// Remove the new 3rd item.
	if value := l.Remove(2); value != "4" {
		t.Error("Error removing new 3rd item")
		t.Log("\tExpected: 4")
		t.Log("\tReceived:", value)
	}
	checkString(t, l, "1, 2, [5 6]")
	checkLength(t, l, 3)

	// Remove the last item.
	value := l.Remove(l.Length() - 1)
	if len(value.([]byte)) != 2 || value.([]byte)[0] != 0x05 || value.([]byte)[1] != 0x06 {
		t.Error("Error removing last item")
		t.Log("\tExpected: [0x05, 0x06")
		t.Log("\tReceived:", value)
	}
	checkString(t, l, "1, 2")
	checkLength(t, l, 2)

	// Remove the first item.
	if value := l.Remove(0); value != 1 {
		t.Error("Error removing first item")
		t.Log("\tExpected: 1")
		t.Log("\tReceived:", value)
	}
	checkString(t, l, "2")
	checkLength(t, l, 1)

	// Remove the last remaining item.
	if value := l.Remove(0); value != 2 {
		t.Error("Error removing last remaining item")
		t.Log("\tExpected: 2")
		t.Log("\tReceived:", value)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

	// Make sure nothing is left.
	if value := l.Remove(0); value != nil {
		t.Error("Unexpectedly found an item")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", value)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)
}

func TestRemoveMatch(t *testing.T) {
	l := hlist.New()
	l.Append(1, "apples", 3, []int{4, 5}, 3.14)
	checkString(t, l, "1, apples, 3, [4 5], 3.14")
	checkLength(t, l, 5)

	l.RemoveMatch(3)
	checkString(t, l, "1, apples, [4 5], 3.14")
	checkLength(t, l, 4)

	l.RemoveMatch(1)
	checkString(t, l, "apples, [4 5], 3.14")
	checkLength(t, l, 3)

	l.RemoveMatch(nil)
	checkString(t, l, "apples, [4 5], 3.14")
	checkLength(t, l, 3)

	l.RemoveMatch("apples")
	checkString(t, l, "[4 5], 3.14")
	checkLength(t, l, 2)

	l.RemoveMatch(3.14)
	checkString(t, l, "[4 5]")
	checkLength(t, l, 1)

	// Try to remove a non-existent item.
	l.RemoveMatch(10)
	checkString(t, l, "[4 5]")
	checkLength(t, l, 1)

	l.RemoveMatch([]int{4, 5})
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)
}

func TestCopy(t *testing.T) {
	l := hlist.New()

	// Copy an empty list.
	nl, err := l.Copy()
	if err != nil {
		t.Error(err)
	}
	checkString(t, nl, "<empty>")
	checkLength(t, nl, 0)

	// Copy a non-empty list.
	l.Append("sizzle", 1e5, 3.1415, 15)
	checkString(t, l, "sizzle, 100000, 3.1415, 15")
	checkLength(t, l, 4)

	nl, err = l.Copy()
	if err != nil {
		t.Error(err)
	}
	checkString(t, nl, "sizzle, 100000, 3.1415, 15")
	checkLength(t, nl, 4)

	// Make sure the two lists are separate.
	checkString(t, l, "sizzle, 100000, 3.1415, 15")
	checkLength(t, l, 4)
	checkString(t, nl, "sizzle, 100000, 3.1415, 15")
	checkLength(t, nl, 4)

	l.Remove(1)
	checkString(t, l, "sizzle, 3.1415, 15")
	checkLength(t, l, 3)
	checkString(t, nl, "sizzle, 100000, 3.1415, 15")
	checkLength(t, nl, 4)

	nl.Remove(2)
	checkString(t, l, "sizzle, 3.1415, 15")
	checkLength(t, l, 3)
	checkString(t, nl, "sizzle, 100000, 15")
	checkLength(t, nl, 3)
}

func TestSame(t *testing.T) {
	l1 := hlist.New()
	l2 := l1

	// Test an empty list.
	if !l1.Same(l2) {
		t.Error("Lists differ (test #1)")
	}
	if !l2.Same(l1) {
		t.Error("Lists differ (test #2)")
	}
	checkString(t, l1, "<empty>")
	checkLength(t, l1, 0)
	checkString(t, l2, "<empty>")
	checkLength(t, l2, 0)

	// And an item and test again.
	l1.Append("item")
	if !l1.Same(l2) {
		t.Error("Lists differ (test #3)")
	}
	if !l2.Same(l1) {
		t.Error("Lists differ (test #4)")
	}
	checkString(t, l1, "item")
	checkLength(t, l1, 1)
	checkString(t, l2, "item")
	checkLength(t, l2, 1)

	// And multiple items to the 2nd reference.
	l2.Append(3.14, []int{1, 2, 3}, 5)
	if !l1.Same(l2) {
		t.Error("Lists differ (test #5)")
	}
	if !l2.Same(l1) {
		t.Error("Lists differ (test #6)")
	}
	checkString(t, l1, "item, 3.14, [1 2 3], 5")
	checkLength(t, l1, 4)
	checkString(t, l2, "item, 3.14, [1 2 3], 5")
	checkLength(t, l2, 4)

	// Remove an item and test again.
	l1.Remove(3)
	if !l1.Same(l2) {
		t.Error("Lists differ (test #7)")
	}
	if !l2.Same(l1) {
		t.Error("Lists differ (test #8)")
	}
	checkString(t, l1, "item, 3.14, [1 2 3]")
	checkLength(t, l1, 3)
	checkString(t, l2, "item, 3.14, [1 2 3]")
	checkLength(t, l2, 3)

	// Test copying, and make sure that the old lists and the new list are no longer the same.
	l3, err := l1.Copy()
	if err != nil {
		t.Error(err)
	}
	if l1.Same(l3) {
		t.Error("Lists are unexpectedly the same (test #9)")
	}
	if l2.Same(l3) {
		t.Error("Lists are unexpectedly the same (test #10)")
	}

	// Test clearing a list.
	l2.Clear()
	if !l1.Same(l2) {
		t.Error("Lists differ (test #11)")
	}
	if !l2.Same(l1) {
		t.Error("Lists differ (test #12)")
	}
	checkString(t, l1, "<empty>")
	checkLength(t, l1, 0)
	checkString(t, l2, "<empty>")
	checkLength(t, l2, 0)

	// Test reassigning to make a same list.
	l3 = l2
	if !l1.Same(l3) {
		t.Error("Lists differ (test #13)")
	}
	if !l2.Same(l3) {
		t.Error("Lists differ (test #14)")
	}

	// Make two lists that have the same contents but are not the same underlying lists.
	l1 = hlist.New()
	l1.Append("apple", "banana", "carrot")
	l2 = hlist.New()
	l2.Append("apple", "banana", "carrot")
	if l1.Same(l2) {
		t.Error("Lists are unexpectedly the same (test #15)")
	}
	if l2.Same(l1) {
		t.Error("Lists are unexpectedly the same (test #16)")
	}
}

func TestTwin(t *testing.T) {
	// Make two lists that have the same contents but are not the same underlying lists.
	l1 := hlist.New()
	l1.Append("apple", "banana", "carrot")
	l2 := hlist.New()
	l2.Append("apple", "banana", "carrot")
	if !l1.Twin(l2) {
		t.Error("Lists unexpectedly failed twin test (test #1)")
	}
	if !l2.Twin(l1) {
		t.Error("Lists unexpectedly failed twin test (test #2)")
	}

	// Test on two lists that are the same.
	l2 = l1
	if !l1.Same(l2) {
		t.Error("Lists are the same. Not running twin test.")
	} else {
		if l1.Twin(l2) {
			t.Error("Lists are twins but should not be (test #3)")
		}
		if l2.Twin(l1) {
			t.Error("Lists are twins but should not be (test #4)")
		}
	}

	// Make sure that typically uncomparable items can be matched.
	type s struct {
		r rune
		i []int
		n int
	}
	l1 = hlist.New()
	l1.Append("apple", []float32{1.23, 2.34, 3.45}, s{'a', []int{8, 9}, 1e3})
	l2 = hlist.New()
	l2.Append("apple", []float32{1.23, 2.34, 3.45}, s{'a', []int{8, 9}, 1e3})
	if !l1.Twin(l2) {
		t.Error("Lists unexpectedly failed twin test (test #5)")
	}
	if !l2.Twin(l1) {
		t.Error("Lists unexpectedly failed twin test (test #6)")
	}

	// Tweak l2 a bit.
	l2 = hlist.New()
	l2.Append("apple", []float32{1.23, 2.34, 3.45}, s{'a', []int{8, 999}, 1e3})
	if l1.Twin(l2) {
		t.Error("Lists are twins but should not be (test #7)")
	}
	if l2.Twin(l1) {
		t.Error("Lists unexpectedly failed twin test (test #8)")
	}
}

func TestMerge(t *testing.T) {
	// Test merging two good lists togther.
	l := hlist.New()
	l.Append(0, 1, 2, 3, 4)
	checkString(t, l, "0, 1, 2, 3, 4")
	checkLength(t, l, 5)

	nl := hlist.New()
	nl.Append(5, 6, 7, 8, 9)
	checkString(t, nl, "5, 6, 7, 8, 9")
	checkLength(t, nl, 5)

	if err := l.Merge(nl); err != nil {
		t.Error(err)
	}
	checkString(t, l, "0, 1, 2, 3, 4, 5, 6, 7, 8, 9")
	checkLength(t, l, 10)
	checkString(t, nl, "<empty>")
	checkLength(t, nl, 0)

	// Test merging a good list and a bad list.
	l.Clear()
	l.Append(0, 1, 2, 3, 4)
	checkString(t, l, "0, 1, 2, 3, 4")
	checkLength(t, l, 5)

	var lp *hlist.List
	checkString(t, lp, "<nil>")
	checkLength(t, lp, -1)

	if err := l.Merge(lp); err != nil {
		t.Error(err)
	}
	checkString(t, l, "0, 1, 2, 3, 4")
	checkLength(t, l, 5)
	checkString(t, lp, "<nil>")
	checkLength(t, lp, -1)

	// Test merging a bad list with a good list.
	if err := lp.Merge(l); err == nil {
		t.Error("Unexpectedly passed bad merge")
	}
	checkString(t, lp, "<nil>")
	checkLength(t, lp, -1)
	checkString(t, l, "0, 1, 2, 3, 4")
	checkLength(t, l, 5)
}

func TestClear(t *testing.T) {
	l := hlist.New()

	// Add some items to the list.
	l.Append(5, "bronto", "stego", 65e6, "t. rex", 0x0C)
	checkString(t, l, "5, bronto, stego, 6.5e+07, t. rex, 12")
	checkLength(t, l, 6)

	// Test clearing out the items.
	l.Clear()
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)

	// Test that clearing one list does not clear a copy.
	l.Append(5, "bronto", "stego", 65e6, "t. rex", 0x0C)
	nl, err := l.Copy()
	if err != nil {
		t.Error(err)
	}
	checkString(t, l, "5, bronto, stego, 6.5e+07, t. rex, 12")
	checkLength(t, l, 6)
	checkString(t, nl, "5, bronto, stego, 6.5e+07, t. rex, 12")
	checkLength(t, nl, 6)

	l.Clear()
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)
	checkString(t, nl, "5, bronto, stego, 6.5e+07, t. rex, 12")
	checkLength(t, nl, 6)
}

func TestYield(t *testing.T) {
	l := hlist.New()
	l.Append(1, "2", 3.14)
	checkString(t, l, "1, 2, 3.14")
	checkLength(t, l, 3)

	// Test that Yield() returns an unbuffered channel and does not alter the list in any way.
	ch := l.Yield(nil)
	if ch == nil {
		t.Error("Failed to receive channel")
	} else if len(ch) != 0 {
		t.Error("Channel is buffered")
		t.Log("\tExpected: 0")
		t.Log("\tReceived:", len(ch))
	}
	checkString(t, l, "1, 2, 3.14")
	checkLength(t, l, 3)

	// Make sure that all the items received from the channel are correct.
	i := 0
	for v := range ch {
		switch i {
		case 0:
			if v != 1 {
				t.Error("Error receiving 1st item")
				t.Log("\tExpected: 1")
				t.Log("\tReceived:", v)
			}
		case 1:
			if v != "2" {
				t.Error("Error receiving 2nd item")
				t.Log("\tExpected: 2")
				t.Log("\tReceived:", v)
			}
		case 2:
			if v != 3.14 {
				t.Error("Error receiving 3rd item")
				t.Log("\tExpected: 3.14")
				t.Log("\tReceived:", v)
			}
		default:
			t.Error("Still receiving items when list should be exhausted")
			t.Log("\tReceived:", v)
		}
		i++
	}
	checkString(t, l, "1, 2, 3.14")
	checkLength(t, l, 3)

	// Make sure that we received all 3 items.
	if i != 3 {
		t.Error("Failed to receive all items in list")
		t.Log("\tExpected: 3 items")
		t.Log("\tReceived:", i, "items")
	}

	// Make sure that the channel is closed after exhausting the list.
	if v, ok := <-ch; ok {
		t.Error("Channel is not closed")
		t.Log("\tReceived:", v)
	}

	// Test breaking out of loop early and entering back in at same point.
	ch = l.Yield(nil)
	if ch == nil {
		t.Error("Failed to receive channel")
	}
	for v := range ch {
		if v != 1 {
			t.Error("Error receiving 1st item")
			t.Log("\tExpected: 1")
			t.Log("\tReceived:", v)
		}
		break
	}
	checkString(t, l, "1, 2, 3.14")
	checkLength(t, l, 3)

	for v := range ch {
		if v != "2" {
			t.Error("Error receiving 2nd item")
			t.Log("\tExpected: 2")
			t.Log("\tReceived:", v)
		}
		break
	}
	checkString(t, l, "1, 2, 3.14")
	checkLength(t, l, 3)

	if v := <-ch; v != 3.14 {
		t.Error("Error receiving 3rd item")
		t.Log("\tExpected: 3.14")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "1, 2, 3.14")
	checkLength(t, l, 3)

	// Test that modifying the list also affects the channel.
	ch = l.Yield(nil)
	if ch == nil {
		t.Error("Failed to receive channel")
	}
	if v := <-ch; v != 1 {
		t.Error("Error receiving 1st item")
		t.Log("\tExpected: 1")
		t.Log("\tReceived:", v)
	}
	l.Remove(2)
	if v := <-ch; v != "2" {
		t.Error("Error receiving 2nd item")
		t.Log("\tExpected: 2")
		t.Log("\tReceived:", v)
	}
	if v, ok := <-ch; ok {
		t.Error("Channel should be closed")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}
	checkString(t, l, "1, 2")
	checkLength(t, l, 2)

	// Test adding items after channel has been established.
	l.Clear()
	l.Append(1)
	ch = l.Yield(nil)
	if ch == nil {
		t.Error("Failed to receive channel")
	}
	l.Append(2, 3)

	i = 0
	for v := range ch {
		switch i {
		case 0:
			if v != 1 {
				t.Error("Error receiving 1st item")
				t.Log("\tExpected: 1")
				t.Log("\tReceived:", v)
			}
		case 1:
			if v != 2 {
				t.Error("Error receiving 2nd item")
				t.Log("\tExpected: 2")
				t.Log("\tReceived:", v)
			}
		case 2:
			if v != 3 {
				t.Error("Error receiving 3rd item")
				t.Log("\tExpected: 3.14")
				t.Log("\tReceived:", v)
			}
		default:
			t.Error("Still receiving items when list should be exhausted")
			t.Log("\tReceived:", v)
		}
		i++
	}
	checkString(t, l, "1, 2, 3")
	checkLength(t, l, 3)

	// Test that an empty list will return an empty channel.
	l.Clear()
	ch = l.Yield(nil)
	if ch != nil {
		t.Error("Channel should be closed")
	}

	// Test that adding items to an empty list will not affect the channel.
	l.Clear()
	ch = l.Yield(nil)
	l.Append(1, 2, 3)
	if ch != nil {
		t.Error("Channel should be closed")
	}

	// Test breaking iteration before receiving any values.
	l.Clear()
	l.Append(1, 2, 3)
	quit := make(chan struct{})
	ch = l.Yield(quit)
	if ch == nil {
		t.Error("Failed to receive channel")
	}
	quit <- struct{}{}
	if v, ok := <-ch; ok {
		t.Error("Channel should be closed")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}

	// Test breaking iteration after receiving one value.
	l.Clear()
	l.Append(1, 2, 3)
	quit = make(chan struct{})
	ch = l.Yield(quit)
	if ch == nil {
		t.Error("Failed to receive channel")
	}

	if v := <-ch; v != 1 {
		t.Error("Error receiving 1st item")
		t.Log("\tExpected: 1")
		t.Log("\tReceived:", v)
	}

	quit <- struct{}{}
	if v, ok := <-ch; ok {
		t.Error("Channel should be closed")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}

	// Test breaking iteration after receiving all values.
	l.Clear()
	l.Append(1, 2, 3)
	quit = make(chan struct{})
	ch = l.Yield(quit)
	if ch == nil {
		t.Error("Failed to receive channel")
	}

	i = 0
	for v := range ch {
		switch i {
		case 0:
			if v != 1 {
				t.Error("Error receiving 1st item")
				t.Log("\tExpected: 1")
				t.Log("\tReceived:", v)
			}
		case 1:
			if v != 2 {
				t.Error("Error receiving 2nd item")
				t.Log("\tExpected: 2")
				t.Log("\tReceived:", v)
			}
		case 2:
			if v != 3 {
				t.Error("Error receiving 3rd item")
				t.Log("\tExpected: 3.14")
				t.Log("\tReceived:", v)
			}
		default:
			t.Error("Still receiving items when list should be exhausted")
			t.Log("\tReceived:", v)
		}
		i++
	}

	select {
	case quit <- struct{}{}:
		t.Error("nothing should be receiving on quit")
	default:
	}

	if v, ok := <-ch; ok {
		t.Error("Channel should be closed")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}

	// Test that closing the quit channel does not affect the values channel.
	l.Clear()
	l.Append(1, 2, 3)
	quit = make(chan struct{})
	ch = l.Yield(quit)
	if ch == nil {
		t.Error("Failed to receive channel")
	}

	i = 0
	for v := range ch {
		switch i {
		case 0:
			if v != 1 {
				t.Error("Error receiving 1st item")
				t.Log("\tExpected: 1")
				t.Log("\tReceived:", v)
			}
		case 1:
			if v != 2 {
				t.Error("Error receiving 2nd item")
				t.Log("\tExpected: 2")
				t.Log("\tReceived:", v)
			}
		case 2:
			if v != 3 {
				t.Error("Error receiving 3rd item")
				t.Log("\tExpected: 3.14")
				t.Log("\tReceived:", v)
			}
		default:
			t.Error("Still receiving items when list should be exhausted")
			t.Log("\tReceived:", v)
		}

		if i == 0 {
			close(quit)
		}

		i++
	}

	if v, ok := <-ch; ok {
		t.Error("Channel should be closed")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}
}

func TestSort(t *testing.T) {
	l := hlist.New()

	// Test floats.
	cmp := func(l, r interface{}) bool {
		return l.(float64) < r.(float64)
	}
	l.Append(6.1, 7.1, 4.1, 2.1, 3.1, 5.1, 1.1, 8.1)
	checkString(t, l, "6.1, 7.1, 4.1, 2.1, 3.1, 5.1, 1.1, 8.1")
	checkLength(t, l, 8)
	if err := l.Sort(cmp); err != nil {
		t.Error(err)
	}
	checkString(t, l, "1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1, 8.1")
	checkLength(t, l, 8)

	// Test slices.
	cmp = func(l, r interface{}) bool {
		ls := l.([]byte)
		rs := r.([]byte)
		for i, v := range ls {
			if i == len(rs) {
				// If we are equal up to this point, then r is a prefix of l.
				return false
			}
			if v == rs[i] {
				continue
			} else if v < rs[i] {
				return true
			} else {
				return false
			}
		}

		// If we got here, then either l is a prefix of r or both slices are equal.
		return true
	}
	l.Clear()
	l.Append([]byte{0x0A, 0xBB}, []byte{0x0A, 0xAA}, []byte{0x0A}, []byte{0x01, 0x02, 0x03, 0x04}, []byte{0x0A, 0xAA, 0x00})
	checkString(t, l, "[10 187], [10 170], [10], [1 2 3 4], [10 170 0]")
	checkLength(t, l, 5)
	if err := l.Sort(cmp); err != nil {
		t.Error(err)
	}
	checkString(t, l, "[1 2 3 4], [10], [10 170], [10 170 0], [10 187]")
	checkLength(t, l, 5)
}

func TestSortInt(t *testing.T) {
	// Test a power of two.
	l := hlist.New()
	l.Append(8, 7, 6, 5, 4, 3, 2, 1)
	checkString(t, l, "8, 7, 6, 5, 4, 3, 2, 1")
	checkLength(t, l, 8)
	if err := l.SortInt(); err != nil {
		t.Error(err)
	}
	checkString(t, l, "1, 2, 3, 4, 5, 6, 7, 8")
	checkLength(t, l, 8)

	// Test a block with less than one full stack.
	l.Clear()
	l.Append(9, 8, 7, 6, 5, 4, 3, 2, 1)
	checkString(t, l, "9, 8, 7, 6, 5, 4, 3, 2, 1")
	checkLength(t, l, 9)
	if err := l.SortInt(); err != nil {
		t.Error(err)
	}
	checkString(t, l, "1, 2, 3, 4, 5, 6, 7, 8, 9")
	checkLength(t, l, 9)

	// Test a block with more than one full stack but less than a full block
	l.Clear()
	l.Append(11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	checkString(t, l, "11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1")
	checkLength(t, l, 11)
	if err := l.SortInt(); err != nil {
		t.Error(err)
	}
	checkString(t, l, "1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11")
	checkLength(t, l, 11)

	// Test small list sizes.
	l.Clear()
	l.Append(1, 0)
	checkString(t, l, "1, 0")
	checkLength(t, l, 2)
	if err := l.SortInt(); err != nil {
		t.Error(err)
	}
	checkString(t, l, "0, 1")
	checkLength(t, l, 2)

	l.Clear()
	l.Append(1)
	checkString(t, l, "1")
	checkLength(t, l, 1)
	if err := l.SortInt(); err != nil {
		t.Error(err)
	}
	checkString(t, l, "1")
	checkLength(t, l, 1)

	// Test an empty list.
	l.Clear()
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)
	if err := l.SortInt(); err != nil {
		t.Error(err)
	}
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)
}

func TestSortStr(t *testing.T) {
	// Test uniform length and no same characters.
	l := hlist.New()
	l.Append("ccc", "aaa", "bbb")
	checkString(t, l, "ccc, aaa, bbb")
	checkLength(t, l, 3)
	if err := l.SortStr(); err != nil {
		t.Error(err)
	}
	checkString(t, l, "aaa, bbb, ccc")
	checkLength(t, l, 3)

	// Test variable lengths and no same characters.
	l.Clear()
	l.Append("aaaaaaaaaaa", "bbb", "ccccc")
	checkString(t, l, "aaaaaaaaaaa, bbb, ccccc")
	checkLength(t, l, 3)
	if err := l.SortStr(); err != nil {
		t.Error(err)
	}
	checkString(t, l, "aaaaaaaaaaa, bbb, ccccc")
	checkLength(t, l, 3)

	// Test uniform length and similar characters.
	l.Clear()
	l.Append("cabc", "caab", "abab")
	checkString(t, l, "cabc, caab, abab")
	checkLength(t, l, 3)
	if err := l.SortStr(); err != nil {
		t.Error(err)
	}
	checkString(t, l, "abab, caab, cabc")
	checkLength(t, l, 3)

	// Test variable length and similar characters.
	l.Clear()
	l.Append("cababcabbcbababca", "cabacbcabacb", "cabababacba")
	checkString(t, l, "cababcabbcbababca, cabacbcabacb, cabababacba")
	checkLength(t, l, 3)
	if err := l.SortStr(); err != nil {
		t.Error(err)
	}
	checkString(t, l, "cabababacba, cababcabbcbababca, cabacbcabacb")
	checkLength(t, l, 3)

	// Test subsets.
	l.Clear()
	l.Append("cabcdd", "cabc", "cabab")
	checkString(t, l, "cabcdd, cabc, cabab")
	checkLength(t, l, 3)
	if err := l.SortStr(); err != nil {
		t.Error(err)
	}
	checkString(t, l, "cabab, cabc, cabcdd")
	checkLength(t, l, 3)

	// Test Unicode characters. The following runes will be appended to the list in this order:
	// "U+4444 U+3333", "U+1111 U+2222", "U+1111 U+2222 U+4444 U+3333", "U+1100", "U+4444"
	l.Clear()
	l.Append("䑄㌳", "ᄑ∢", "ᄑ∢䑄㌳", "ᄀ", "䑄")
	checkString(t, l, "䑄㌳, ᄑ∢, ᄑ∢䑄㌳, ᄀ, 䑄")
	checkLength(t, l, 5)
	if err := l.SortStr(); err != nil {
		t.Error(err)
	}
	checkString(t, l, "ᄀ, ᄑ∢, ᄑ∢䑄㌳, 䑄, 䑄㌳")
	checkLength(t, l, 5)

	// Test Unicode characters and ASCII characters. The following runes will be appended to the
	// list in this order: "U+4444 U+3333", "U+1111 U+2222", "A U+1111 U+2222 U+4444 U+3333",
	// "U+1111 U+2222 U+4444 U+3333", "U+1100", "U+4444", "U+1111 U+2222 A"
	l.Clear()
	l.Append("䑄㌳", "ᄑ∢", "Aᄑ∢䑄㌳", "ᄑ∢䑄㌳", "ᄀ", "䑄", "ᄑ∢A")
	checkString(t, l, "䑄㌳, ᄑ∢, Aᄑ∢䑄㌳, ᄑ∢䑄㌳, ᄀ, 䑄, ᄑ∢A")
	checkLength(t, l, 7)
	if err := l.SortStr(); err != nil {
		t.Error(err)
	}
	checkString(t, l, "Aᄑ∢䑄㌳, ᄀ, ᄑ∢, ᄑ∢A, ᄑ∢䑄㌳, 䑄, 䑄㌳")
	checkLength(t, l, 7)
}

func checkString(t *testing.T, l *hlist.List, want string) {
	if l.String() != want {
		t.Error("List contents are incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", l)
	}
}

func checkLength(t *testing.T, l *hlist.List, want int) {
	if l.Length() != want {
		t.Error("Incorrect length")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", l.Length())
	}
}
