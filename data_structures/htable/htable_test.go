package htable_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/snhilde/dsa/data_structures/htable"
)

// --- Blanket Tests ---

func TestTBadPtr(t *testing.T) {
	var tb *htable.Table

	// Test Add().
	if err := tb.Add("a", "b", "c"); err == nil {
		t.Error("Unexpectedly passed bad pointer test for Add()")
	}

	// Test Insert().
	if err := tb.Insert(1, "a", "b", "c"); err == nil {
		t.Error("Unexpectedly passed bad pointer test for Insert()")
	}

	// Test AddRow().
	if err := tb.AddRow(htable.NewRow(1, 2, 3)); err == nil {
		t.Error("Unexpectedly passed bad pointer test for AddRow()")
	}

	// Test InsertRow().
	if err := tb.InsertRow(0, htable.NewRow(1, 2, 3)); err == nil {
		t.Error("Unexpectedly passed bad pointer test for InsertRow()")
	}

	// Test RemoveRow().
	if err := tb.RemoveRow(2); err == nil {
		t.Error("Unexpectedly passed bad pointer test for RemoveRow()")
	}

	// Test Clear().
	if err := tb.Clear(); err == nil {
		t.Error("Unexpectedly passed bad pointer test for Clear()")
	}

	// Test ColumnToIndex().
	if tb.ColumnToIndex("1") != -1 {
		t.Error("Unexpectedly passed bad pointer test for ColumnToIndex()")
	}

	// Test String().
	if tb.String() != "<nil>" {
		t.Error("Unexpectedly passed bad pointer test for String()")
	}

	// Test SetItem().
	if err := tb.SetItem("1", 0, 1); err == nil {
		t.Error("Unexpectedly passed bad pointer test for SetItem()")
	}

	// Test SetHeader().
	if err := tb.SetHeader("1", "10"); err == nil {
		t.Error("Unexpectedly passed bad pointer test for SetHeader()")
	}

	// Test Headers().
	if tb.Headers() != nil {
		t.Error("Unexpectedly passed bad pointer test for Headers()")
	}

	// Test Copy().
	if tb.Copy() != nil {
		t.Error("Unexpectedly passed bad pointer test for Copy()")
	}

	// Test Rows().
	if tb.Rows() != -1 {
		t.Error("Unexpectedly passed bad pointer test for Rows()")
	}

	// Test Enabled().
	if tb.Enabled() != -1 {
		t.Error("Unexpectedly passed bad pointer test for Enabled()")
	}

	// Test Disabled().
	if tb.Disabled() != -1 {
		t.Error("Unexpectedly passed bad pointer test for Disabled()")
	}

	// Test Columns().
	if tb.Columns() != -1 {
		t.Error("Unexpectedly passed bad pointer test for Columns()")
	}

	// Test Count().
	if tb.Count() != -1 {
		t.Error("Unexpectedly passed bad pointer test for Count()")
	}

	// Test Same().
	if tb.Same(nil) {
		t.Error("Unexpectedly passed bad pointer test for Same()")
	}

	// Test Row().
	if i, r := tb.Row("a", 5); i != -1 || r != nil {
		t.Error("Unexpectedly passed bad pointer test for Row()")
	}

	// Test RowByIndex().
	if tb.RowByIndex(0) != nil {
		t.Error("Unexpectedly passed bad pointer test for RowByIndex()")
	}

	// Test Item().
	if v := tb.Item("a", 5); v != nil {
		t.Error("Unexpectedly passed bad pointer test for Item()")
	}

	// Test Matches().
	if tb.Matches("a", 0, "item") {
		t.Error("Unexpectedly passed bad pointer test for Matches()")
	}

	// Test Toggle().
	if err := tb.Toggle(0, false); err == nil {
		t.Error("Unexpectedly passed bad pointer test for Toggle()")
	}

	// Test SortByColumn().
	if err := tb.SortByColumn("a", func(left, right interface{}) bool { return true }); err == nil {
		t.Error("Unexpectedly passed bad pointer test for SortByColumn()")
	}

	// Test SortByRow().
	if err := tb.SortByRow(func(left, right *htable.Row) bool { return true }); err == nil {
		t.Error("Unexpectedly passed bad pointer test for SortByRow()")
	}

	// Test CSV().
	if tb.CSV() != "" {
		t.Error("Unexpectedly passed bad pointer test for CSV()")
	}
}

func TestRBadPtr(t *testing.T) {
	var r *htable.Row

	// Test String().
	if r.String() != "<nil>" {
		t.Error("Unexpectedly passed bad pointer test for Row's String()")
	}

	// Test SetItem().
	if err := r.SetItem(0, "item"); err == nil {
		t.Error("Unexpectedly passed bad pointer test for Row's SetItem()")
	}

	// Test Count().
	if r.Count() != -1 {
		t.Error("Unexpectedly passed bad pointer test for Row's Count()")
	}

	// Test Item().
	if r.Item(0) != nil {
		t.Error("Unexpectedly passed bad pointer test for Row's Item()")
	}

	// Test Items().
	if r.Items() != nil {
		t.Error("Unexpectedly passed bad pointer test for Row's Items()")
	}

	// Test Enabled().
	if r.Enabled() {
		t.Error("Unexpectedly passed bad pointer test for Row's Enabled()")
	}

	// Test Matches().
	if r.Matches(0, 1) {
		t.Error("Unexpectedly passed bad pointer test for Row's Matches()")
	}

	// Test Copy().
	if r.Copy() != nil {
		t.Error("Unexpectedly passed bad pointer test for Row's Copy()")
	}
}

