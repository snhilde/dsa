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
		t.Error("unexpectedly passed bad pointer test for AddRow()")
	}

    // Test InsertRow().
	if err := tb.InsertRow(1, "a", "b", "c"); err == nil {
		t.Error("unexpectedly passed bad pointer test for InsertRow()")
	}

	// Test RemoveRow().
	if err := tb.RemoveRow(2); err == nil {
		t.Error("unexpectedly passed bad pointer test for RemoveRow()")
	}

    // Test Rows().
	if n := tb.Rows(); n != -1 {
		t.Error("unexpectedly passed bad pointer test for Rows()")
	}

    // Test Columns().
	if n := tb.Columns(); n != -1 {
		t.Error("unexpectedly passed bad pointer test for Columns()")
	}

    // Test Count().
	if n := tb.Count(); n != -1 {
		t.Error("unexpectedly passed bad pointer test for Count()")
	}

    // Test Row().
	if n := tb.Row("a", 5); n != -1 {
		t.Error("unexpectedly passed bad pointer test for Row()")
	}

	// Test String()
	if s := tb.String(); s != "<nil>" {
		t.Error("unexpectedly passed bad pointer test for String()")
	}
}

func TestBadArgs(t *testing.T) {
	tb, _ := New("1", "2", "3")
	tb.AddRow(1, 2, 3)

	// Test New() - missing column header.
	if tbl, err := New(""); tbl != nil || err == nil {
		t.Error("unexpectedly passed missing column header test for New()")
	}

	// Test AddRow() - too few columns.
	if err := tb.AddRow(1, 2); err == nil {
		t.Error("unexpectedly passed too few columns test for AddRow()")
	}

	// Test AddRow() - too many columns.
	if err := tb.AddRow(1, 2, 3, 4); err == nil {
		t.Error("unexpectedly passed too many columns test for AddRow()")
	}

	// Test AddRow() - wrong types.
	if err := tb.AddRow("item1", "item2", "item3"); err == nil {
		t.Error("unexpectedly passed wrong types test for AddRow()")
	}

    // Test InsertRow() - negative index.
	if err := tb.InsertRow(-1, "item"); err == nil {
		t.Error("unexpectedly passed negative index test for InsertRow()")
	}

    // Test InsertRow() - out-of-bounds index.
	if err := tb.InsertRow(100, "item"); err == nil {
		t.Error("unexpectedly passed out-of-bounds index test for InsertRow()")
	}

    // Test RemoveRow() - negative index.
	if err := tb.RemoveRow(-1); err == nil {
		t.Error("unexpectedly passed negative index test for RemoveRow()")
	}

    // Test RemoveRow() - out-of-bounds index.
	if err := tb.RemoveRow(100); err == nil {
		t.Error("unexpectedly passed out-of-bounds index test for RemoveRow()")
	}

    // Test Row() - empty column header.
	if r := tb.Row("", 5); r != -1 {
		t.Error("unexpectedly passed empty column header test for Row()")
	}

    // Test Row() - missing item.
	if r := tb.Row("header", nil); r != -1 {
		t.Error("unexpectedly passed missing item test for Row()")
	}
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

	tb, _ = New("1", "2", "3")
	if err := tb.AddRow("a", "b", "c"); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[a b c]")
	checkCount(t, tb, 3)

	if err := tb.AddRow("d", "e", "f"); err != nil {
		t.Error(err)
	}
	checkString(t, tb, "[a b c], [d e f]")
	checkCount(t, tb, 6)

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


// --- Helper Functions ---
func checkString(t *testing.T, tb *Table, want string) {
	s := tb.String()
	if s != want {
		t.Error("Table items are incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", s)
	}
}

func checkCount(t *testing.T, tb *Table, want int) {
	if tb.Count() != want {
		t.Error("Table count is incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", tb.Count())
	}
}
