package htable

import (
	"testing"
	"fmt"
)


// --- Blanket Tests ---
func TestTBadPtr(t *testing.T) {
	var tb *Table

    // Test Add().
	if err := tb.Add("a", "b", "c"); err == nil {
		t.Error("Unexpectedly passed bad pointer test for Add()")
	}

    // Test Insert().
	if err := tb.Insert(1, "a", "b", "c"); err == nil {
		t.Error("Unexpectedly passed bad pointer test for Insert()")
	}

    // Test AddRow().
	if err := tb.AddRow(NewRow(1, 2, 3)); err == nil {
		t.Error("Unexpectedly passed bad pointer test for AddRow()")
	}

    // Test InsertRow().
	if err := tb.InsertRow(0, NewRow(1, 2, 3)); err == nil {
		t.Error("Unexpectedly passed bad pointer test for InsertRow()")
	}

	// Test RemoveRow().
	if err := tb.RemoveRow(2); err == nil {
		t.Error("Unexpectedly passed bad pointer test for RemoveRow()")
	}

    // Test Rows().
	if n := tb.Rows(); n != -1 {
		t.Error("Unexpectedly passed bad pointer test for Rows()")
	}

    // Test Columns().
	if n := tb.Columns(); n != -1 {
		t.Error("Unexpectedly passed bad pointer test for Columns()")
	}

    // Test Count().
	if n := tb.Count(); n != -1 {
		t.Error("Unexpectedly passed bad pointer test for Count()")
	}

	// Test String()
	if s := tb.String(); s != "<nil>" {
		t.Error("Unexpectedly passed bad pointer test for String()")
	}

    // Test Row().
	if i, r := tb.Row("a", 5); i != -1 || r != nil {
		t.Error("Unexpectedly passed bad pointer test for Row()")
	}

    // Test Item().
	if v := tb.Item(5, "a"); v != nil {
		t.Error("Unexpectedly passed bad pointer test for Item()")
	}

	// Test Matches().
	if tb.Matches(0, "a", "item") {
		t.Error("Unexpectedly passed bad pointer test for Matches()")
	}

	// Test Toggle().
	if err := tb.Toggle(0, false); err == nil {
		t.Error("Unexpectedly passed bad pointer test for Toggle()")
	}
}

func TestRBadPtr(t *testing.T) {
	var r *Row

	if v := r.Item("1"); v != nil {
		t.Error("Unexpectedly passed bad pointer test for Row's Item()")
	}

	if r.Matches("1", 1) {
		t.Error("Unexpectedly passed bad pointer test for Row's Matches()")
	}

	if r.String() != "<nil>" {
		t.Error("Unexpectedly passed bad pointer test for Row's String()")
	}

	if r.Count() != -1 {
		t.Error("Unexpectedly passed bad pointer test for Row's Count()")
	}
}