func TestTBadArgs(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")
	tb.Add(1, 2, 3)

	// Test New() - duplicate column headers.
	if tbl, err := htable.New("aaa", "bbb", "aaa", "ccc"); tbl != nil || err == nil {
		t.Error("Unexpectedly passed duplicate column header test for New()")
	}

	// Test New() - missing column header.
	if tbl, err := htable.New(""); tbl != nil || err == nil {
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

	// Test Insert() - too few columns.
	if err := tb.Insert(0, 1, 2); err == nil {
		t.Error("Unexpectedly passed too few columns test for Insert()")
	}

	// Test Insert() - too many columns.
	if err := tb.Insert(0, 1, 2, 3, 4); err == nil {
		t.Error("Unexpectedly passed too many columns test for Insert()")
	}

	// Test Insert() - wrong type in first position.
	if err := tb.Insert(0, "item1", 5, 6); err == nil {
		t.Error("Unexpectedly passed wrong type (first) test for Insert()")
	}

	// Test Insert() - wrong type in middle position.
	if err := tb.Insert(0, 4, "item5", 6); err == nil {
		t.Error("Unexpectedly passed wrong type (middle) test for Insert()")
	}

	// Test Insert() - wrong type in last position.
	if err := tb.Insert(0, 4, 5, "item6"); err == nil {
		t.Error("Unexpectedly passed wrong type (last) test for Insert()")
	}

	// Test AddRow() - too few columns.
	if err := tb.AddRow(htable.NewRow(1, 2)); err == nil {
		t.Error("Unexpectedly passed too few columns test for AddRow()")
	}

	// Test AddRow() - too many columns.
	if err := tb.AddRow(htable.NewRow(1, 2, 3, 4)); err == nil {
		t.Error("Unexpectedly passed too many columns test for AddRow()")
	}

	// Test AddRow() - wrong type in first position.
	if err := tb.AddRow(htable.NewRow("item1", 5, 6)); err == nil {
		t.Error("Unexpectedly passed wrong type (first) test for AddRow()")
	}

	// Test AddRow() - wrong type in middle position.
	if err := tb.AddRow(htable.NewRow(4, "item5", 6)); err == nil {
		t.Error("Unexpectedly passed wrong type (middle) test for AddRow()")
	}

	// Test AddRow() - wrong type in last position.
	if err := tb.AddRow(htable.NewRow(4, 5, "item6")); err == nil {
		t.Error("Unexpectedly passed wrong type (last) test for AddRow()")
	}

	// Test InsertRow() - negative index.
	if err := tb.InsertRow(-1, htable.NewRow(-1, 4, 5, 6)); err == nil {
		t.Error("Unexpectedly passed negative index test for InsertRow()")
	}

	// Test InsertRow() - out-of-bounds index.
	if err := tb.InsertRow(100, htable.NewRow(100, 4, 5, 6)); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test for InsertRow()")
	}

	// Test InsertRow() - too few columns.
	if err := tb.InsertRow(0, htable.NewRow(1, 2)); err == nil {
		t.Error("Unexpectedly passed too few columns test for InsertRow()")
	}

	// Test InsertRow() - too many columns.
	if err := tb.InsertRow(0, htable.NewRow(1, 2, 3, 4)); err == nil {
		t.Error("Unexpectedly passed too many columns test for InsertRow()")
	}

	// Test InsertRow() - wrong type in first position.
	if err := tb.InsertRow(0, htable.NewRow("item1", 5, 6)); err == nil {
		t.Error("Unexpectedly passed wrong type (first) test for InsertRow()")
	}

	// Test InsertRow() - wrong type in middle position.
	if err := tb.InsertRow(0, htable.NewRow(4, "item5", 6)); err == nil {
		t.Error("Unexpectedly passed wrong type (middle) test for InsertRow()")
	}

	// Test InsertRow() - wrong type in last position.
	if err := tb.InsertRow(0, htable.NewRow(4, 5, "item6")); err == nil {
		t.Error("Unexpectedly passed wrong type (last) test for InsertRow()")
	}

	// Test RemoveRow() - negative index.
	if err := tb.RemoveRow(-1); err == nil {
		t.Error("Unexpectedly passed negative index test for RemoveRow()")
	}

	// Test RemoveRow() - out-of-bounds index.
	if err := tb.RemoveRow(100); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test for RemoveRow()")
	}

	// Test ColumnToIndex() - empty column header.
	if tb.ColumnToIndex("") != -1 {
		t.Error("Unexpectedly passed empty column header test for ColumnToIndex()")
	}

	// Test ColumnToIndex() - invalid column header.
	if tb.ColumnToIndex("4") != -1 {
		t.Error("Unexpectedly passed invalid column header test for ColumnToIndex()")
	}

	// Test SetItem() - negative index.
	if err := tb.SetItem("1", -1, "value"); err == nil {
		t.Error("Unexpectedly passed negative index test for SetItem()")
	}

	// Test SetItem() - out-of-bounds index.
	if err := tb.SetItem("1", 100, "value"); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test for SetItem()")
	}

	// Test SetItem() - missing column header.
	if err := tb.SetItem("", 0, "value"); err == nil {
		t.Error("Unexpectedly passed missing column header test for SetItem()")
	}

	// Test SetItem() - invalid column header.
	if err := tb.SetItem("4", 0, "value"); err == nil {
		t.Error("Unexpectedly passed invalid column header test for SetItem()")
	}

	// Test SetHeader() - empty column header.
	if err := tb.SetHeader("", "30"); err == nil {
		t.Error("Unexpectedly passed empty column header test for SetHeader()")
	}

	// Test SetHeader() - invalid column header.
	if err := tb.SetHeader("4", "30"); err == nil {
		t.Error("Unexpectedly passed invalid column header test for SetHeader()")
	}

	// Test SetHeader() - empty name.
	if err := tb.SetHeader("3", ""); err == nil {
		t.Error("Unexpectedly passed empty name test for SetHeader()")
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

	// Test RowByIndex() - negative index.
	if tb.RowByIndex(-1) != nil {
		t.Error("Unexpectedly passed negative index test for RowByIndex()")
	}

	// Test RowByIndex() - out-of-bounds index.
	if tb.RowByIndex(100) != nil {
		t.Error("Unexpectedly passed out-of-bounds index test for RowByIndex()")
	}

	// Test Item() - negative index.
	if tb.Item("1", -1) != nil {
		t.Error("Unexpectedly passed negative index test for Item()")
	}

	// Test Item() - out-of-bounds index.
	if tb.Item("1", 100) != nil {
		t.Error("Unexpectedly passed out-of-bounds index test for Item()")
	}

	// Test Item() - empty column header.
	if tb.Item("", 0) != nil {
		t.Error("Unexpectedly passed empty column header test for Item()")
	}

	// Test Item() - invalid column header.
	if tb.Item("4", 0) != nil {
		t.Error("Unexpectedly passed invalid column header test for Item()")
	}

	// Test Same() - nil table.
	if tb.Same(nil) {
		t.Error("Unexpectedly passed nil table test for Same()")
	}

	// Test Matches() - negative index.
	if tb.Matches("1", -1, 1) {
		t.Error("Unexpectedly passed negative index test for Matches()")
	}

	// Test Matches() - out-of-bounds index.
	if tb.Matches("1", 100, 1) {
		t.Error("Unexpectedly passed out-of-bounds index test for Matches()")
	}

	// Test Matches() - empty column header.
	if tb.Matches("", 0, 1) {
		t.Error("Unexpectedly passed empty column header test for Matches()")
	}

	// Test Matches() - invalid column header.
	if tb.Matches("4", 0, 4) {
		t.Error("Unexpectedly passed invalid column header test for Matches()")
	}

	// Test Matches() - missing item.
	if tb.Matches("1", 0, nil) {
		t.Error("Unexpectedly passed missing item test for Matches()")
	}

	// Test Toggle() - negative index.
	if err := tb.Toggle(-1, true); err == nil {
		t.Error("Unexpectedly passed negative index test for Toggle()")
	}

	// Test Toggle() - out-of-bounds index.
	if err := tb.Toggle(100, true); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test for Toggle()")
	}

	// Test SortByColumn() - invalid column header.
	if err := tb.SortByColumn("4", func(left, right interface{}) bool { return true }); err == nil {
		t.Error("Unexpectedly passed invalid column header test for SortByColumn()")
	}

	// Test SortByColumn() - missing comparison function.
	if err := tb.SortByColumn("1", nil); err == nil {
		t.Error("Unexpectedly passed missing comparison function test for SortByColumn()")
	}

	// Test SortByRow() - missing comparison function.
	if err := tb.SortByRow(nil); err == nil {
		t.Error("Unexpectedly passed missing comparison function test for SortByRow()")
	}

	// Test Add() - adding table to itself.
	tb, _ = htable.New("1")
	tb2, _ := htable.New("2")
	if err := tb.Add(tb2); err != nil {
		t.Error(err)
	}
	if err := tb.Add(tb); err == nil {
		t.Error("Unexpectedly passed adding table to itself test for Add()")
	}

	// Test Insert() - adding table to itself.
	tb, _ = htable.New("1")
	tb2, _ = htable.New("2")
	if err := tb.Add(tb2); err != nil {
		t.Error(err)
	}
	if err := tb.Insert(0, tb); err == nil {
		t.Error("Unexpectedly passed adding table to itself test for Insert()")
	}
}

func TestRBadArgs(t *testing.T) {
	r := htable.NewRow(1, 2, 3)

	// Test SetItem() - negative index.
	if err := r.SetItem(-1, nil); err == nil {
		t.Error("Unexpectedly passed negative index test for SetItem()")
	}

	// Test SetItem() - out-of-bounds index.
	if err := r.SetItem(100, nil); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test for SetItem()")
	}

	// Test Item() - negative index.
	if r.Item(-1) != nil {
		t.Error("Unexpectedly passed negative index test for Item()")
	}

	// Test Item() - out-of-bounds index.
	if r.Item(100) != nil {
		t.Error("Unexpectedly passed out-of-bounds index test for Item()")
	}

	// Test Matches() - negative index.
	if r.Matches(-1, 1) {
		t.Error("Unexpectedly passed negative index test for Matches()")
	}

	// Test Matches() - out-of-bounds index.
	if r.Matches(100, 1) {
		t.Error("Unexpectedly passed out-of-bounds index test for Matches()")
	}
}

func TestBadTypes(t *testing.T) {
	type s1 struct {
		s string
		n int
	}
	type s2 struct {
		i int
		v float32
	}

	types := []interface{}{
		int(1),
		int8(1),
		int16(1),
		int32(1),
		int64(1),
		uint(1),
		uint8(1),
		uint16(1),
		uint32(1),
		uint64(1),
		float32(1.1),
		float64(1.1),
		complex64(1 + 10i),
		complex128(4 + 40i),
		s1{"string", 1},
		s2{2, 3.14},
		TestTNew,
		checkTString,
		checkTCount,
	}

	for i, headerType := range types {
		tb, err := htable.New("1")
		if err != nil {
			t.Error(err)
		}

		// Set the type for this column.
		if err := tb.Add(headerType); err != nil {
			t.Error(err)
		}

		// Make sure that we can't put any other type in this column.
		for j, newType := range types {
			if i == j {
				continue
			}

			if err := tb.Add(newType); err == nil {
				t.Error("Unexpectedly passed adding a row with the wrong type")
				t.Log("\tHeader Type:", reflect.TypeOf(headerType))
				t.Log("\tNew Type:", reflect.TypeOf(newType))
			}
		}
	}
}

// --- Function Tests ---

func TestTNew(t *testing.T) {
	n, err := htable.New("a", "b", "c")
	if err != nil {
		t.Error(err)
	} else if n == nil {
		t.Error("Failed to create new table")
	}
}

func TestRNewRow(t *testing.T) {
	r := htable.NewRow(1, 2, 3)
	checkRString(t, r, "{1, 2, 3}")
	checkRCount(t, r, 3)

	type s struct {
		s1 string
		s2 string
	}
	r = htable.NewRow(1, 2, s{s1: "a", s2: "b"})
	checkRString(t, r, "{1, 2, {a b}}")
	checkRCount(t, r, 3)

	tmp := htable.NewRow("other", "row")
	r = htable.NewRow("this row", "here", 5, tmp)
	checkRString(t, r, "{this row, here, 5, {other, row}}")
	checkRCount(t, r, 4)
}

