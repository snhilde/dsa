package htable

import (
	"testing"
)


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
}

func TestBadArgs(t *testing.T) {
	tb, _ := New("item")

	// Test New() - missing column header.
	if tbl, err := New(""); tbl != nil || err == nil {
		t.Error("unexpectedly passed missing column header test for New()")
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


func TestNew(t *testing.T) {
	n, err := New("a", "b", "c")
	if err != nil {
		t.Error(err)
	} else if n == nil {
		t.Error("Failed to create new table")
	}
}