func TestTBadArgs(t *testing.T) {
	tb, _ := New("1", "2", "3")
	tb.Add(1, 2, 3)

	// Test New() - missing column header.
	if tbl, err := New(""); tbl != nil || err == nil {
		t.Error("Unexpectedly passed missing column header test for New()")
	}

	// Test Add() - too few columns.
	if err := tb.Add(1, 2); err == nil {
		t.Error("Unexpectedly passed too few columns test for Add()")
	}

	// Test Add() - too many columns.
	if err := tb.Add(1, 2, 3, 4); err == nil {
		t.Error("Unexpectedly passed too many columns test for Add()")
	}

	// Test Add() - wrong type in first position.
	if err := tb.Add("item1", 5, 6); err == nil {
		t.Error("Unexpectedly passed wrong type (first) test for Add()")
	}

	// Test Add() - wrong type in middle position.
	if err := tb.Add(4, "item5", 6); err == nil {
		t.Error("Unexpectedly passed wrong type (middle) test for Add()")
	}

	// Test Add() - wrong type in last position.
	if err := tb.Add(4, 5, "item6"); err == nil {
		t.Error("Unexpectedly passed wrong type (last) test for Add()")
	}

    // Test Insert() - negative index.
	if err := tb.Insert(-1, 4, 5, 6); err == nil {
		t.Error("Unexpectedly passed negative index test for Insert()")
	}

    // Test Insert() - out-of-bounds index.
	if err := tb.Insert(100, 4, 5, 6); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test for Insert()")
	}

    // Test RemoveRow() - negative index.
	if err := tb.RemoveRow(-1); err == nil {
		t.Error("Unexpectedly passed negative index test for RemoveRow()")
	}

    // Test RemoveRow() - out-of-bounds index.
	if err := tb.RemoveRow(100); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test for RemoveRow()")
	}

    // Test Row() - empty column header.
	if i, r := tb.Row("", 5); i != -1 || r != nil {
		t.Error("Unexpectedly passed empty column header test for Row()")
	}

    // Test Row() - invalid column header.
	if i, r := tb.Row("4", 4); i != -1 || r != nil {
		t.Error("Unexpectedly passed invalid column header test for Row()")
	}

    // Test Row() - missing item.
	if i, r := tb.Row("1", nil); i != -1 || r != nil {
		t.Error("Unexpectedly passed missing item test for Row()")
	}

    // Test Item() - negative index.
	if v := tb.Item(-1, "1"); v != nil {
		t.Error("Unexpectedly passed negative index test for Item()")
	}

    // Test Item() - out-of-bounds index.
	if v := tb.Item(100, "1"); v != nil {
		t.Error("Unexpectedly passed out-of-bounds index test for Item()")
	}

    // Test Item() - empty column header.
	if v := tb.Item(0, ""); v != nil {
		t.Error("Unexpectedly passed empty column header test for Item()")
	}

    // Test Item() - invalid column header.
	if v := tb.Item(0, "4"); v != nil {
		t.Error("Unexpectedly passed invalid column header test for Item()")
	}

    // Test Matches() - negative index.
	if tb.Matches(-1, "1", 1) {
		t.Error("Unexpectedly passed negative index test for Matches()")
	}

    // Test Matches() - out-of-bounds index.
	if tb.Matches(100, "1", 1) {
		t.Error("Unexpectedly passed out-of-bounds index test for Matches()")
	}

    // Test Matches() - empty column header.
	if tb.Matches(0, "", 1) {
		t.Error("Unexpectedly passed empty column header test for Matches()")
	}

    // Test Matches() - invalid column header.
	if tb.Matches(0, "4", 4) {
		t.Error("Unexpectedly passed invalid column header test for Matches()")
	}

    // Test Matches() - missing item.
	if tb.Matches(0, "1", nil) {
		t.Error("Unexpectedly passed missing item test for Matches()")
	}
}

func TestRBadArgs(t *testing.T) {
	r := NewRow(1, 2, 3)

	// Test Item() - no column headers.
	if v := r.Item("1"); v != nil {
		t.Error("Unexpectedly passed no column headers test for Item()")
	}

	// Test Item() - missing column header.
	if v := r.Item(""); v != nil {
		t.Error("Unexpectedly passed missing column header test for Item()")
	}

	// Test Matches() - no column headers.
	if r.Matches("1", 1) {
		t.Error("Unexpectedly passed no column headers test for Matches()")
	}

	// Test Matches() - missing column header.
	if r.Matches("", 1) {
		t.Error("Unexpectedly passed missing column header test for Matches()")
	}
}

func TestBadTypes(t *testing.T) {
	// TODO
}


// --- Function Tests ---
func TestTNew(t *testing.T) {
	n, err := New("a", "b", "c")
	if err != nil {
		t.Error(err)
	} else if n == nil {
		t.Error("Failed to create new table")
	}
}

func TestRNewRow(t *testing.T) {
	r := NewRow(1, 2, 3)
	checkRString(t, r, "{1, 2, 3}")
	checkRCount(t, r, 3)

	type s struct {
		s1 string
		s2 string
	}
	r = NewRow(1, 2, s{s1:"a", s2:"b"})
	checkRString(t, r, "{1, 2, {a b}}")
	checkRCount(t, r, 3)

	tmp := NewRow("other", "row")
	r = NewRow("this row", "here", 5, tmp)
	checkRString(t, r, "{this row, here, 5, {other, row}}")
	checkRCount(t, r, 4)
}