// --- Table's Method Tests ---

func TestTAdd(t *testing.T) {
	// Testing rows of integers.
	tb, _ := htable.New("1", "2", "3")
	if err := tb.Add(1, 2, 3); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}")
	checkTCSV(t, tb, "1,2,3", "1,2,3")
	checkTCount(t, tb, 3)

	if err := tb.Add(4, 5, 6); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6")
	checkTCount(t, tb, 6)

	// Test rows of characters.
	tb, _ = htable.New("1", "2", "3")
	if err := tb.Add('a', 'b', 'c'); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 97, 2: 98, 3: 99}")
	checkTCSV(t, tb, "1,2,3", "97,98,99")
	checkTCount(t, tb, 3)

	if err := tb.Add('d', 'e', 'f'); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 97, 2: 98, 3: 99}, {1: 100, 2: 101, 3: 102}")
	checkTCSV(t, tb, "1,2,3", "97,98,99\r\n100,101,102")
	checkTCount(t, tb, 6)

	// Test rows of int slices.
	tb, _ = htable.New("1", "2", "3")
	if err := tb.Add([]int{10, 20}, []int{30, 40}, []int{50, 60}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: [10 20], 2: [30 40], 3: [50 60]}")
	checkTCSV(t, tb, "1,2,3", "[10 20],[30 40],[50 60]")
	checkTCount(t, tb, 3)

	if err := tb.Add([]int{100, 200}, []int{300, 400}, []int{500, 600}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: [10 20], 2: [30 40], 3: [50 60]}, {1: [100 200], 2: [300 400], 3: [500 600]}")
	checkTCSV(t, tb, "1,2,3", "[10 20],[30 40],[50 60]\r\n[100 200],[300 400],[500 600]")
	checkTCount(t, tb, 6)

	// Test mixed-value rows.
	tb, _ = htable.New("1", "2", "3")
	if err := tb.Add(1.1, "b", []byte{0x03}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1.1, 2: b, 3: [3]}")
	checkTCSV(t, tb, "1,2,3", "1.1,b,[3]")
	checkTCount(t, tb, 3)

	if err := tb.Add(4.4, "e", []byte{0x06}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1.1, 2: b, 3: [3]}, {1: 4.4, 2: e, 3: [6]}")
	checkTCSV(t, tb, "1,2,3", "1.1,b,[3]\r\n4.4,e,[6]")
	checkTCount(t, tb, 6)

	// Test rows of strings with spaces.
	tb, _ = htable.New("1", "2", "3")
	if err := tb.Add("before ", "in between", " after"); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: before , 2: in between, 3:  after}")
	checkTCSV(t, tb, "1,2,3", "before ,in between, after")
	checkTCount(t, tb, 3)

	if err := tb.Add("AAA", "   ", ""); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: before , 2: in between, 3:  after}, {1: AAA, 2:    , 3: }")
	checkTCSV(t, tb, "1,2,3", "before ,in between, after\r\nAAA,   ,")
	checkTCount(t, tb, 6)

	// Test rows masquerading as interfaces.
	tb, _ = htable.New("1", "2", "3")
	if err := tb.Add(interface{}(1), interface{}(2), interface{}(3)); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}")
	checkTCSV(t, tb, "1,2,3", "1,2,3")
	checkTCount(t, tb, 3)

	if err := tb.Add(4, 5, 6); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6")
	checkTCount(t, tb, 6)

	// Test rows of functions.
	tb, _ = htable.New("1", "2", "3")
	if err := tb.Add(TestTNew, checkTString, TestTInsert); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, fmt.Sprintf("{1: %p, 2: %p, 3: %p}", TestTNew, checkTString, TestTInsert))
	checkTCSV(t, tb, "1,2,3", fmt.Sprintf("%p,%p,%p", TestTNew, checkTString, TestTInsert))
	checkTCount(t, tb, 3)

	if err := tb.Add(TestTAdd, checkTString, TestTRows); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, fmt.Sprintf("{1: %p, 2: %p, 3: %p}, {1: %p, 2: %p, 3: %p}", TestTNew, checkTString, TestTInsert, TestTAdd, checkTString, TestTRows))
	checkTCSV(t, tb, "1,2,3", fmt.Sprintf("%p,%p,%p\r\n%p,%p,%p", TestTNew, checkTString, TestTInsert, TestTAdd, checkTString, TestTRows))
	checkTCount(t, tb, 6)
}

func TestTInsert(t *testing.T) {
	// Test inserting a row at the beginning.
	tb, _ := htable.New("1", "2", "3")
	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6")
	checkTCount(t, tb, 6)
	if err := tb.Insert(0, -1, -2, -3); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: -1, 2: -2, 3: -3}, {1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "-1,-2,-3\r\n1,2,3\r\n4,5,6")
	checkTCount(t, tb, 9)

	// Test inserting a row in the middle.
	tb, _ = htable.New("1", "2", "3")
	tb.Add("a", "b", "c")
	tb.Add("d", "e", "f")
	checkTString(t, tb, "{1: a, 2: b, 3: c}, {1: d, 2: e, 3: f}")
	checkTCSV(t, tb, "1,2,3", "a,b,c\r\nd,e,f")
	checkTCount(t, tb, 6)
	if err := tb.Insert(1, "x", "y", "z"); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: a, 2: b, 3: c}, {1: x, 2: y, 3: z}, {1: d, 2: e, 3: f}")
	checkTCSV(t, tb, "1,2,3", "a,b,c\r\nx,y,z\r\nd,e,f")
	checkTCount(t, tb, 9)

	// Test inserting a row at the end.
	tb, _ = htable.New("1", "2", "3")
	tb.Add([]int{10, 20}, []int{30, 40}, []int{50, 60})
	tb.Add([]int{100, 200}, []int{300, 400}, []int{500, 600})
	checkTString(t, tb, "{1: [10 20], 2: [30 40], 3: [50 60]}, {1: [100 200], 2: [300 400], 3: [500 600]}")
	checkTCSV(t, tb, "1,2,3", "[10 20],[30 40],[50 60]\r\n[100 200],[300 400],[500 600]")
	checkTCount(t, tb, 6)
	if err := tb.Insert(2, []int{-1, -2, -3}, []int{-4, -5, -6}, []int{-7, -8, -9}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: [10 20], 2: [30 40], 3: [50 60]}, {1: [100 200], 2: [300 400], 3: [500 600]}, {1: [-1 -2 -3], 2: [-4 -5 -6], 3: [-7 -8 -9]}")
	checkTCSV(t, tb, "1,2,3", "[10 20],[30 40],[50 60]\r\n[100 200],[300 400],[500 600]\r\n[-1 -2 -3],[-4 -5 -6],[-7 -8 -9]")
	checkTCount(t, tb, 9)

	// Test inserting a row beyond the table's current boundaries.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(1.1, "b", []byte{0x03})
	tb.Add(4.4, "e", []byte{0x06})
	checkTString(t, tb, "{1: 1.1, 2: b, 3: [3]}, {1: 4.4, 2: e, 3: [6]}")
	checkTCSV(t, tb, "1,2,3", "1.1,b,[3]\r\n4.4,e,[6]")
	checkTCount(t, tb, 6)
	if err := tb.Insert(3, 7.7, "h", []byte{0x09}); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test")
	}
	checkTString(t, tb, "{1: 1.1, 2: b, 3: [3]}, {1: 4.4, 2: e, 3: [6]}")
	checkTCSV(t, tb, "1,2,3", "1.1,b,[3]\r\n4.4,e,[6]")
	checkTCount(t, tb, 6)

	// Test only inserting instead of appending.
	tb, _ = htable.New("1", "2", "3")
	if err := tb.Insert(0, "before ", "in between", " after"); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: before , 2: in between, 3:  after}")
	checkTCSV(t, tb, "1,2,3", "before ,in between, after")
	checkTCount(t, tb, 3)
	if err := tb.Insert(1, "AAA", "   ", ""); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: before , 2: in between, 3:  after}, {1: AAA, 2:    , 3: }")
	checkTCSV(t, tb, "1,2,3", "before ,in between, after\r\nAAA,   ,")
	checkTCount(t, tb, 6)
	if err := tb.Insert(0, "first", "second", "third"); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: first, 2: second, 3: third}, {1: before , 2: in between, 3:  after}, {1: AAA, 2:    , 3: }")
	checkTCSV(t, tb, "1,2,3", "first,second,third\r\nbefore ,in between, after\r\nAAA,   ,")
	checkTCount(t, tb, 9)
}

