package htable

import (
	"testing"
	"fmt"
)


// --- Blanket Tests ---
func TestBadPtr(t *testing.T) {
	var tb *Table

    // Test AddRow().
	if err := tb.AddRow("a", "b", "c"); err == nil {
		t.Error("Unexpectedly passed bad pointer test for AddRow()")
	}

    // Test InsertRow().
	if err := tb.InsertRow(1, "a", "b", "c"); err == nil {
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

    // Test Row().
	if n := tb.Row("a", 5); n != -1 {
		t.Error("Unexpectedly passed bad pointer test for Row()")
	}

	// Test String()
	if s := tb.String(); s != "<nil>" {
		t.Error("Unexpectedly passed bad pointer test for String()")
	}
}

func TestBadArgs(t *testing.T) {
	tb, _ := New("1", "2", "3")
	tb.AddRow(1, 2, 3)

	// Test New() - missing column header.
	if tbl, err := New(""); tbl != nil || err == nil {
		t.Error("Unexpectedly passed missing column header test for New()")
	}

	// Test AddRow() - too few columns.
	if err := tb.AddRow(1, 2); err == nil {
		t.Error("Unexpectedly passed too few columns test for AddRow()")
	}

	// Test AddRow() - too many columns.
	if err := tb.AddRow(1, 2, 3, 4); err == nil {
		t.Error("Unexpectedly passed too many columns test for AddRow()")
	}

	// Test AddRow() - wrong types.
	if err := tb.AddRow("item1", "item2", "item3"); err == nil {
		t.Error("Unexpectedly passed wrong types test for AddRow()")
	}

    // Test InsertRow() - negative index.
	if err := tb.InsertRow(-1, "item"); err == nil {
		t.Error("Unexpectedly passed negative index test for InsertRow()")
	}

    // Test InsertRow() - out-of-bounds index.
	if err := tb.InsertRow(100, "item"); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test for InsertRow()")
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
	if r := tb.Row("", 5); r != -1 {
		t.Error("Unexpectedly passed empty column header test for Row()")
	}

    // Test Row() - missing item.
	if r := tb.Row("header", nil); r != -1 {
		t.Error("Unexpectedly passed missing item test for Row()")
	}
}

func TestBadTypes(t *testing.T) {
}


// --- Function Tests ---
func TestNew(t *testing.T) {
	n, err := New("a", "b", "c")
	if err != nil {
		t.Error(err)
	} else if n == nil {
		t.Error("Failed to create new table")
	}
}


// --- Method Tests ---
func TestAddRow(t *testing.T) {
	// Testing rows of integers.
	tb, _ := New("1", "2", "3")
	if err := tb.AddRow(1, 2, 3); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[1 2 3]")
	checkCount(t, tb, 3)

	if err := tb.AddRow(4, 5, 6); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[1 2 3], [4 5 6]")
	checkCount(t, tb, 6)

	// Test rows of characters.
	tb, _ = New("1", "2", "3")
	if err := tb.AddRow('a', 'b', 'c'); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[97 98 99]")
	checkCount(t, tb, 3)

	if err := tb.AddRow('d', 'e', 'f'); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[97 98 99], [100 101 102]")
	checkCount(t, tb, 6)

	// Test rows of int slices.
	tb, _ = New("1", "2", "3")
	if err := tb.AddRow([]int{10, 20}, []int{30, 40}, []int{50, 60}); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[[10 20] [30 40] [50 60]]")
	checkCount(t, tb, 3)

	if err := tb.AddRow([]int{100, 200}, []int{300, 400}, []int{500, 600}); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[[10 20] [30 40] [50 60]], [[100 200] [300 400] [500 600]]")
	checkCount(t, tb, 6)

	// Test mixed-value rows.
	tb, _ = New("1", "2", "3")
	if err := tb.AddRow(1.1, "b", []byte{0x03}); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[1.1 b [3]]")
	checkCount(t, tb, 3)

	if err := tb.AddRow(4.4, "e", []byte{0x06}); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[1.1 b [3]], [4.4 e [6]]")
	checkCount(t, tb, 6)

	// Test rows of strings with spaces.
	tb, _ = New("1", "2", "3")
	if err := tb.AddRow("before ", "in between", " after"); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[before  in between  after]")
	checkCount(t, tb, 3)

	if err := tb.AddRow("AAA", "   ", ""); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[before  in between  after], [AAA     ]")
	checkCount(t, tb, 6)

	// Test rows masquerading as interfaces.
	tb, _ = New("1", "2", "3")
	if err := tb.AddRow(interface{}(1), interface{}(2), interface{}(3)); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[1 2 3]")
	checkCount(t, tb, 3)

	if err := tb.AddRow(4, 5, 6); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[1 2 3], [4 5 6]")
	checkCount(t, tb, 6)

	// Test rows of functions.
	tb, _ = New("1", "2", "3")
	if err := tb.AddRow(TestNew, checkCount, TestInsertRow); err != nil {
		t.Error(err)
	}
	checkString(t, tb, fmt.Sprintf("[%p %p %p]", TestNew, checkCount, TestInsertRow))
	checkCount(t, tb, 3)

	if err := tb.AddRow(checkString, checkString, TestAddRow); err != nil {
		t.Error(err)
	}
	checkString(t, tb, fmt.Sprintf("[%p %p %p], [%p %p %p]", TestNew, checkCount, TestInsertRow, checkString, checkString, TestAddRow))
	checkCount(t, tb, 6)
}

func TestInsertRow(t *testing.T) {
	// Test inserting a row at the beginning.
	tb, _ := New("1", "2", "3")
	tb.AddRow(1, 2, 3)
	tb.AddRow(4, 5, 6)
	checkString(t, tb, "[1 2 3], [4 5 6]")
	checkCount(t, tb, 6)
	if err := tb.InsertRow(0, -1, -2, -3); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[-1 -2 -3], [1 2 3], [4 5 6]")
	checkCount(t, tb, 9)

	// Test inserting a row in the middle.
	tb, _ = New("1", "2", "3")
	tb.AddRow("a", "b", "c")
	tb.AddRow("d", "e", "f")
	checkString(t, tb, "[a b c], [d e f]")
	checkCount(t, tb, 6)
	if err := tb.InsertRow(1, "x", "y", "z"); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[a b c], [x y z], [d e f]")
	checkCount(t, tb, 9)

	// Test inserting a row at the end.
	tb, _ = New("1", "2", "3")
	tb.AddRow([]int{10, 20}, []int{30, 40}, []int{50, 60})
	tb.AddRow([]int{100, 200}, []int{300, 400}, []int{500, 600})
	checkString(t, tb, "[[10 20] [30 40] [50 60]], [[100 200] [300 400] [500 600]]")
	checkCount(t, tb, 6)
	if err := tb.InsertRow(2, []int{-1, -2, -3}, []int{-4, -5, -6}, []int{-7, -8, -9}); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[[10 20] [30 40] [50 60]], [[100 200] [300 400] [500 600]], [[-1 -2 -3] [-4 -5 -6] [-7 -8 -9]]")
	checkCount(t, tb, 9)

	// Test inserting a row beyond the table's current boundaries.
	tb, _ = New("1", "2", "3")
	tb.AddRow(1.1, "b", []byte{0x03})
	tb.AddRow(4.4, "e", []byte{0x06})
	checkString(t, tb, "[1.1 b [3]], [4.4 e [6]]")
	checkCount(t, tb, 6)
	if err := tb.InsertRow(3, 7.7, "h", []byte{0x09}); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test")
	}
	checkString(t, tb, "[1.1 b [3]], [4.4 e [6]]")
	checkCount(t, tb, 6)

	// Test only inserting instead of appending.
	tb, _ = New("1", "2", "3")
	if err := tb.InsertRow(0, "before ", "in between", " after"); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[before  in between  after]")
	checkCount(t, tb, 3)
	if err := tb.InsertRow(1, "AAA", "   ", ""); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[before  in between  after], [AAA     ]")
	checkCount(t, tb, 6)
	if err := tb.InsertRow(0, "first", "second", "third"); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[first second third], [before  in between  after], [AAA     ]")
	checkCount(t, tb, 9)
}

