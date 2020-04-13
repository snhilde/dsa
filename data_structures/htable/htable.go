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


var (
	// This is the standard error message when trying to use an invalid table.
	badTable = errors.New("Table must be created with New() first")
	// This is the standard error message when trying to use an invalid row.
	badRow = errors.New("Row must be created with NewRow() first")
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
	// Build the row.
	r := NewRow(items...)

	// Validate that the overall length and the types of all the items are correct.
	if err := t.validateRow(r); err != nil {
		return err
	}

	// Add the row to the table.
	return t.rows.Append(r)
}

// Insert creates a new row with the items and inserts it into the table at the specified index.
func (t *Table) Insert(index int, items ...interface{}) error {
	// Build the row.
	r := NewRow(items...)

	// Validate that the overall length and the types of all the items are correct.
	if err := t.validateRow(r); err != nil {
		return err
	}

	// Add the row to the table.
	return t.rows.Insert(index, r)
}

// AddRow adds the row to the end of the table.
func (t *Table) AddRow(r *Row) error {
	// Validate that the overall length and the types of all the items are correct.
	if err := t.validateRow(r); err != nil {
		return err
	}

	// Add the row to the table.
	return t.rows.Append(r)
}

// InsertRow inserts the row into the table at the specified index.
func (t *Table) InsertRow(index int, r *Row) error {
	// Validate that the overall length and the types of all the items are correct.
	if err := t.validateRow(r); err != nil {
		return err
	}

	// Add the row to the table.
	return t.rows.Insert(index, r)
}