func TestTAddRow(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)

	// Test adding a row with the correct number of columns.
	r := htable.NewRow(1, 2, 3)
	if err := tb.AddRow(r); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}")
	checkTCSV(t, tb, "1,2,3", "1,2,3")
	checkTCount(t, tb, 3)

	// Test adding a row with too few columns.
	r = htable.NewRow(1, 2)
	if err := tb.AddRow(r); err == nil {
		t.Error("Unexpectedly passed adding a row with too few columns")
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}")
	checkTCSV(t, tb, "1,2,3", "1,2,3")
	checkTCount(t, tb, 3)

	// Test adding a row with too many columns.
	r = htable.NewRow(1, 2, 3, 4)
	if err := tb.AddRow(r); err == nil {
		t.Error("Unexpectedly passed adding a row with too many columns")
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}")
	checkTCSV(t, tb, "1,2,3", "1,2,3")
	checkTCount(t, tb, 3)

	// Test adding a row with the right types.
	r = htable.NewRow(4, 5, 6)
	if err := tb.AddRow(r); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6")
	checkTCount(t, tb, 6)

	// Test adding a row with the first value being the wrong type.
	r = htable.NewRow("7", 8, 9)
	if err := tb.AddRow(r); err == nil {
		t.Error("Unexpectedly passed adding a row with the first value being the wrong type")
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6")
	checkTCount(t, tb, 6)

	// Test adding a row with the middle value being the wrong type.
	r = htable.NewRow(7, "8", 9)
	if err := tb.AddRow(r); err == nil {
		t.Error("Unexpectedly passed adding a row with the middle value being the wrong type")
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6")
	checkTCount(t, tb, 6)

	// Test adding a row with the last value being the wrong type.
	r = htable.NewRow(7, 8, "9")
	if err := tb.AddRow(r); err == nil {
		t.Error("Unexpectedly passed adding a row with the last value being the wrong type")
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6")
	checkTCount(t, tb, 6)

	// Test adding a row with all values being the wrong type.
	r = htable.NewRow("7", "8", "9")
	if err := tb.AddRow(r); err == nil {
		t.Error("Unexpectedly passed adding a row with all values being the wrong type")
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6")
	checkTCount(t, tb, 6)
}

func TestTInsertRow(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)

	// Test inserting a row with the correct number of columns.
	r := htable.NewRow(1, 2, 3)
	if err := tb.InsertRow(0, r); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}")
	checkTCSV(t, tb, "1,2,3", "1,2,3")
	checkTCount(t, tb, 3)

	// Test adding a row with too few columns.
	r = htable.NewRow(1, 2)
	if err := tb.InsertRow(0, r); err == nil {
		t.Error("Unexpectedly passed adding a row with too few columns")
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}")
	checkTCSV(t, tb, "1,2,3", "1,2,3")
	checkTCount(t, tb, 3)

	// Test adding a row with too many columns.
	r = htable.NewRow(1, 2, 3, 4)
	if err := tb.InsertRow(0, r); err == nil {
		t.Error("Unexpectedly passed adding a row with too many columns")
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}")
	checkTCSV(t, tb, "1,2,3", "1,2,3")
	checkTCount(t, tb, 3)

	// Test inserting a row before the current first row.
	r = htable.NewRow(-1, -2, -3)
	if err := tb.InsertRow(0, r); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: -1, 2: -2, 3: -3}, {1: 1, 2: 2, 3: 3}")
	checkTCSV(t, tb, "1,2,3", "-1,-2,-3\r\n1,2,3")
	checkTCount(t, tb, 6)

	// Test inserting a row after the current first row.
	r = htable.NewRow(0, 0, 0)
	if err := tb.InsertRow(1, r); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: -1, 2: -2, 3: -3}, {1: 0, 2: 0, 3: 0}, {1: 1, 2: 2, 3: 3}")
	checkTCSV(t, tb, "1,2,3", "-1,-2,-3\r\n0,0,0\r\n1,2,3")
	checkTCount(t, tb, 9)

	// Test inserting a row at the end.
	r = htable.NewRow(4, 5, 6)
	if err := tb.InsertRow(tb.Rows(), r); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: -1, 2: -2, 3: -3}, {1: 0, 2: 0, 3: 0}, {1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "-1,-2,-3\r\n0,0,0\r\n1,2,3\r\n4,5,6")
	checkTCount(t, tb, 12)

	// Test inserting a row with the first value being the wrong type.
	r = htable.NewRow("7", 8, 9)
	if err := tb.InsertRow(1, r); err == nil {
		t.Error("Unexpectedly passed inserting a row with the first value being the wrong type")
	}
	checkTString(t, tb, "{1: -1, 2: -2, 3: -3}, {1: 0, 2: 0, 3: 0}, {1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "-1,-2,-3\r\n0,0,0\r\n1,2,3\r\n4,5,6")
	checkTCount(t, tb, 12)

	// Test inserting a row with the middle value being the wrong type.
	r = htable.NewRow(7, "8", 9)
	if err := tb.InsertRow(1, r); err == nil {
		t.Error("Unexpectedly passed inserting a row with the middle value being the wrong type")
	}
	checkTString(t, tb, "{1: -1, 2: -2, 3: -3}, {1: 0, 2: 0, 3: 0}, {1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "-1,-2,-3\r\n0,0,0\r\n1,2,3\r\n4,5,6")
	checkTCount(t, tb, 12)

	// Test inserting a row with the last value being the wrong type.
	r = htable.NewRow(7, 8, "9")
	if err := tb.InsertRow(1, r); err == nil {
		t.Error("Unexpectedly passed inserting a row with the last value being the wrong type")
	}
	checkTString(t, tb, "{1: -1, 2: -2, 3: -3}, {1: 0, 2: 0, 3: 0}, {1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "-1,-2,-3\r\n0,0,0\r\n1,2,3\r\n4,5,6")
	checkTCount(t, tb, 12)

	// Test inserting a row with all values being the wrong type.
	r = htable.NewRow("7", "8", "9")
	if err := tb.InsertRow(1, r); err == nil {
		t.Error("Unexpectedly passed inserting a row with all values being the wrong type")
	}
	checkTString(t, tb, "{1: -1, 2: -2, 3: -3}, {1: 0, 2: 0, 3: 0}, {1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "-1,-2,-3\r\n0,0,0\r\n1,2,3\r\n4,5,6")
	checkTCount(t, tb, 12)
}

func TestTRemoveRow(t *testing.T) {
	// Test removing a row at the beginning.
	tb, _ := htable.New("1", "2", "3")
	tb.Add(-1, -2, -3)
	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)
	checkTString(t, tb, "{1: -1, 2: -2, 3: -3}, {1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "-1,-2,-3\r\n1,2,3\r\n4,5,6")
	checkTCount(t, tb, 9)
	if err := tb.RemoveRow(0); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6")
	checkTCount(t, tb, 6)

	// Test removing a row in the middle.
	tb, _ = htable.New("1", "2", "3")
	tb.Add("a", "b", "c")
	tb.Add("x", "y", "z")
	tb.Add("d", "e", "f")
	checkTString(t, tb, "{1: a, 2: b, 3: c}, {1: x, 2: y, 3: z}, {1: d, 2: e, 3: f}")
	checkTCSV(t, tb, "1,2,3", "a,b,c\r\nx,y,z\r\nd,e,f")
	checkTCount(t, tb, 9)
	if err := tb.RemoveRow(1); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: a, 2: b, 3: c}, {1: d, 2: e, 3: f}")
	checkTCSV(t, tb, "1,2,3", "a,b,c\r\nd,e,f")
	checkTCount(t, tb, 6)

	// Test removing a row at the end.
	tb, _ = htable.New("1", "2", "3")
	tb.Add([]int{10, 20}, []int{30, 40}, []int{50, 60})
	tb.Add([]int{100, 200}, []int{300, 400}, []int{500, 600})
	tb.Add([]int{-1, -2, -3}, []int{-4, -5, -6}, []int{-7, -8, -9})
	checkTString(t, tb, "{1: [10 20], 2: [30 40], 3: [50 60]}, {1: [100 200], 2: [300 400], 3: [500 600]}, {1: [-1 -2 -3], 2: [-4 -5 -6], 3: [-7 -8 -9]}")
	checkTCSV(t, tb, "1,2,3", "[10 20],[30 40],[50 60]\r\n[100 200],[300 400],[500 600]\r\n[-1 -2 -3],[-4 -5 -6],[-7 -8 -9]")
	checkTCount(t, tb, 9)
	if err := tb.RemoveRow(2); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: [10 20], 2: [30 40], 3: [50 60]}, {1: [100 200], 2: [300 400], 3: [500 600]}")
	checkTCSV(t, tb, "1,2,3", "[10 20],[30 40],[50 60]\r\n[100 200],[300 400],[500 600]")
	checkTCount(t, tb, 6)

	// Test removing a row beyond the table's current boundaries.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(1.1, "b", []byte{0x03})
	tb.Add(4.4, "e", []byte{0x06})
	checkTString(t, tb, "{1: 1.1, 2: b, 3: [3]}, {1: 4.4, 2: e, 3: [6]}")
	checkTCSV(t, tb, "1,2,3", "1.1,b,[3]\r\n4.4,e,[6]")
	checkTCount(t, tb, 6)
	if err := tb.RemoveRow(3); err == nil {
		t.Error("Unexpectedly passed out-of-bounds index test")
	}
	checkTString(t, tb, "{1: 1.1, 2: b, 3: [3]}, {1: 4.4, 2: e, 3: [6]}")
	checkTCSV(t, tb, "1,2,3", "1.1,b,[3]\r\n4.4,e,[6]")
	checkTCount(t, tb, 6)

	// Test removing when there aren't any rows in the table.
	tb, _ = htable.New("1", "2", "3")
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)
	if err := tb.RemoveRow(0); err == nil {
		t.Error("Unexpectedly passed removing from empty table test")
	}
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)
}