func TestRemoveRow(t *testing.T) {
	// Test removing a row at the beginning.
	tb, _ := New("1", "2", "3")
	tb.AddRow(-1, -2, -3)
	tb.AddRow(1, 2, 3)
	tb.AddRow(4, 5, 6)
	checkString(t, tb, "[-1 -2 -3], [1 2 3], [4 5 6]")
	checkCount(t, tb, 9)
	if err := tb.RemoveRow(0); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[1 2 3], [4 5 6]")
	checkCount(t, tb, 6)

	// Test removing a row in the middle.
	tb, _ = New("1", "2", "3")
	tb.AddRow("a", "b", "c")
	tb.AddRow("x", "y", "z")
	tb.AddRow("d", "e", "f")
	checkString(t, tb, "[a b c], [x y z], [d e f]")
	checkCount(t, tb, 9)
	if err := tb.RemoveRow(1); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[a b c], [d e f]")
	checkCount(t, tb, 6)

	// Test removing a row at the end.
	tb, _ = New("1", "2", "3")
	tb.AddRow([]int{10, 20}, []int{30, 40}, []int{50, 60})
	tb.AddRow([]int{100, 200}, []int{300, 400}, []int{500, 600})
	tb.AddRow([]int{-1, -2, -3}, []int{-4, -5, -6}, []int{-7, -8, -9})
	checkString(t, tb, "[[10 20] [30 40] [50 60]], [[100 200] [300 400] [500 600]], [[-1 -2 -3] [-4 -5 -6] [-7 -8 -9]]")
	checkCount(t, tb, 9)
	if err := tb.RemoveRow(2); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[[10 20] [30 40] [50 60]], [[100 200] [300 400] [500 600]]")
	checkCount(t, tb, 6)

	// Test removing a row beyond the table's current boundaries.
	tb, _ = New("1", "2", "3")
	tb.AddRow(1.1, "b", []byte{0x03})
	tb.AddRow(4.4, "e", []byte{0x06})
	checkString(t, tb, "[1.1 b [3]], [4.4 e [6]]")
	checkCount(t, tb, 6)
	if err := tb.RemoveRow(3); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test")
	}
	checkString(t, tb, "[1.1 b [3]], [4.4 e [6]]")
	checkCount(t, tb, 6)

	// Test removing when there aren't any rows in the table.
	tb, _ = New("1", "2", "3")
	checkString(t, tb, "<empty>")
	checkCount(t, tb, 0)
	if err := tb.RemoveRow(0); err == nil {
		t.Error("Unexpectedly passed removing from empty table test")
	}
	checkString(t, tb, "<empty>")
	checkCount(t, tb, 0)
}

