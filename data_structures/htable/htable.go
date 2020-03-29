// Package htable is a data structure of rows and columns, with each row having the same number of items and each column
// holding the same type of data. The tables provide easy building and quick lookup for a fast implementation of storing
// and accessing uniform lists.
package htable

import (
	"github.com/snhilde/dsa/data_structures/hlist"
	"reflect"
	"errors"
	"fmt"
	"strings"
)


// Table is the main type in this package. It holds all the rows of data.
type Table struct {
	h     []string       // Column headers
	types []reflect.Kind // Types of each column (must be consistent for all rows)
	rows   *hlist.List   // Linked list of rows
}


// New creates a new table. The strings will denote the names of each column, used during lookup.
func New(headers ...string) (*Table, error) {
	if headers == nil || len(headers) == 0 {
		return nil, errors.New("Missing column headers")
	}

	// Validate the column headers.
	for i, v := range headers {
		// Make sure every column has a header.
		if v == "" {
			return nil, errors.New(fmt.Sprintf("Column %v has an empty header", i))
		}
		// Make sure none of the columns match each other. Not the most efficient, but necessary.
		for j, w := range headers[i+1:] {
			if v == w {
				return nil, errors.New(fmt.Sprintf("Columns %v and %v have the same header", i, j))
			}
		}
	}

	var t Table
	t.h = headers
	t.types = make([]reflect.Kind, len(headers))
	t.rows = hlist.New()

	return &t, nil
}


// Add creates a new row with the items and adds it to the end of the table.
func (t *Table) Add(items ...interface{}) error {
	row, err := t.newRow(items...)
	if err != nil {
		return err
	}

	return t.rows.Append(row)
}

// InsertRow inserts a new row of items at the specified index.
func (t *Table) InsertRow(index int, items ...interface{}) error {
	row, err := t.newRow(items...)
	if err != nil {
		return err
	}

	return t.rows.Insert(index, row)
}

// RemoveRow deletes a row from the table.
func (t *Table) RemoveRow(index int) error {
	if t == nil {
		return tErr()
	}

	v := t.rows.Remove(index)
	if v == nil {
		// hlist.Remove will return the value at the index. Because our rows can never be nil, if we receive a nil
		// value, then it means an error occurred.
		return errors.New(fmt.Sprintf("Failed to remove row %v", index))
	}

	// All good
	return nil
}

// Rows returns the number of rows in the table, or -1 on error. This will include all rows, regardless of check status.
func (t *Table) Rows() int {
	if t == nil {
		return -1
	}

	return t.rows.Length()
}

// Columns returns the number of columns in the table, or -1 on error.
func (t *Table) Columns() int {
	if t == nil {
		return -1
	}

	return len(t.h)
}

// Count returns the number of items in the table, or -1 on error. This will include all items, regardless of check status.
func (t *Table) Count() int {
	r := t.Rows()
	c := t.Columns()

	if r == -1 || c == -1 {
		return -1
	}

	return r * c
}

// String returns a printable list of values, grouped by row.
func (t *Table) String() string {
	if t == nil {
		return "<nil>"
	} else if t.Count() == 0 {
		return "<empty>"
	}

	var b strings.Builder
	rows := t.rows.YieldAll()
	for v := range rows {
		row := v.(*Row)
		if row.check {
			b.WriteString(fmt.Sprintf("%v, ", row.v))
		}
	}

	s := b.String()
	if s == "" {
		s = "<empty>"
	}

	return strings.TrimSuffix(s, ", ")
}

// Row returns the first index and first row that contains the item in the specified column, or -1 and nil if not found
// or error.
func (t *Table) Row(col string, item interface{}) (int, *Row) {
	if t == nil {
		return -1, nil
	}

	// Find out which column we need to match on.
	c := -1
	for i, v := range t.h {
		if col == v {
			c = i
			break
		}
	}

	// Make sure we found a column.
	if c == -1 {
		return -1, nil
	}

	// Get our iterator to go through the rows.
	i := 0
	quit := make(chan interface{})
	rows := t.rows.Yield(quit)
	for v := range rows {
		row := v.(*Row)
		if row.check {
			if reflect.DeepEqual(item, row.v[c]) {
				// Break out of the list iteration. If Yield's goroutine has already exited (because the list was fully
				// traversed), then it won't receive the message to quit. We'll try to send the quit message, and then
				// we'll exit.
				select {
				case quit <- 0:
				default:
					break
				}
				return i, row
			}
		}
		i++
	}

	// If we're here, then we didn't find anything.
	return -1, nil
}

// Item returns the item at the specified coordinates, or nil if there is no item at the coordinates.
func (t *Table) Item(row int, col string) interface{} {
	if t == nil {
		return nil
	}

	// Grab our row.
	r := t.rows.Value(row)
	if r == nil {
		return nil
	}

	return r.(*Row).Item(col)
}

// Matches returns true if the value matches the item at the specified coordinates or false if there is no match.
// Matching can occur on rows toggled off.
func (t *Table) Matches(row int, col string, v interface{}) bool {
	item := t.Item(row, col)
	return reflect.DeepEqual(v, item)
}

// Toggle sets the row at the specified index to either be checked or skipped during table lookups (like Row and Count).
func (t *Table) Toggle(row int, check bool) error {
	if t == nil {
		return tErr()
	}

	tmp := t.rows.Value(row)
	if tmp == nil {
		return errors.New("Invalid index")
	}

	r := tmp.(*Row)
	r.check = check

	return nil
}


// Row holds all the data for each row in the table.
type Row struct {
	h     []string      // Column headers
	v     []interface{} // Column values
	check   bool
}

// Item returns the value at the specified column for this row, or nil if not found or error.
func (r *Row) Item(col string) interface{} {
	if r == nil {
		return nil
	}

	for i, v := range r.h {
		if col == v {
			return r.v[i]
		}
	}

	// If we're here, then we didn't find anything.
	return nil
}

// Matches returns true if the value matches the item in the specified column or false if there is no match.
// Matching can occur on rows toggled off.
func (r *Row) Matches(col string, v interface{}) bool {
	item := r.Item(col)
	return reflect.DeepEqual(v, item)
}


func tErr() error {
	return errors.New("Table must be created with New() first")
}

func (t *Table) newRow(items ...interface{}) (*Row, error) {
	if t == nil {
		return nil, tErr()
	} else if n := t.Columns(); n != len(items) {
		return nil, errors.New(fmt.Sprintf("Number of items (%v) does not match number of columns (%v)", len(items), n))
	}

	first := false
	if t.Rows() == 0 {
		// This is the first row being added to the table. It will set the type of each column in the table.
		first = true
	}

	// Build out our row.
	row := new(Row)
	row.h = t.h
	row.v = make([]interface{}, t.Columns())
	row.check = true

	// Add the items to the row.
	for i, v := range items {
		rv := reflect.ValueOf(v)
		k := rv.Kind()
		switch k {
		// Always default to the common type.
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			k = reflect.Int
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			k = reflect.Uint
		case reflect.Float32, reflect.Float64:
			k = reflect.Float64
		case reflect.Complex64, reflect.Complex128:
			k = reflect.Complex128
		}

		if first {
			t.types[i] = k
		} else {
			// Make sure the type of this element matches the prototype.
			if k != t.types[i] {
				return nil, errors.New(fmt.Sprintf("Item %v's type (%v) does not match column's prototype (%v)", i, k, t.types[i]))
			}
		}
		row.v[i] = v
	}

	return row, nil
}