func TestTClear(t *testing.T) {
	// Try clearing an empty table.
	tb, _ := htable.New("1", "2", "3")
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)

	if err := tb.Clear(); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)

	// Add some rows and clear again.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)
	tb.Add(7, 8, 9)
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}, {1: 7, 2: 8, 3: 9}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6\r\n7,8,9")
	checkTCount(t, tb, 9)

	if err := tb.Clear(); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)

	// Add some rows, delete one row, and then clear.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)
	tb.Add(7, 8, 9)
	tb.RemoveRow(1)
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 7, 2: 8, 3: 9}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n7,8,9")
	checkTCount(t, tb, 6)

	if err := tb.Clear(); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)

	// Add some rows, delete all the rows, and then clear.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)
	tb.Add(7, 8, 9)
	tb.RemoveRow(0)
	tb.RemoveRow(0)
	tb.RemoveRow(0)
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)

	if err := tb.Clear(); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)

	// Add some rows, disable a couple of rows, and clear the table.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)
	tb.Add(7, 8, 9)
	tb.Toggle(1, false)
	tb.Toggle(2, false)
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}")
	checkTCSV(t, tb, "1,2,3", "1,2,3")
	checkTCount(t, tb, 9)

	if err := tb.Clear(); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)
}

func TestTColumnToIndex(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")

	// Make sure each column lines up with the correct index.
	for i, v := range []string{"1", "2", "3"} {
		if n := tb.ColumnToIndex(v); n != i {
			t.Error("Column does not align to expected index")
			t.Log("\tExpected:", i)
			t.Log("\tReceived:", n)
		}
	}
}

func TestTSetItem(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)

	// Make sure you can't set any values with an empty table.
	if err := tb.SetItem("1", 0, 5); err == nil {
		t.Error("Unexpectedly passed setting a value with an empty table")
	}

	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)
	tb.Add(7, 8, 9)
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}, {1: 7, 2: 8, 3: 9}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6\r\n7,8,9")
	checkTCount(t, tb, 9)

	// Change each value for all three rows.
	for row := 0; row < 3; row++ {
		for i, v := range []string{"1", "2", "3"} {
			n := ((row * 3) + i + 1) * 10
			if err := tb.SetItem(v, row, n); err != nil {
				t.Error(err)
			}
		}
	}
	checkTString(t, tb, "{1: 10, 2: 20, 3: 30}, {1: 40, 2: 50, 3: 60}, {1: 70, 2: 80, 3: 90}")
	checkTCSV(t, tb, "1,2,3", "10,20,30\r\n40,50,60\r\n70,80,90")
	checkTCount(t, tb, 9)

	// Make sure you can't change an item's type.
	if err := tb.SetItem("1", 0, "a"); err == nil {
		t.Error("Unexpectedly passed changing an item's type")
	}
}

func TestTSetHeader(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")

	was := []string{"1", "2", "3"}
	to := []string{"4", "5", "6"}
	for i, v := range was {
		// First, make sure the current header is correct.
		h := tb.Headers()
		if v != h[i] {
			t.Error("Received incorrect header")
			t.Log("\tExpected:", v)
			t.Log("\tReceived:", h[i])
		}

		// Now, change the header.
		if err := tb.SetHeader(v, to[i]); err != nil {
			t.Error(err)
		}

		// Make sure the old value is gone and the new one is in place.
		h = tb.Headers()
		if to[i] != h[i] {
			t.Error("Did not change header")
			t.Log("\tExpected:", to[i])
			t.Log("\tReceived:", h[i])
		}
	}
}

func TestTHeaders(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")

	exp := []string{"1", "2", "3"}
	h := tb.Headers()
	for i, v := range h {
		if v != exp[i] {
			t.Error("Received incorrect header")
			t.Log("\tExpected:", exp[i])
			t.Log("\tReceived:", v)
		}
	}

	// Make sure you returned slice of headers doesn't reference the same underlying array as the table's headers.
	h[0] = "4"
	hh := tb.Headers()
	if hh[0] == "4" {
		t.Error("Changed table's column header")
	}
}

func TestTCopy(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")
	tb.Add(3, 10, 21)
	tb.Add(2, 11, 23)
	tb.Add(1, 12, 22)
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tb, 9)

	// Make a copy of the table and check that everything is the same.
	tCopy := tb.Copy()
	checkTString(t, tCopy, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tCopy, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tCopy, 9)

	// Make sure that a change in one table does not affect the other table.
	tb.Add(0, 0, 0)
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}, {1: 0, 2: 0, 3: 0}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22\r\n0,0,0")
	checkTCount(t, tb, 12)
	checkTString(t, tCopy, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tCopy, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tCopy, 9)

	tCopy.RemoveRow(0)
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}, {1: 0, 2: 0, 3: 0}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22\r\n0,0,0")
	checkTCount(t, tb, 12)
	checkTString(t, tCopy, "{1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tCopy, "1,2,3", "2,11,23\r\n1,12,22")
	checkTCount(t, tCopy, 6)

	tCopy.SetHeader("1", "9")
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}, {1: 0, 2: 0, 3: 0}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22\r\n0,0,0")
	checkTCount(t, tb, 12)
	checkTString(t, tCopy, "{9: 2, 2: 11, 3: 23}, {9: 1, 2: 12, 3: 22}")
	checkTCSV(t, tCopy, "9,2,3", "2,11,23\r\n1,12,22")
	checkTCount(t, tCopy, 6)
}