// --- Table's Method Tests ---
func TestTAdd(t *testing.T) {
	// Testing rows of integers.
	tb, _ := New("1", "2", "3")
	if err := tb.Add(1, 2, 3); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}")
	checkTCount(t, tb, 3)

	if err := tb.Add(4, 5, 6); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCount(t, tb, 6)

	// Test rows of characters.
	tb, _ = New("1", "2", "3")
	if err := tb.Add('a', 'b', 'c'); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 97, 2: 98, 3: 99}")
	checkTCount(t, tb, 3)

	if err := tb.Add('d', 'e', 'f'); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 97, 2: 98, 3: 99}, {1: 100, 2: 101, 3: 102}")
	checkTCount(t, tb, 6)

	// Test rows of int slices.
	tb, _ = New("1", "2", "3")
	if err := tb.Add([]int{10, 20}, []int{30, 40}, []int{50, 60}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: [10 20], 2: [30 40], 3: [50 60]}")
	checkTCount(t, tb, 3)

	if err := tb.Add([]int{100, 200}, []int{300, 400}, []int{500, 600}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: [10 20], 2: [30 40], 3: [50 60]}, {1: [100 200], 2: [300 400], 3: [500 600]}")
	checkTCount(t, tb, 6)

	// Test mixed-value rows.
	tb, _ = New("1", "2", "3")
	if err := tb.Add(1.1, "b", []byte{0x03}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1.1, 2: b, 3: [3]}")
	checkTCount(t, tb, 3)

	if err := tb.Add(4.4, "e", []byte{0x06}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1.1, 2: b, 3: [3]}, {1: 4.4, 2: e, 3: [6]}")
	checkTCount(t, tb, 6)

	// Test rows of strings with spaces.
	tb, _ = New("1", "2", "3")
	if err := tb.Add("before ", "in between", " after"); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: before , 2: in between, 3:  after}")
	checkTCount(t, tb, 3)

	if err := tb.Add("AAA", "   ", ""); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: before , 2: in between, 3:  after}, {1: AAA, 2:    , 3: }")
	checkTCount(t, tb, 6)

	// Test rows masquerading as interfaces.
	tb, _ = New("1", "2", "3")
	if err := tb.Add(interface{}(1), interface{}(2), interface{}(3)); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}")
	checkTCount(t, tb, 3)

	if err := tb.Add(4, 5, 6); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCount(t, tb, 6)

	// Test rows of functions.
	tb, _ = New("1", "2", "3")
	if err := tb.Add(TestTNew, checkTCount, TestTInsert); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, fmt.Sprintf("{1: %p, 2: %p, 3: %p}", TestTNew, checkTCount, TestTInsert))
	checkTCount(t, tb, 3)

	if err := tb.Add(checkTString, checkTString, TestTAdd); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, fmt.Sprintf("{1: %p, 2: %p, 3: %p}, {1: %p, 2: %p, 3: %p}", TestTNew, checkTCount, TestTInsert, checkTString, checkTString, TestTAdd))
	checkTCount(t, tb, 6)
}

func TestTInsert(t *testing.T) {
	// Test inserting a row at the beginning.
	tb, _ := New("1", "2", "3")
	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCount(t, tb, 6)
	if err := tb.Insert(0, -1, -2, -3); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: -1, 2: -2, 3: -3}, {1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCount(t, tb, 9)

	// Test inserting a row in the middle.
	tb, _ = New("1", "2", "3")
	tb.Add("a", "b", "c")
	tb.Add("d", "e", "f")
	checkTString(t, tb, "{1: a, 2: b, 3: c}, {1: d, 2: e, 3: f}")
	checkTCount(t, tb, 6)
	if err := tb.Insert(1, "x", "y", "z"); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: a, 2: b, 3: c}, {1: x, 2: y, 3: z}, {1: d, 2: e, 3: f}")
	checkTCount(t, tb, 9)

	// Test inserting a row at the end.
	tb, _ = New("1", "2", "3")
	tb.Add([]int{10, 20}, []int{30, 40}, []int{50, 60})
	tb.Add([]int{100, 200}, []int{300, 400}, []int{500, 600})
	checkTString(t, tb, "{1: [10 20], 2: [30 40], 3: [50 60]}, {1: [100 200], 2: [300 400], 3: [500 600]}")
	checkTCount(t, tb, 6)
	if err := tb.Insert(2, []int{-1, -2, -3}, []int{-4, -5, -6}, []int{-7, -8, -9}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: [10 20], 2: [30 40], 3: [50 60]}, {1: [100 200], 2: [300 400], 3: [500 600]}, {1: [-1 -2 -3], 2: [-4 -5 -6], 3: [-7 -8 -9]}")
	checkTCount(t, tb, 9)

	// Test inserting a row beyond the table's current boundaries.
	tb, _ = New("1", "2", "3")
	tb.Add(1.1, "b", []byte{0x03})
	tb.Add(4.4, "e", []byte{0x06})
	checkTString(t, tb, "{1: 1.1, 2: b, 3: [3]}, {1: 4.4, 2: e, 3: [6]}")
	checkTCount(t, tb, 6)
	if err := tb.Insert(3, 7.7, "h", []byte{0x09}); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test")
	}
	checkTString(t, tb, "{1: 1.1, 2: b, 3: [3]}, {1: 4.4, 2: e, 3: [6]}")
	checkTCount(t, tb, 6)

	// Test only inserting instead of appending.
	tb, _ = New("1", "2", "3")
	if err := tb.Insert(0, "before ", "in between", " after"); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: before , 2: in between, 3:  after}")
	checkTCount(t, tb, 3)
	if err := tb.Insert(1, "AAA", "   ", ""); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: before , 2: in between, 3:  after}, {1: AAA, 2:    , 3: }")
	checkTCount(t, tb, 6)
	if err := tb.Insert(0, "first", "second", "third"); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: first, 2: second, 3: third}, {1: before , 2: in between, 3:  after}, {1: AAA, 2:    , 3: }")
	checkTCount(t, tb, 9)
}