// RemoveRow deletes a row from the table.
func (t *Table) RemoveRow(index int) error {
	if t == nil {
		return badTable
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

// Clear erases the rows in the table but leaves the column headers and column types.
func (t *Table) Clear() error {
	if t == nil {
		return badTable
	}

	return t.rows.Clear()
}

// ColumnToIndex translates the column's string name to its index in the table. This will return -1 on error.
func (t *Table) ColumnToIndex(col string) int {
	if t == nil {
		return -1
	}

	for i, v := range t.h {
		if col == v {
			return i
		}
	}

	// If we're here, then we didn't find anything.
	return -1
}

// String returns a formatted list of the items in the table, row by row.
func (t *Table) String() string {
	if t == nil {
		return "<nil>"
	} else if t.Count() == 0 {
		return "<empty>"
	}

	var b strings.Builder
	rows := t.rows.YieldAll()
	for r := range rows {
		row := r.(*Row)
		if row.enabled {
			var tmp strings.Builder
			for i, v := range row.v {
				tmp.WriteString(fmt.Sprintf("%v: %v, ", t.h[i], v))
			}
			s := tmp.String()
			s = strings.TrimSuffix(s, ", ")
			s = strings.Join([]string{"{", s, "}"}, "")
			b.WriteString(fmt.Sprintf("%v, ", s))
		}
	}

	s := b.String()
	if s == "" {
		s = "<empty>"
	}

	return strings.TrimSuffix(s, ", ")
}

// SetItem changes the value of the item at the specified coordinates.
func (t *Table) SetItem(row int, col string, value interface{}) error {
	if t == nil {
		return badTable
	} else if row < 0 || t.Rows() <= row {
		return errors.New("Invalid row")
	}

	// Grab our row.
	r := t.rows.Value(row)
	if r == nil {
		return errors.New("Missing row")
	}

	// Figure out the index of the column.
	i := t.ColumnToIndex(col)
	if i < 0 {
		return errors.New("Invalid column")
	}

	// Change the value.
	nr := r.(*Row)
	nr.SetItem(i, value)

	// Add the row back in to the table, which will also validate the new value.
	if err := t.RemoveRow(row); err != nil {
		return err
	}
	if err := t.InsertRow(row, nr); err != nil {
		return err
	}

	return nil
}

// SetHeader changes the specified column's header to name.
func (t *Table) SetHeader(col string, name string) error {
	if t == nil {
		return badTable
	} else if col == "" {
		return errors.New("Invalid column")
	} else if name == "" {
		return errors.New("Missing name")
	}

	for i, v := range t.h {
		if col == v {
			t.h[i] = name
			return nil
		}
	}

	// If we're here, then we didn't find the column.
	return errors.New("Missing column")
}

// Headers returns a copy of the table's column headers.
func (t *Table) Headers() []string {
	if t == nil {
		return nil
	}

	h := make([]string, t.Columns())
	copy(h, t.h)

	return h
}

// Rows returns the number of rows in the table, or -1 on error. This will include all rows, regardless of enabled status.
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

// Count returns the number of items in the table, or -1 on error. This will include all items, regardless of enabled status.
func (t *Table) Count() int {
	r := t.Rows()
	c := t.Columns()

	if r == -1 || c == -1 {
		return -1
	}

	return r * c
}

// Row returns the index and Row type of the first row that contains the item in the specified column, or -1 and nil if
// not found or error.
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
		if row.enabled {
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

	// Figure out the index of the column.
	i := t.ColumnToIndex(col)
	if i < 0 {
		return nil
	}

	return r.(*Row).Item(i)
}

// Matches returns true if the value matches the item at the specified coordinates or false if there is no match.
// Matching can occur on disabled rows.
func (t *Table) Matches(row int, col string, v interface{}) bool {
	item := t.Item(row, col)
	return reflect.DeepEqual(v, item)
}

// Toggle sets the row at the specified index to either be checked or skipped during table lookups (like Row and Count).
func (t *Table) Toggle(row int, enabled bool) error {
	if t == nil {
		return badTable
	}

	tmp := t.rows.Value(row)
	if tmp == nil {
		return errors.New("Invalid index")
	}

	r := tmp.(*Row)
	r.enabled = enabled

	return nil
}

// WriteCSV converts the table into rows of comma-separated values, with each row delineated by \r\n newlines.
func (t *Table) WriteCSV() string {
	if t == nil || t.Rows() < 1 {
		return ""
	}

	var b strings.Builder
	rows := t.rows.YieldAll()
	for r := range rows {
		row := r.(*Row)
		if row.enabled {
			items := make([]string, len(row.v))
			for i, v := range row.v {
				items[i] = fmt.Sprintf("%v", v)
			}
			b.WriteString(strings.Join(items, ","))
			b.WriteString("\r\n")
		}
	}

	// Remove the last newline before returning the string.
	return strings.TrimSuffix(b.String(), "\r\n")
}


// Row holds all the data for each row in the table.
type Row struct {
	v       []interface{} // Column values
	enabled   bool
}

// NewRow creates a new row with the given items.
func NewRow(items ...interface{}) *Row {
	if items == nil || len(items) == 0 {
		return nil
	}

	r := new(Row)

	// Add the items.
	r.v = make([]interface{}, len(items))
	copy(r.v, items)

	// Default to enabling this row.
	r.enabled = true

	return r
}

// String returns a formatted list of the items in the row.
func (r *Row) String() string {
	if r == nil {
		return "<nil>"
	} else if len(r.v) == 0 {
		return "<empty>"
	}

	var b strings.Builder
	for _, v := range r.v {
		b.WriteString(fmt.Sprintf("%v, ", v))
	}

	s := strings.TrimSuffix(b.String(), ", ")

	return strings.Join([]string{"{", s, "}"}, "")
}

// SetItem changes the value of the item in the specified column.
func (r *Row) SetItem(index int, value interface{}) error {
	if r == nil {
		return badRow
	} else if index < 0 || r.Count() <= index {
		return errors.New("Invalid column")
	}

	r.v[index] = value

	return nil
}

// Count returns the number of items in the row, or -1 on error.
func (r *Row) Count() int {
	if r == nil {
		return -1
	}

	return len(r.v)
}

// Item returns the row's value at the specified index, or nil if not found or error.
func (r *Row) Item(index int) interface{} {
	if r == nil {
		return nil
	} else if index < 0 {
		return nil
	} else if index >= len(r.v) {
		return nil
	}

	return r.v[index]
}

// Matches returns true if the value matches the item in the specified column or false if there is no match.
// Matching can occur on disabled rows.
func (r *Row) Matches(index int, v interface{}) bool {
	item := r.Item(index)
	return reflect.DeepEqual(v, item)
}


func (t *Table) validateRow(r *Row) error {
	if t == nil {
		return badTable
	} else if n := t.Columns(); n != len(r.v) {
		return errors.New(fmt.Sprintf("Number of items (%v) does not match number of columns (%v)", len(r.v), n))
	}

	first := false
	if t.Rows() == 0 {
		// This is the first row being added to the table. It will set the type of each column in the table.
		first = true
	}

	// Validate the types.
	for i, v := range r.v {
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
				return errors.New(fmt.Sprintf("Item %v's type (%v) does not match column's prototype (%v)", i, k, t.types[i]))
			}
		}
	}

	return nil
}