func TestTRows(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")
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

func TestTEnabled(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")
	if n := tb.Enabled(); n != 0 {
		t.Error("Enabled count is incorrect")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", n)
	}

	tb.Add(1, 2, 3)
	if n := tb.Enabled(); n != 1 {
		t.Error("Enabled count is incorrect")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", n)
	}

	tb.Toggle(0, false)
	if n := tb.Enabled(); n != 0 {
		t.Error("Enabled count is incorrect")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", n)
	}

	tb.Toggle(0, true)
	if n := tb.Enabled(); n != 1 {
		t.Error("Enabled count is incorrect")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", n)
	}

	tb.Add(4, 5, 6)
	if n := tb.Enabled(); n != 2 {
		t.Error("Enabled count is incorrect")
		t.Log("\tExpected:", 2)
		t.Log("\tReceived:", n)
	}

	tb.RemoveRow(0)
	if n := tb.Enabled(); n != 1 {
		t.Error("Enabled count is incorrect")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", n)
	}

	tb.Insert(1, 7, 8, 9)
	if n := tb.Enabled(); n != 2 {
		t.Error("Enabled count is incorrect")
		t.Log("\tExpected:", 2)
		t.Log("\tReceived:", n)
	}
}

func TestTDisabled(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")
	if n := tb.Disabled(); n != 0 {
		t.Error("Disabled count is incorrect")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", n)
	}

	tb.Add(1, 2, 3)
	if n := tb.Disabled(); n != 0 {
		t.Error("Disabled count is incorrect")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", n)
	}

	tb.Toggle(0, false)
	if n := tb.Disabled(); n != 1 {
		t.Error("Disabled count is incorrect")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", n)
	}

	tb.Toggle(0, true)
	if n := tb.Disabled(); n != 0 {
		t.Error("Disabled count is incorrect")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", n)
	}

	tb.Add(4, 5, 6)
	if n := tb.Disabled(); n != 0 {
		t.Error("Disabled count is incorrect")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", n)
	}

	tb.Toggle(0, false)
	if n := tb.Disabled(); n != 1 {
		t.Error("Disabled count is incorrect")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", n)
	}

	tb.RemoveRow(0)
	if n := tb.Disabled(); n != 0 {
		t.Error("Disabled count is incorrect")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", n)
	}

	tb.Insert(1, 7, 8, 9)
	if n := tb.Disabled(); n != 0 {
		t.Error("Disabled count is incorrect")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", n)
	}
}

func TestTColumns(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")
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

	tb, _ = htable.New("1")
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
	tb, _ := htable.New("1", "2", "3")
	checkTCount(t, tb, 0)

	tb.Add(1, 2, 3)
	checkTCount(t, tb, 3)

	tb.Add(4, 5, 6)
	checkTCount(t, tb, 6)

	tb.RemoveRow(0)
	checkTCount(t, tb, 3)

	tb.Insert(1, 7, 8, 9)
	checkTCount(t, tb, 6)

	tb, _ = htable.New("1")
	checkTCount(t, tb, 0)

	tb.Add(1)
	checkTCount(t, tb, 1)

	tb.RemoveRow(0)
	checkTCount(t, tb, 0)
}

func TestTSame(t *testing.T) {
	// Test that the same tables are the same.
	tb, _ := htable.New("1", "2", "3")
	checkTCount(t, tb, 0)

	tb.Add(1, 2, 3)
	checkTCount(t, tb, 3)

	tb2 := tb
	if !tb.Same(tb2) {
		t.Error("tables should be the same")
	}

	// Test that identical but not clone tables are not the same.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(1, 2, 3)
	checkTCount(t, tb, 3)

	tb2, _ = htable.New("1", "2", "3")
	tb2.Add(1, 2, 3)
	checkTCount(t, tb2, 3)

	if tb.Same(tb2) {
		t.Error("identical tables are not the same")
	}

	// Test that non-identical tables are not the same.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(1, 2, 3)
	checkTCount(t, tb, 3)

	tb2, _ = htable.New("a")
	tb2.Add("line")
	checkTCount(t, tb2, 1)

	if tb.Same(tb2) {
		t.Error("non-identical tables are not the same")
	}
}

func TestTRowItem(t *testing.T) {
	// Set up a new table.
	tb, _ := htable.New("1", "2", "3")
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
			if v := tb.Item(col, row); v != exp {
				t.Error("Item value incorrect")
				t.Log("\tExpected:", exp)
				t.Log("\tReceived:", v)
			}
		}
	}

	// Set up a new table.
	tb, _ = htable.New("name", "left-handed", "age")
	tb.Add("Swari", true, 30)
	tb.Add("Kathy", false, 40)
	tb.Add("Joe", false, 189)

	// Find the first person who is left-handed.
	if i, _ := tb.Row("left-handed", true); i != 0 {
		t.Error("Matched incorrect row")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", i)
	} else if v := tb.Item("name", i); v != "Swari" {
		t.Error("Item value incorrect")
		t.Log("\tExpected:", "Swari")
		t.Log("\tReceived:", v)
	}

	// Find the first person who is right-handed.
	if i, _ := tb.Row("left-handed", false); i != 1 {
		t.Error("Matched incorrect row")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", i)
	} else if v := tb.Item("name", i); v != "Kathy" {
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
	} else if v := tb.Item("height", i); v != nil {
		t.Error("Unexpectedly found item")
		t.Log("\tExpected: nil")
		t.Log("\tReceived:", v)
	}

	// Store and run functions.
	hi := func() string { return "hello" }
	reply := func() string { return "how are you" }

	tb, _ = htable.New("name", "func")
	tb.Add("hi", hi)
	tb.Add("reply", reply)

	// Find and run the first function.
	if i, _ := tb.Row("name", "hi"); i != 0 {
		t.Error("Matched incorrect row")
		t.Log("\tExpected:", 0)
		t.Log("\tReceived:", i)
	} else {
		f := tb.Item("func", i)
		if ff, ok := f.(func() string); !ok {
			t.Error("Returned function is not of type func() string")
		} else if v := ff(); v != "hello" {
			t.Error("Function result incorrect")
			t.Log("\tExpected:", "hello")
			t.Log("\tReceived:", v)
		}
	}

	// Find and run the second function.
	if i, _ := tb.Row("name", "reply"); i != 1 {
		t.Error("Matched incorrect row")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", i)
	} else {
		f := tb.Item("func", i)
		if ff, ok := f.(func() string); !ok {
			t.Error("Returned function is not of type func() string")
		} else if v := ff(); v != "how are you" {
			t.Error("Function result incorrect")
			t.Log("\tExpected:", "how are you")
			t.Log("\tReceived:", v)
		}
	}
}

func TestTRowByIndex(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")
	tb.Add(1, "a", 1.1)
	tb.Add(2, "b", 2.2)
	tb.Add(3, "c", 3.3)
	checkTString(t, tb, "{1: 1, 2: a, 3: 1.1}, {1: 2, 2: b, 3: 2.2}, {1: 3, 2: c, 3: 3.3}")
	checkTCSV(t, tb, "1,2,3", "1,a,1.1\r\n2,b,2.2\r\n3,c,3.3")
	checkTCount(t, tb, 9)

	// Grab some rows and make sure they're correct.
	row := tb.RowByIndex(0)
	checkRString(t, row, "{1, a, 1.1}")
	checkRCount(t, row, 3)

	row = tb.RowByIndex(1)
	checkRString(t, row, "{2, b, 2.2}")
	checkRCount(t, row, 3)

	row = tb.RowByIndex(2)
	checkRString(t, row, "{3, c, 3.3}")
	checkRCount(t, row, 3)

	// Make sure that modifying the returned row also modifies the row in the table.
	row = tb.RowByIndex(0)
	row.SetItem(0, 4)
	checkRString(t, row, "{4, a, 1.1}")
	checkRCount(t, row, 3)
	checkTString(t, tb, "{1: 4, 2: a, 3: 1.1}, {1: 2, 2: b, 3: 2.2}, {1: 3, 2: c, 3: 3.3}")
	checkTCSV(t, tb, "1,2,3", "4,a,1.1\r\n2,b,2.2\r\n3,c,3.3")
	checkTCount(t, tb, 9)
}

func TestTMatches(t *testing.T) {
	// Set up a new table.
	tb, _ := htable.New("1", "2", "3")
	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)
	tb.Add(7, 8, 9)

	// Make sure we have a match at each position.
	cols := []string{"1", "2", "3"}
	for row := 0; row < 3; row++ {
		for i, col := range cols {
			v := (row * 3) + i + 1
			if !tb.Matches(col, row, v) {
				t.Error("Did not match")
			}
		}
	}

	// Set up a new table.
	tb, _ = htable.New("name", "left-handed", "age")
	tb.Add("Swari", true, 30)
	tb.Add("Kathy", false, 40)
	tb.Add("Joe", false, 189)

	// Make sure Swari is left-handed.
	if !tb.Matches("left-handed", 0, true) {
		t.Error("Did not match")
	}

	// Make sure Joe is 189.
	if !tb.Matches("age", 2, 189) {
		t.Error("Did not match")
	}

	// Look for a name that doesn't exist.
	if tb.Matches("name", 0, "April") {
		t.Error("Unexpectedly matched")
	}
}

func TestTToggle(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")
	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)
	tb.Add(7, 8, 9)
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}, {1: 7, 2: 8, 3: 9}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6\r\n7,8,9")
	checkTCount(t, tb, 9)

	// Disable first row.
	if err := tb.Toggle(0, false); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 4, 2: 5, 3: 6}, {1: 7, 2: 8, 3: 9}")
	checkTCSV(t, tb, "1,2,3", "4,5,6\r\n7,8,9")
	checkTCount(t, tb, 9)

	// Enable first row.
	if err := tb.Toggle(0, true); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}, {1: 7, 2: 8, 3: 9}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6\r\n7,8,9")
	checkTCount(t, tb, 9)

	// Disable middle row.
	if err := tb.Toggle(1, false); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 7, 2: 8, 3: 9}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n7,8,9")
	checkTCount(t, tb, 9)

	// Enable middle row.
	if err := tb.Toggle(1, true); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}, {1: 7, 2: 8, 3: 9}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6\r\n7,8,9")
	checkTCount(t, tb, 9)

	// Disable last row.
	if err := tb.Toggle(2, false); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6")
	checkTCount(t, tb, 9)

	// Enable last row.
	if err := tb.Toggle(2, true); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}, {1: 7, 2: 8, 3: 9}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6\r\n7,8,9")
	checkTCount(t, tb, 9)

	// Enable already-enabled row.
	if err := tb.Toggle(0, true); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 2, 3: 3}, {1: 4, 2: 5, 3: 6}, {1: 7, 2: 8, 3: 9}")
	checkTCSV(t, tb, "1,2,3", "1,2,3\r\n4,5,6\r\n7,8,9")
	checkTCount(t, tb, 9)

	// Disable already-disabled row.
	if err := tb.Toggle(0, false); err != nil {
		t.Error(err)
	}
	if err := tb.Toggle(0, false); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 4, 2: 5, 3: 6}, {1: 7, 2: 8, 3: 9}")
	checkTCSV(t, tb, "1,2,3", "4,5,6\r\n7,8,9")
	checkTCount(t, tb, 9)

	// Remove disabled row.
	tb.RemoveRow(0)
	checkTString(t, tb, "{1: 4, 2: 5, 3: 6}, {1: 7, 2: 8, 3: 9}")
	checkTCSV(t, tb, "1,2,3", "4,5,6\r\n7,8,9")
	checkTCount(t, tb, 6)

	// Remove row before disable row, then enable row.
	if err := tb.Toggle(1, false); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 4, 2: 5, 3: 6}")
	checkTCSV(t, tb, "1,2,3", "4,5,6")
	checkTCount(t, tb, 6)

	tb.RemoveRow(0)
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 3)

	if err := tb.Toggle(0, true); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 7, 2: 8, 3: 9}")
	checkTCSV(t, tb, "1,2,3", "7,8,9")
	checkTCount(t, tb, 3)

	// Try to toggle a row in an empty table.
	tb.RemoveRow(0)
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)

	if err := tb.Toggle(0, true); err == nil {
		t.Error("Unexpectedly toggle nonexistant row")
	}
	checkTString(t, tb, "<empty>")
	checkTCSV(t, tb, "1,2,3", "")
	checkTCount(t, tb, 0)
}