func TestTAddRow(t *testing.T) {
	// TODO
}

func TestTInsertRow(t *testing.T) {
	// TODO
}

func TestTRemoveRow(t *testing.T) {
	// Test removing a row at the beginning.
	tb, _ := New("1", "2", "3")
	tb.Add(-1, -2, -3)
	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)
	checkTString(t, tb, "{1: -1, 2: -2, 3: -3}, {1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCount(t, tb, 9)
	if err := tb.RemoveRow(0); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCount(t, tb, 6)

	// Test removing a row in the middle.
	tb, _ = New("1", "2", "3")
	tb.Add("a", "b", "c")
	tb.Add("x", "y", "z")
	tb.Add("d", "e", "f")
	checkTString(t, tb, "{1: a, 2: b, 3: c}, {1: x, 2: y, 3: z}, {1: d, 2: e, 3: f}")
	checkTCount(t, tb, 9)
	if err := tb.RemoveRow(1); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: a, 2: b, 3: c}, {1: d, 2: e, 3: f}")
	checkTCount(t, tb, 6)

	// Test removing a row at the end.
	tb, _ = New("1", "2", "3")
	tb.Add([]int{10, 20}, []int{30, 40}, []int{50, 60})
	tb.Add([]int{100, 200}, []int{300, 400}, []int{500, 600})
	tb.Add([]int{-1, -2, -3}, []int{-4, -5, -6}, []int{-7, -8, -9})
	checkTString(t, tb, "{1: [10 20], 2: [30 40], 3: [50 60]}, {1: [100 200], 2: [300 400], 3: [500 600]}, {1: [-1 -2 -3], 2: [-4 -5 -6], 3: [-7 -8 -9]}")
	checkTCount(t, tb, 9)
	if err := tb.RemoveRow(2); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: [10 20], 2: [30 40], 3: [50 60]}, {1: [100 200], 2: [300 400], 3: [500 600]}")
	checkTCount(t, tb, 6)

	// Test removing a row beyond the table's current boundaries.
	tb, _ = New("1", "2", "3")
	tb.Add(1.1, "b", []byte{0x03})
	tb.Add(4.4, "e", []byte{0x06})
	checkTString(t, tb, "{1: 1.1, 2: b, 3: [3]}, {1: 4.4, 2: e, 3: [6]}")
	checkTCount(t, tb, 6)
	if err := tb.RemoveRow(3); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test")
	}
	checkTString(t, tb, "{1: 1.1, 2: b, 3: [3]}, {1: 4.4, 2: e, 3: [6]}")
	checkTCount(t, tb, 6)

	// Test removing when there aren't any rows in the table.
	tb, _ = New("1", "2", "3")
	checkTString(t, tb, "<empty>")
	checkTCount(t, tb, 0)
	if err := tb.RemoveRow(0); err == nil {
		t.Error("Unexpectedly passed removing from empty table test")
	}
	checkTString(t, tb, "<empty>")
	checkTCount(t, tb, 0)
}

