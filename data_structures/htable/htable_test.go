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
	tb.RemoveRow(2)

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