func TestTSortByColumn(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")
	tb.Add(3, 10, 21)
	tb.Add(2, 11, 23)
	tb.Add(1, 12, 22)
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tb, 9)

	// Make sure that only items in the first column are passed to the comparison function. All
	// items will have values between 1 and 9.
	if err := tb.SortByColumn("1", func(l, r interface{}) bool {
		left := l.(int)
		right := r.(int)
		if left < 1 || left > 9 || right < 1 || right > 9 {
			t.Error("Incorrect column passed to comparison function")
		}
		return true
	}); err != nil {
		t.Error(err)
	}

	// Make sure that only items from the second column are passed to the comparison function. All
	// items will have values between 10 and 19.
	if err := tb.SortByColumn("2", func(l, r interface{}) bool {
		left := l.(int)
		right := r.(int)
		if left < 10 || left > 19 || right < 10 || right > 19 {
			t.Error("Incorrect column passed to comparison function")
		}
		return true
	}); err != nil {
		t.Error(err)
	}

	// Make sure that only items from the third column are passed to the comparison function. All
	// items will have values between 20 and 29.
	if err := tb.SortByColumn("3", func(l, r interface{}) bool {
		left := l.(int)
		right := r.(int)
		if left < 20 || left > 29 || right < 20 || right > 29 {
			t.Error("Incorrect column passed to comparison function")
		}
		return true
	}); err != nil {
		t.Error(err)
	}

	// Make sure the table is sorted correctly by the first column.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(3, 10, 21)
	tb.Add(2, 11, 23)
	tb.Add(1, 12, 22)
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tb, 9)

	if err := tb.SortByColumn("1", func(left, right interface{}) bool {
		return left.(int) < right.(int)
	}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 12, 3: 22}, {1: 2, 2: 11, 3: 23}, {1: 3, 2: 10, 3: 21}")
	checkTCSV(t, tb, "1,2,3", "1,12,22\r\n2,11,23\r\n3,10,21")
	checkTCount(t, tb, 9)

	// Make sure the table is sorted correctly by the second column.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(3, 10, 21)
	tb.Add(2, 11, 23)
	tb.Add(1, 12, 22)
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tb, 9)

	if err := tb.SortByColumn("2", func(left, right interface{}) bool {
		return left.(int) < right.(int)
	}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tb, 9)

	// Make sure the table is sorted correctly by the third column.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(3, 10, 21)
	tb.Add(2, 11, 23)
	tb.Add(1, 12, 22)
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tb, 9)

	if err := tb.SortByColumn("3", func(left, right interface{}) bool {
		return left.(int) < right.(int)
	}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 1, 2: 12, 3: 22}, {1: 2, 2: 11, 3: 23}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n1,12,22\r\n2,11,23")
	checkTCount(t, tb, 9)
}