func TestTRows(t *testing.T) {
	tb, _ := New("1", "2", "3")
	if n := tb.Rows(); n != 0 {
		t.Error("Row count is incorrect")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", n)
	}

	tb.Add(1, 2, 3)
	if n := tb.Rows(); n != 1 {
		t.Error("Row count is incorrect")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", n)
	}

	tb.Add(4, 5, 6)
	if n := tb.Rows(); n != 2 {
		t.Error("Row count is incorrect")
		t.Log("\tExpected:", 2)
		t.Log("\tReceived:", n)
	}

	tb.RemoveRow(0)
	if n := tb.Rows(); n != 1 {
		t.Error("Row count is incorrect")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", n)
	}

	tb.Insert(1, 7, 8, 9)
	if n := tb.Rows(); n != 2 {
		t.Error("Row count is incorrect")
		t.Log("\tExpected:", 2)
		t.Log("\tReceived:", n)
	}
}

func TestTColumns(t *testing.T) {
	tb, _ := New("1", "2", "3")
	if n := tb.Columns(); n != 3 {
		t.Error("Column count is incorrect")
		t.Log("\tExpected:", 3)
		t.Log("\tReceived:", n)
	}

	tb.Add(1, 2, 3)
	if n := tb.Columns(); n != 3 {
		t.Error("Column count is incorrect")
		t.Log("\tExpected:", 3)
		t.Log("\tReceived:", n)
	}

	tb.Add(4, 5, 6)
	if n := tb.Columns(); n != 3 {
		t.Error("Column count is incorrect")
		t.Log("\tExpected:", 3)
		t.Log("\tReceived:", n)
	}

	tb.RemoveRow(0)
	if n := tb.Columns(); n != 3 {
		t.Error("Column count is incorrect")
		t.Log("\tExpected:", 3)
		t.Log("\tReceived:", n)
	}

	tb.Insert(1, 7, 8, 9)
	if n := tb.Columns(); n != 3 {
		t.Error("Column count is incorrect")
		t.Log("\tExpected:", 3)
		t.Log("\tReceived:", n)
	}

	tb, _ = New("1")
	if n := tb.Columns(); n != 1 {
		t.Error("Column count is incorrect")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", n)
	}

	tb.Add(1)
	if n := tb.Columns(); n != 1 {
		t.Error("Column count is incorrect")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", n)
	}

	tb.RemoveRow(0)
	if n := tb.Columns(); n != 1 {
		t.Error("Column count is incorrect")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", n)
	}
}

func TestTCount(t *testing.T) {
	tb, _ := New("1", "2", "3")
	checkTCount(t, tb, 0)

	tb.Add(1, 2, 3)
	checkTCount(t, tb, 3)

	tb.Add(4, 5, 6)
	checkTCount(t, tb, 6)

	tb.RemoveRow(0)
	checkTCount(t, tb, 3)

	tb.Insert(1, 7, 8, 9)
	checkTCount(t, tb, 6)

	tb, _ = New("1")
	checkTCount(t, tb, 0)

	tb.Add(1)
	checkTCount(t, tb, 1)

	tb.RemoveRow(0)
	checkTCount(t, tb, 0)
}

