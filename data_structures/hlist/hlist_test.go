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
	if l.Same(New()) {
		t.Error("unexpectedly passed Same() test with bad pointer")
	}

	// Test Twin().
	if l.Twin(New()) {
		t.Error("unexpectedly passed Twin() test with bad pointer")
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

func TestExists(t *testing.T) {
	l := New()
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

	// Try to find a non-existant item.
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

	// Try to remove a non-existant item.
	l.RemoveMatch(10)
	checkString(t, l, "[4 5]")
	checkLength(t, l, 1)

	l.RemoveMatch([]int{4, 5})
	checkString(t, l, "<empty>")
	checkLength(t, l, 0)
}

func TestCopy(t *testing.T) {
	l := New()

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
	if l.head == nl.head {
		t.Error("Lists are clones, not copies")
	}

	// Make sure the two lists are separate.
	checkString(t, l, "sizzle, 100000, 3.1415, 15")
	checkLength(t, l, 4)
	checkString(t, nl, "sizzle, 100000, 3.1415, 15")
	checkLength(t, nl, 4)
	if l.head == nl.head {
		t.Error("Lists are clones, not copies")
	}

	l.Remove(1)
	checkString(t, l, "sizzle, 3.1415, 15")
	checkLength(t, l, 3)
	checkString(t, nl, "sizzle, 100000, 3.1415, 15")
	checkLength(t, nl, 4)
	if l.head == nl.head {
		t.Error("Lists are clones, not copies")
	}

	nl.Remove(2)
	checkString(t, l, "sizzle, 3.1415, 15")
	checkLength(t, l, 3)
	checkString(t, nl, "sizzle, 100000, 15")
	checkLength(t, nl, 3)
	if l.head == nl.head {
		t.Error("Lists are clones, not copies")
	}
}

func TestSame(t *testing.T) {
	l1 := New()
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
	l1 = New()
	l1.Append("apple", "banana", "carrot")
	l2 = New()
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
	l1 := New()
	l1.Append("apple", "banana", "carrot")
	l2 := New()
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
		r rune; i []int; n int
	}
	l1 = New()
	l1.Append("apple", []float32{1.23, 2.34, 3.45}, s{'a', []int{8, 9}, 1e3})
	l2 = New()
	l2.Append("apple", []float32{1.23, 2.34, 3.45}, s{'a', []int{8, 9}, 1e3})
	if !l1.Twin(l2) {
		t.Error("Lists unexpectedly failed twin test (test #5)")
	}
	if !l2.Twin(l1) {
		t.Error("Lists unexpectedly failed twin test (test #6)")
	}

	// Tweak l2 a bit.
	l2 = New()
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
	l := New()
	l.Append(0, 1, 2, 3, 4)
	checkString(t, l, "0, 1, 2, 3, 4")
	checkLength(t, l, 5)

	nl := New()
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

	var lp *List
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
	l := New()

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

func TestSort(t *testing.T) {
	l := New()

	// Test floats.
	cmp := func(l, r interface{}) bool {
		if l.(float64) < r.(float64) {
			return true
		}
		return false
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
				continue;
			} else if v < rs[i] {
				return true
			} else {
				return false
			}
		}

		// If we got here, then either l is a prefix of r or both slices are equal.
		return true;
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
	l := New()
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
	l := New()
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

	// Test Unicode characters and ASCII characters. The following runes will be appended to the list in this order:
	// "U+4444 U+3333", "U+1111 U+2222", "A U+1111 U+2222 U+4444 U+3333", "U+1111 U+2222 U+4444 U+3333", "U+1100", "U+4444", "U+1111 U+2222 A"
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