func TestRows(t *testing.T) {
	tb, _ := New("1", "2", "3")
	if n := tb.Rows(); n != 0 {
		t.Error("Row count is incorrect")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", n)
	}

	tb.AddRow(1, 2, 3)
	if n := tb.Rows(); n != 1 {
		t.Error("Row count is incorrect")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", n)
	}

	tb.AddRow(4, 5, 6)
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

	tb.InsertRow(1, 7, 8, 9)
	if n := tb.Rows(); n != 2 {
		t.Error("Row count is incorrect")
		t.Log("\tExpected:", 2)
		t.Log("\tReceived:", n)
	}
}

func TestColumns(t *testing.T) {
	tb, _ := New("1", "2", "3")
	if n := tb.Columns(); n != 3 {
		t.Error("Column count is incorrect")
		t.Log("\tExpected:", 3)
		t.Log("\tReceived:", n)
	}

	tb.AddRow(1, 2, 3)
	if n := tb.Columns(); n != 3 {
		t.Error("Column count is incorrect")
		t.Log("\tExpected:", 3)
		t.Log("\tReceived:", n)
	}

	tb.AddRow(4, 5, 6)
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

	tb.InsertRow(1, 7, 8, 9)
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

	tb.AddRow(1)
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

func TestCount(t *testing.T) {
}

func TestRow(t *testing.T) {
}


// --- Helper Functions ---
func checkString(t *testing.T, tb *Table, want string) {
	if s := tb.String(); s != want {
		t.Error("Table items are incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", s)
	}
}

func checkCount(t *testing.T, tb *Table, want int) {
	if n := tb.Count(); n != want {
		t.Error("Table count is incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", n)
	}
}