func TestTRowItem(t *testing.T) {
	// Set up a new table.
	tb, _ := New("1", "2", "3")
	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)
	tb.Add(7, 8, 9)

	// Validate the row of each of the 9 values.
	cols := []string{"1", "2", "3"}
	for row := 0; row < 3; row++ {
		for i, col := range cols {
			v := (row * 3) + i + 1
			if i, _ := tb.Row(col, v); i != row {
				t.Error("Matched incorrect row")
				t.Log("\tExpected:", row)
				t.Log("\tReceived:", i)
			}
		}
	}

	// Validate each of the 9 items.
	for row := 0; row < 3; row++ {
		for i, col := range cols {
			exp := (row * 3) + i + 1
			if v := tb.Item(row, col); v != exp {
				t.Error("Item value incorrect")
				t.Log("\tExpected:", exp)
				t.Log("\tReceived:", v)
			}
		}
	}

	// Set up a new table.
	tb, _ = New("name", "left-handed", "age")
	tb.Add("Swari", true,  30)
	tb.Add("Kathy", false, 40)
	tb.Add("Joe",   false, 189)

	// Find the first person who is left-handed.
	if i, _ := tb.Row("left-handed", true); i != 0 {
		t.Error("Matched incorrect row")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", i)
	} else if v := tb.Item(i, "name"); v != "Swari" {
		t.Error("Item value incorrect")
		t.Log("\tExpected:", "Swari")
		t.Log("\tReceived:", v)
	}

	// Find the first person who is right-handed.
	if i, _ := tb.Row("left-handed", false); i != 1 {
		t.Error("Matched incorrect row")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", i)
	} else if v := tb.Item(i, "name"); v != "Kathy" {
		t.Error("Item value incorrect")
		t.Log("\tExpected:", "Kathy")
		t.Log("\tReceived:", v)
	}

	// Look for a name that doesn't exist.
	if i, _ := tb.Row("name", "Dr. Gutten"); i != -1 {
		t.Error("Unexpectedly found row")
		t.Log("\tExpected:", -1)
		t.Log("\tReceived:", i)
	}

	// Look for an item that doesn't exist.
	if i, _ := tb.Row("name", "Kathy"); i != 1 {
		t.Error("Matched incorrect row")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", i)
	} else if v := tb.Item(i, "height"); v != nil {
		t.Error("Unexpectedly found item")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}

	// Store and run functions.
	hi := func() string { return "hello" }
	dbl := func(n int) int { return n * 2 }

	tb, _ = New("name", "func")
	tb.Add("hi", hi)
	tb.Add("dbl", dbl)

	// Find and run the first function.
	if i, _ := tb.Row("name", "hi"); i != 0 {
		t.Error("Matched incorrect row")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", i)
	} else {
		f := tb.Item(i, "func")
		if ff, ok := f.(func() string); !ok {
			t.Error("Returned function is not of type func() string")
		} else if v := ff(); v != "hello" {
			t.Error("Function result incorrect")
			t.Log("\tExpected:", "hello" )
			t.Log("\tReceived:", v)
		}
	}

	// Find and run the second function.
	if i, _ := tb.Row("name", "dbl"); i != 1 {
		t.Error("Matched incorrect row")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", i)
	} else {
		f := tb.Item(i, "func")
		if ff, ok := f.(func(int) int); !ok {
			t.Error("Returned function is not of type func(int) int")
		} else if n := ff(2); n != 4 {
			t.Error("Function result incorrect")
			t.Log("\tExpected:", 4 )
			t.Log("\tReceived:", n)
		}
	}
}

func TestTMatches(t *testing.T) {
	// Set up a new table.
	tb, _ := New("1", "2", "3")
	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)
	tb.Add(7, 8, 9)

	// Make sure we have a match at each position.
	cols := []string{"1", "2", "3"}
	for row := 0; row < 3; row++ {
		for i, col := range cols {
			v := (row * 3) + i + 1
			if !tb.Matches(row, col, v) {
				t.Error("Did not match")
			}
		}
	}

	// Set up a new table.
	tb, _ = New("name", "left-handed", "age")
	tb.Add("Swari", true,  30)
	tb.Add("Kathy", false, 40)
	tb.Add("Joe",   false, 189)

	// Make sure Swari is left-handed.
	if !tb.Matches(0, "left-handed", true) {
		t.Error("Did not match")
	}

	// Make sure Joe is 189.
	if !tb.Matches(2, "age", 189) {
		t.Error("Did not match")
	}

	// Look for a name that doesn't exist.
	if tb.Matches(0, "name", "April") {
		t.Error("Unexpectedly matched")
	}
}

func TestTToggle(t *testing.T) {
	// TODO
}


// --- Row's Method Tests ---
func TestRItem(t *testing.T) {
	// TODO
}

func TestRMatches(t *testing.T) {
	// TODO
}


// --- Helper Functions ---
func checkTString(t *testing.T, tb *Table, want string) {
	if s := tb.String(); s != want {
		t.Error("Table items are incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", s)
	}
}

func checkTCount(t *testing.T, tb *Table, want int) {
	if n := tb.Count(); n != want {
		t.Error("Table count is incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", n)
	}
}

func checkRString(t *testing.T, r *Row, want string) {
	if s := r.String(); s != want {
		t.Error("Table items are incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", s)
	}
}

func checkRCount(t *testing.T, r *Row, want int) {
	if n := r.Count(); n != want {
		t.Error("Table count is incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", n)
	}
}