func TestTSortByRow(t *testing.T) {
	tb, _ := htable.New("1", "2", "3")
	tb.Add(3, 10, 21)
	tb.Add(2, 11, 23)
	tb.Add(1, 12, 22)
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tb, 9)

	// Make sure that the rows passed to the comparison function are copies and do not affect the
	// rows in the table.
	if err := tb.SortByRow(func(left, right *htable.Row) bool {
		left.SetItem(0, 100)
		left.SetItem(1, 101)
		left.SetItem(2, 102)
		right.SetItem(0, 200)
		right.SetItem(1, 201)
		right.SetItem(2, 202)
		return true
	}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tb, 9)

	// Make sure the table is sorted correctly by the first column.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(3, 10, 21)
	tb.Add(2, 11, 23)
	tb.Add(1, 12, 22)
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tb, 9)

	if err := tb.SortByRow(func(left, right *htable.Row) bool {
		leftItem := left.Item(0)
		rightItem := right.Item(0)
		return leftItem.(int) < rightItem.(int)
	}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 12, 3: 22}, {1: 2, 2: 11, 3: 23}, {1: 3, 2: 10, 3: 21}")
	checkTCSV(t, tb, "1,2,3", "1,12,22\r\n2,11,23\r\n3,10,21")
	checkTCount(t, tb, 9)

	// Make sure the table is sorted correctly by the second column.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(3, 10, 21)
	tb.Add(2, 11, 23)
	tb.Add(1, 12, 22)
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tb, 9)

	if err := tb.SortByRow(func(left, right *htable.Row) bool {
		leftItem := left.Item(1)
		rightItem := right.Item(1)
		return leftItem.(int) < rightItem.(int)
	}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tb, 9)

	// Make sure the table is sorted correctly by the third column.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(3, 10, 21)
	tb.Add(2, 11, 23)
	tb.Add(1, 12, 22)
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 2, 2: 11, 3: 23}, {1: 1, 2: 12, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n2,11,23\r\n1,12,22")
	checkTCount(t, tb, 9)

	if err := tb.SortByRow(func(left, right *htable.Row) bool {
		leftItem := left.Item(2)
		rightItem := right.Item(2)
		return leftItem.(int) < rightItem.(int)
	}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 3, 2: 10, 3: 21}, {1: 1, 2: 12, 3: 22}, {1: 2, 2: 11, 3: 23}")
	checkTCSV(t, tb, "1,2,3", "3,10,21\r\n1,12,22\r\n2,11,23")
	checkTCount(t, tb, 9)

	// Make sure that we can sort a table according to multiple factors.
	tb, _ = htable.New("1", "2", "3")
	tb.Add(3, 11, 22)
	tb.Add(3, 10, 23)
	tb.Add(3, 11, 21)
	tb.Add(1, 20, 22)
	tb.Add(2, 30, 22)
	tb.Add(2, 31, 22)
	checkTString(t, tb, "{1: 3, 2: 11, 3: 22}, {1: 3, 2: 10, 3: 23}, {1: 3, 2: 11, 3: 21}, {1: 1, 2: 20, 3: 22}, {1: 2, 2: 30, 3: 22}, {1: 2, 2: 31, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "3,11,22\r\n3,10,23\r\n3,11,21\r\n1,20,22\r\n2,30,22\r\n2,31,22")
	checkTCount(t, tb, 18)

	if err := tb.SortByRow(func(left, right *htable.Row) bool {
		// Try to sort on the first column first.
		leftItem := left.Item(0).(int)
		rightItem := right.Item(0).(int)
		if leftItem != rightItem {
			return leftItem < rightItem
		}

		// Next, try to sort on the second column.
		leftItem = left.Item(1).(int)
		rightItem = right.Item(1).(int)
		if leftItem != rightItem {
			return leftItem < rightItem
		}

		// Finally, sort on the third column.
		leftItem = left.Item(2).(int)
		rightItem = right.Item(2).(int)
		return leftItem < rightItem
	}); err != nil {
		t.Error(err)
	}
	checkTString(t, tb, "{1: 1, 2: 20, 3: 22}, {1: 2, 2: 30, 3: 22}, {1: 2, 2: 31, 3: 22}, {1: 3, 2: 10, 3: 23}, {1: 3, 2: 11, 3: 21}, {1: 3, 2: 11, 3: 22}")
	checkTCSV(t, tb, "1,2,3", "1,20,22\r\n2,30,22\r\n2,31,22\r\n3,10,23\r\n3,11,21\r\n3,11,22")
	checkTCount(t, tb, 18)
}

// --- Row's Method Tests ---

func TestRSetItem(t *testing.T) {
	r := htable.NewRow(1, 2, 3)
	checkRString(t, r, "{1, 2, 3}")
	checkRCount(t, r, 3)

	// Try setting the first item.
	if err := r.SetItem(0, 4); err != nil {
		t.Error(err)
	}
	checkRString(t, r, "{4, 2, 3}")
	checkRCount(t, r, 3)

	// Try setting the middle item.
	if err := r.SetItem(1, 5); err != nil {
		t.Error(err)
	}
	checkRString(t, r, "{4, 5, 3}")
	checkRCount(t, r, 3)

	// Try setting the last item.
	if err := r.SetItem(2, 6); err != nil {
		t.Error(err)
	}
	checkRString(t, r, "{4, 5, 6}")
	checkRCount(t, r, 3)

	// Make sure you can set a different type for rows (tables should fail).
	if err := r.SetItem(0, 3.14); err != nil {
		t.Error(err)
	}
	checkRString(t, r, "{3.14, 5, 6}")
	checkRCount(t, r, 3)

	if err := r.SetItem(1, "scooby"); err != nil {
		t.Error(err)
	}
	checkRString(t, r, "{3.14, scooby, 6}")
	checkRCount(t, r, 3)

	if err := r.SetItem(2, []int{10, 20, 30}); err != nil {
		t.Error(err)
	}
	checkRString(t, r, "{3.14, scooby, [10 20 30]}")
	checkRCount(t, r, 3)
}

func TestRItem(t *testing.T) {
	r := htable.NewRow(1, 2, 3)
	checkRString(t, r, "{1, 2, 3}")
	checkRCount(t, r, 3)

	if v := r.Item(0); v != 1 {
		t.Error("Item is row is incorrect")
		t.Log("\tExpected:", 1)
		t.Log("\tReceived:", v)
	}

	if v := r.Item(1); v != 2 {
		t.Error("Item is row is incorrect")
		t.Log("\tExpected:", 2)
		t.Log("\tReceived:", v)
	}

	if v := r.Item(2); v != 3 {
		t.Error("Item is row is incorrect")
		t.Log("\tExpected:", 3)
		t.Log("\tReceived:", v)
	}
}

func TestRItems(t *testing.T) {
	r := htable.NewRow(1, 2, 3)
	checkRString(t, r, "{1, 2, 3}")
	checkRCount(t, r, 3)

	// Make sure that the returned items are correct.
	items := r.Items()
	check := []interface{}{1, 2, 3}
	if !reflect.DeepEqual(items, check) {
		t.Error("Items are incorrect")
		t.Log("\tExpected:", check)
		t.Log("\tReceived:", items)
	}

	// Make sure that a change to the row does not affect the items already returned.
	r.SetItem(0, 4)
	checkRString(t, r, "{4, 2, 3}")
	checkRCount(t, r, 3)
	if !reflect.DeepEqual(items, check) {
		t.Error("Items are incorrect")
		t.Log("\tExpected:", check)
		t.Log("\tReceived:", items)
	}

	// Make sure that the previous change to the row is reflect in the new items.
	items = r.Items()
	check = []interface{}{4, 2, 3}
	if !reflect.DeepEqual(items, check) {
		t.Error("Items are incorrect")
		t.Log("\tExpected:", check)
		t.Log("\tReceived:", items)
	}
}

func TestREnabledToggleRow(t *testing.T) {
	r := htable.NewRow(1, 2, 3)
	checkRString(t, r, "{1, 2, 3}")
	checkRCount(t, r, 3)

	// Make sure that the row is enabled by default.
	if !r.Enabled() {
		t.Error("Row is not marked as enabled by default")
	}

	// Make sure that toggling the row to disabled works.
	r.ToggleRow(false)
	if r.Enabled() {
		t.Error("Row was not toggled to disabled")
	}

	// Toggle is back now.
	r.ToggleRow(true)
	if !r.Enabled() {
		t.Error("Row was not toggled to enabled")
	}

	// Make sure that we can set it to enabled when it's already enabled.
	r.ToggleRow(true)
	r.ToggleRow(true)
	if !r.Enabled() {
		t.Error("Row did not stay enabled")
	}

	// Make sure that we can set it to disabled when it's already disabled.
	r.ToggleRow(false)
	r.ToggleRow(false)
	if r.Enabled() {
		t.Error("Row did not stay disabled")
	}

	// Test that we can toggle rows in a table and have that value reflected correctly in the row
	// itself.
	tb, _ := htable.New("1", "2", "3")
	tb.Add(1, 2, 3)
	tb.Add(4, 5, 6)

	// Right now, both rows should be enabled.
	if row := tb.RowByIndex(0); !row.Enabled() {
		t.Error("Row 0 is not marked as enabled by default")
	}
	if row := tb.RowByIndex(1); !row.Enabled() {
		t.Error("Row 1 is not marked as enabled by default")
	}

	// Make sure that disabling a row in the table is reflected correctly in the row itself.
	tb.Toggle(0, false)
	if row := tb.RowByIndex(0); row.Enabled() {
		t.Error("Row 0 was not toggled to disabled")
	}

	// Make sure that disabling again keeps the value as disabled.
	tb.Toggle(0, false)
	if row := tb.RowByIndex(0); row.Enabled() {
		t.Error("Row 0 did not stay disabled")
	}

	// Make sure that double enabling works.
	tb.Toggle(0, true)
	tb.Toggle(0, true)
	if row := tb.RowByIndex(0); !row.Enabled() {
		t.Error("Row 0 did not stay enabled")
	}

	// Make sure that removing a row does not affect the enabled/disabled status of other rows.
	tb.Toggle(1, false)
	tb.RemoveRow(0)
	if row := tb.RowByIndex(0); row.Enabled() {
		t.Error("Row 0 did not stay disabled after deletion")
	}

	// Make sure that inserting a row does not affect the enabled/disabled status of other rows.
	tb.Insert(0, 7, 8, 9)
	tb.Insert(1, 10, 11, 12)
	if row := tb.RowByIndex(0); !row.Enabled() {
		t.Error("Row 1 did not stay enabled after insertion")
	}
	if row := tb.RowByIndex(2); row.Enabled() {
		t.Error("Row 2 did not stay disabled after insertion")
	}

	// Make sure that pulling out a row and changing its value does affect the row still in the table.
	row := tb.RowByIndex(0)
	row.ToggleRow(false)
	if r := tb.RowByIndex(0); r.Enabled() {
		t.Error("Row 0 was not toggled manually")
	}
}

func TestRMatches(t *testing.T) {
	r := htable.NewRow(1, 2, 3)
	checkRString(t, r, "{1, 2, 3}")
	checkRCount(t, r, 3)

	// Test same types.
	for i, v := range []int{1, 2, 3} {
		// Make sure same values match.
		if !r.Matches(i, v) {
			t.Error("Doesn't match")
		}

		// Make sure different values don't match.
		if r.Matches(i, v+1) {
			t.Error("Matches, but shouldn't")
		}
	}

	// Test different types.
	for i, v := range []string{"1", "2", "3"} {
		if r.Matches(i, v) {
			t.Error("Matches, but shouldn't")
		}
	}

	// Test different types but same binary representation.
	for i, v := range []byte{0x01, 0x02, 0x03} {
		if r.Matches(i, v) {
			t.Error("Matches, but shouldn't")
		}
	}
}

func TestRCopy(t *testing.T) {
	r := htable.NewRow(1, 2, 3)
	checkRString(t, r, "{1, 2, 3}")
	checkRCount(t, r, 3)

	// Make a copy and check that the new row is identical to the original row.
	nr := r.Copy()
	checkRString(t, nr, "{1, 2, 3}")
	checkRCount(t, nr, 3)

	// Make sure that a change in one row doesn't affect the other row.
	r.SetItem(0, 4)
	checkRString(t, r, "{4, 2, 3}")
	checkRCount(t, r, 3)
	checkRString(t, nr, "{1, 2, 3}")
	checkRCount(t, nr, 3)

	nr.SetItem(1, 9)
	checkRString(t, r, "{4, 2, 3}")
	checkRCount(t, r, 3)
	checkRString(t, nr, "{1, 9, 3}")
	checkRCount(t, nr, 3)
}

// --- Helper Functions ---

func checkTString(t *testing.T, tb *htable.Table, want string) {
	if s := tb.String(); s != want {
		t.Error("Table items are incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", s)
	}
}

func checkTCSV(t *testing.T, tb *htable.Table, headers, want string) {
	all := headers
	if want != "" {
		all += "\r\n" + want
	}
	if csv := tb.CSV(); csv != all {
		t.Error("CSV is incorrect")
		t.Log("\tExpected:\n" + all)
		t.Log("\tReceived:\n" + csv)
	}
}

func checkTCount(t *testing.T, tb *htable.Table, want int) {
	if n := tb.Count(); n != want {
		t.Error("Table count is incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", n)
	}
}

func checkRString(t *testing.T, r *htable.Row, want string) {
	if s := r.String(); s != want {
		t.Error("Table items are incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", s)
	}
}

func checkRCount(t *testing.T, r *htable.Row, want int) {
	if n := r.Count(); n != want {
		t.Error("Table count is incorrect")
		t.Log("\tExpected:", want)
		t.Log("\tReceived:", n)
	}
}
